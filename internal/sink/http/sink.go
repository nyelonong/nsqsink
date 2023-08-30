package http

import "github.com/nyelonong/nsqsink/internal/sink"

type HTTP struct {
	// URL is the URL to send the data to.
	URL string `json:"url" yaml:"url"`

	// Method is the HTTP method to use.
	Method string `json:"method" yaml:"method"`

	// Headers is a map of headers to add to the request.
	Headers map[string]string `json:"headers" yaml:"headers"`

	// Query is a map of query parameters to add to the request.
	Query map[string]string `json:"query" yaml:"query"`

	// Timeout is the maximum duration to wait for a response.
	Timeout string `json:"timeout" yaml:"timeout"`

	// Backoff is the amount of time to wait before retrying a failed request.
	Backoff string `json:"backoff" yaml:"backoff"`

	// MaxRetries is the maximum number of times to retry a failed request.
	MaxRetries int `json:"max_retries" yaml:"max_retries"`

	Sink sink.Sink
}

func (h *HTTP) Write(data []byte) error {
	return nil
}

func (h *HTTP) Close() error {
	return nil
}
