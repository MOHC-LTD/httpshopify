package http

import "time"

// Option allows the client to be fully configurable
type Option interface {
	// configure configures the client with this option
	configure(client *Client) error
}

type BackOffOptions struct {
	retryCount        int
	retryBaseDuration time.Duration
	retryMaxDuration  time.Duration
}

func (option BackOffOptions) configure(client *Client) error {
	client.retryCount = option.retryCount
	client.retryBaseDuration = option.retryBaseDuration
	client.retryMaxDuration = option.retryMaxDuration

	return nil
}

func WithBackoffOptions(retryCount int, retryBaseDuration time.Duration, retryMaxDuration time.Duration) BackOffOptions {
	return BackOffOptions{
		retryCount,
		retryBaseDuration,
		retryMaxDuration,
	}
}
