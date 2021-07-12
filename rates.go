package httpshopify

// RateLimit represents a Shopify rest rate limit.
/*
	These limits use the leaky bucket algorithm.
	For more details, see https://shopify.dev/api/usage/rate-limits.
*/
type RateLimit struct {
	BucketSize int
	LeakRate   int
}

// RateLimitDefault builds the default rate limit for the rest API.
/*
	This limit is as follows:
	BucketSize: 40 requests/app/store
	LeakRate: 2/second

	Source: https://shopify.dev/api/usage/rate-limits.
*/
func RateLimitDefault() RateLimit {
	return RateLimit{
		BucketSize: 40,
		LeakRate:   2,
	}
}

// RateLimitPlus builds the rate limit for the rest API of a Shopify plus store.
/*
	This limit is as follows:
	BucketSize: 80 requests/app/store
	LeakRate: 4/second

	Source: https://shopify.dev/api/usage/rate-limits.
*/
func RateLimitPlus() RateLimit {
	return RateLimit{
		BucketSize: 40,
		LeakRate:   2,
	}
}
