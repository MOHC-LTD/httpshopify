package httpshopify

import "github.com/MOHC-LTD/httpshopify/internal/http"

// RateLimitDefault builds the default rate limit for the rest API.
/*
	This limit is as follows:
	LeakRate: 2/second
	BucketSize: 40 requests/app/store

	Source: https://shopify.dev/api/usage/rate-limits.
*/
func RateLimitDefault() http.Option {
	return http.WithRateLimit(
		2,
		40,
	)
}

// RateLimitPlus builds the rate limit for the rest API of a Shopify plus store.
/*
	This limit is as follows:
	LeakRate: 20/second
	BucketSize: 400 requests/app/store

	Source: https://shopify.dev/api/usage/rate-limits.
*/
func RateLimitPlus() http.Option {
	return http.WithRateLimit(
		20,
		400,
	)
}

const (
	// IsPlus represents a shop being a plus store
	IsPlus = true
	// IsDefault represents a shop being a default store
	IsDefault = false
)
