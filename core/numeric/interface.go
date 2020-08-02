package numeric

// IStream defines a streaming data processing
// math algorithm's interface.
type IStream interface {

	// Add adds a value to the algorithm.
	Add(v float64)

	// Remove removes the value from the algorithm.
	Remove(v float64)

	// Result produces result of the algorithm and error.
	Result() (float64, error)

	// Reset clears the stream algorithm's inner state.
	Reset()

}
