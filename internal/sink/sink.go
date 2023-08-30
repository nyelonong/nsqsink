package sink

type Sink interface {
	// Write writes the given data to the sink.
	Write(data []byte) error

	// Close closes the sink.
	Close() error

	// // Name returns the name of the sink.
	// Name() string

	// // String returns a string representation of the sink.
	// String() string

	// // Type returns the type of the sink.
	// Type() string

	// // SetType sets the type of the sink.
	// SetType(string)

	// // SetName sets the name of the sink.
	// SetName(string)

	// // SetConfig sets the config of the sink.
	// SetConfig(map[string]interface{})
}
