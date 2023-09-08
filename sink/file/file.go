package file

import (
	"context"
	"fmt"
	"os"

	"github.com/nyelonong/nsqsink/sink"
)

type Client struct {
	file *os.File
}

func NewSink(fileName string) (sink.Sinker, error) {
	f, err := os.Create(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to open file because of %s", err.Error())
	}

	return &Client{
		file: f,
	}, nil
}

func (c *Client) Write(ctx context.Context, data []byte) ([]byte, error) {
	_, err := c.file.Write(data)
	if err != nil {
		c.file.Close()
		return nil, err
	}
	return data, nil
}

func (c *Client) Close(ctx context.Context) error {
	return c.file.Close()
}
