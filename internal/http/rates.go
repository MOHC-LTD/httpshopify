package http

import "golang.org/x/time/rate"

// OptionRateLimit holds configuration for rate limiting requests.
type OptionRateLimit struct {
	rate  int
	burst int
}

func (option OptionRateLimit) configure(client *Client) error {
	client.limiter = rate.NewLimiter(rate.Limit(option.rate), option.burst)

	return nil
}

// WithRateLimit allows configuring the rate limit for requests.
/*
	Based on the https://pkg.go.dev/golang.org/x/time package.
*/
func WithRateLimit(rate int, burst int) OptionRateLimit {
	return OptionRateLimit{
		rate,
		burst,
	}
}
