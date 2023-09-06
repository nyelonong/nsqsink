package http

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var (
	ErrMaxRetriesExceeded = errors.New("max retries exceeded")
)

const (
	// DefaultTimeout is the default timeout for the HTTP request.
	DefaultTimeout = 5 * time.Second

	// DefaultBackoff is the default amount of time to wait before retrying a failed request.
	DefaultBackoff = 5 * time.Second

	// DefaultMaxRetries is the default maximum number of times to retry a failed request.
	DefaultMaxRetries = 1
)

type Client struct {
	// client is the HTTP client to use.
	client *http.Client

	// URL is the URL to send the data to.
	URL string `json:"url" yaml:"url"`

	// Headers is a map of headers to add to the request.
	Headers map[string]string `json:"headers" yaml:"headers"`

	// Backoff is the amount of time to wait before retrying a failed request.
	Backoff time.Duration `json:"backoff" yaml:"backoff"`

	// MaxRetries is the maximum number of times to retry a failed request.
	MaxRetries int `json:"max_retries" yaml:"max_retries"`

	// RetryStatusCode is the status code to retry on.
	RetryStatusCode []int `json:"retry_status_code" yaml:"retry_status_code"`

	// Method is the HTTP method to use.
	Method string `json:"method" yaml:"method"`

	// URLParam is the URL parameters to use.
	Param url.Values `json:"param" yaml:"param"`

	// BasicAuth is the basic authentication to use.
	BasicAuth BasicAuth `json:"basic_auth" yaml:"basic_auth"`

	// BearerAuth is the bearer authentication to use.
	BearerAuth string `json:"bearer_auth" yaml:"bearer_auth"`

	// DataFormat is the data format to use.
	DataFormat string `json:"data_format" yaml:"data_format"`
}

type BasicAuth struct {
	Username string
	Password string
}

type Option func(*Client) error

func WithTimeout(timeout time.Duration) Option {
	return func(h *Client) error {
		h.client.Timeout = timeout
		return nil
	}
}

func NewClient(options ...Option) (*Client, error) {
	h := &Client{
		Headers: map[string]string{},
		client: &http.Client{
			Timeout: DefaultTimeout,
		},
		Backoff:    DefaultBackoff,
		MaxRetries: DefaultMaxRetries,
	}

	for _, option := range options {
		if err := option(h); err != nil {
			return h, err
		}
	}

	return h, nil
}

func (c *Client) Write(ctx context.Context, data []byte) ([]byte, error) {
	c.setContentType()

	resp, err := c.do(ctx, data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (c *Client) Close() error {
	c.client.CloseIdleConnections()
	return nil
}

func (c *Client) do(ctx context.Context, body interface{}) (*http.Response, error) {
	bodyReader, err := toReader(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(c.Method, c.URL, bodyReader)
	if err != nil {
		return nil, err
	}

	for key, value := range c.Headers {
		req.Header.Set(key, value)
	}

	req = req.WithContext(ctx)
	req.Close = true

	var resp *http.Response

	for i := 0; i < c.MaxRetries; i++ {
		resp, err = c.client.Do(req)
		if err != nil {
			time.Sleep(c.Backoff)
			continue
		}

		if c.RetryStatusCode != nil {
			for _, code := range c.RetryStatusCode {
				if resp.StatusCode == code {
					resp.Body.Close()
					time.Sleep(c.Backoff)
					continue
				}
			}
		}

		return resp, nil
	}

	return nil, ErrMaxRetriesExceeded
}

func (c *Client) setContentType() {
	switch c.DataFormat {
	case "json":
		c.Headers["Content-Type"] = "application/json"
	case "xml":
		c.Headers["Content-Type"] = "application/xml"
	case "form":
		c.Headers["Content-Type"] = "application/x-www-form-urlencoded"
	case "text":
		c.Headers["Content-Type"] = "text/plain"
	case "html":
		c.Headers["Content-Type"] = "text/html"
	case "multipart":
		c.Headers["Content-Type"] = "multipart/form-data"
	default:
		c.Headers["Content-Type"] = "application/json"
	}
}

func toReader(body interface{}) (io.Reader, error) {
	var bodyReader io.Reader

	switch v := body.(type) {
	case io.Reader:
		bodyReader = v
	case string:
		bodyReader = strings.NewReader(v)
	case []byte:
		bodyReader = bytes.NewReader(v)
	case nil:
	default:
		b, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}

		bodyReader = bytes.NewReader(b)
	}

	return bodyReader, nil
}