package http

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"golang.org/x/time/rate"
)

// Client is a HTTP client
type Client struct {
	client            *http.Client
	defaultHeaders    RequestHeaders
	limiter           *rate.Limiter
	RetryMaxDuration  time.Duration
	RetryBaseDuration time.Duration
	RetryCount        int
}

// NewClient builds a new HTTP client
func NewClient(options ...Option) Client {
	client := Client{
		client: &http.Client{},
	}

	for _, option := range options {
		option.configure(&client)
	}

	return client
}

// AppendDefaultHeaders appends the default headers to the passed ones.
func (c Client) AppendDefaultHeaders(headers RequestHeaders) RequestHeaders {
	for _, header := range c.defaultHeaders {
		if !headers.Includes(header.Name) {
			headers = append(headers, header)
		}
	}

	return headers
}

// Do does a request
func (c Client) Do(method string, url string, headers RequestHeaders, body io.Reader) ([]byte, ResponseHeaders, error) {
	if c.limiter != nil {
		ctx := context.Background()
		c.limiter.Wait(ctx)
	}

	headers = c.AppendDefaultHeaders(headers)

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, ResponseHeaders{}, err
	}

	for _, header := range headers {
		req.Header.Set(header.Name, header.Value)
	}

	var resp *http.Response
	for i := 0; i < c.RetryCount+1; i++ {
		resp, err = c.client.Do(req)

		// Return early if error
		if err != nil {
			if i == c.RetryCount {
				return nil, ResponseHeaders{}, err
			}

			continue
		}

		if resp != nil {
			defer resp.Body.Close()
		}

		// Break if client config not set for retries or retries maxed out
		if c.RetryCount == 0 || c.RetryCount == i {
			break
		}

		// Break since we have a response
		if resp.StatusCode != 429 && resp.StatusCode < 500 {
			break
		}

		// Retry using Shopify's retry after duration if any
		if retryAfterHeader := resp.Header.Get("Retry-After"); retryAfterHeader != "" {
			retryAfter, _ := strconv.Atoi(retryAfterHeader)
			waitTime := time.Duration(retryAfter) * time.Second
			fmt.Printf("Rate limit was hit. Retry %d/%d starting in %v seconds\n", i+1, c.RetryCount, waitTime.Seconds())
			time.Sleep(waitTime)

			continue
		}

		// Retry defaulting to exponential backoff and jitter
		waitTime := withExponentialBackOff(i, c.RetryBaseDuration, c.RetryMaxDuration)
		withJitter(&waitTime)
		fmt.Printf("Retry %d/%d starting in %v seconds.\n", i+1, c.RetryCount, waitTime.Seconds())
		time.Sleep(waitTime)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, ResponseHeaders{}, err
	}

	err = HandleStatus(resp.StatusCode, responseBody)
	if err != nil {
		return nil, ResponseHeaders{}, err
	}

	return responseBody, ResponseHeaders{resp.Header}, nil
}

func withExponentialBackOff(retry int, base time.Duration, max time.Duration) time.Duration {
	backoff := base * (1 << retry)

	if backoff > max {
		return max
	}

	return backoff
}

func withJitter(waitTime *time.Duration) {
	*waitTime += time.Duration(rand.Float64() * float64(time.Second))
}
