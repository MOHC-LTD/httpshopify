package http

import (
	"bytes"
	"context"
	"io"
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
	retryMaxDuration  time.Duration
	retryBaseDuration time.Duration
	retryCount        int
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

func (c Client) retryDuration(retry int) time.Duration {
	duration := c.retryBaseDuration * (1 << retry)
	jitter := time.Duration(rand.Float64() * float64(time.Second))

	if duration > c.retryMaxDuration {
		return c.retryMaxDuration + jitter
	}

	return duration + jitter
}

// Do does a request
func (c Client) Do(method string, url string, headers RequestHeaders, body io.Reader) ([]byte, ResponseHeaders, error) {
	var requestBody []byte
	var err error

	if body != nil {
		requestBody, err = io.ReadAll(body)
		if err != nil {
			return nil, ResponseHeaders{}, err
		}
	}

	if c.limiter != nil {
		ctx := context.Background()
		c.limiter.Wait(ctx)
	}

	headers = c.AppendDefaultHeaders(headers)

	var resp *http.Response

	for i := 0; i < c.retryCount+1; i++ {
		req, err := http.NewRequest(method, url, bytes.NewReader(requestBody))
		if err != nil {
			return nil, ResponseHeaders{}, err
		}

		for _, header := range headers {
			req.Header.Set(header.Name, header.Value)
		}

		resp, err = c.client.Do(req)

		// Break if client config not set for retries
		if c.retryCount == i {
			break
		}

		// Default wait time to exp backoff
		waitTime := c.retryDuration(i)

		if err == nil {
			defer resp.Body.Close()

			isRetryResponse := resp.StatusCode == 429 || resp.StatusCode >= 500
			if !isRetryResponse {
				break
			}

			// Use Shopify's retry after duration if exists
			if retryAfterHeader := resp.Header.Get("Retry-After"); retryAfterHeader != "" {
				retryAfter, _ := strconv.ParseFloat(retryAfterHeader, 64)
				waitTime = time.Duration(retryAfter * float64(time.Second))
			}
		}

		time.Sleep(waitTime)
	}

	if err != nil {
		return nil, ResponseHeaders{}, err
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, ResponseHeaders{}, err
	}

	err = HandleStatus(resp.StatusCode, responseBody)
	if err != nil {
		return nil, ResponseHeaders{}, err
	}

	return responseBody, ResponseHeaders{resp.Header}, nil
}
