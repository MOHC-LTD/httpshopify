package httpshopify

import "time"

type Options struct {
	retryCount        int
	retryBaseDuration time.Duration
	retryMaxDuration  time.Duration
}

// OptionFunc is a function that sets options on the Options struct
type OptionFunc func(*Options)

// WithExponentialBackoff configures the client to use exponential backoff
// on responses that indicate the request can be retried. Typically this is when
// Shopify returns a 429 or 503 HTTP status code
func WithExponentialBackoff(retryCount int, retryBaseDuration time.Duration, retryMaxDuration time.Duration) OptionFunc {
	return func(o *Options) {
		o.retryCount = retryCount
		o.retryBaseDuration = retryBaseDuration
		o.retryMaxDuration = retryMaxDuration
	}
}
