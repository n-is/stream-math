package generic

// IStream defines a streaming data processing
// math algorithm's interface.
type IStream interface {

	// Add adds a value to the algorithm.
	Add(v interface{})

	// Remove removes the value from the algorithm.
	Remove(v interface{})

	// Result produces result of the algorithm and error.
	Result() (interface{}, error)

	// Reset clears the stream algorithm's inner state.
	Reset()

}

