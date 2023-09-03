package streamer

import "context"

// Streamer is an interface which can be implemented
// to handle messages produced by a stream source
type Streamer interface {
	// RegisterHandler method
	// method to register handler to the streamer, mapping event to handler
	RegisterHandler(ctx context.Context, e Event, h Handler) error

	// Run method
	// method to run all handler in the streamer
	Run() error

	// Stop method
	// method to stop all handler in the streamer
	Stop() error
}

// StreamerType
// type for streamer type
type StreamerType string

// const streamer type
const (
	StreamerNSQ StreamerType = "nsq"
	// StreamerKafka string = "kafka"
)
