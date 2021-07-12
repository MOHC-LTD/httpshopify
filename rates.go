package httpshopify

import "github.com/MOHC-LTD/httpshopify/internal/http"

// RateLimitDefault builds the default rate limit for the rest API.
/*
	This limit is as follows:
	BucketSize: 40 requests/app/store
	LeakRate: 2/second

	Source: https://shopify.dev/api/usage/rate-limits.
*/
func RateLimitDefault() http.Option {
	return http.WithRateLimit(
		40,
		2,
	)
}

// RateLimitPlus builds the rate limit for the rest API of a Shopify plus store.
/*
	This limit is as follows:
	BucketSize: 80 requests/app/store
	LeakRate: 4/second

	Source: https://shopify.dev/api/usage/rate-limits.
*/
func RateLimitPlus() http.Option {
	return http.WithRateLimit(
		80,
		4,
	)
}

const (
	// IsPlus represents a shop being a plus store
	IsPlus = true
	// IsDefault represents a shop being a default store
	IsDefault = false
)
