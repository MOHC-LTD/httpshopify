package http

// Option allows the client to be fully configurable
type Option interface {
	// configure configures the client with this option
	configure(client *Client) error
}
