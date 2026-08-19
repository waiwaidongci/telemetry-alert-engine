package fanoutpipeline

// NewErrorChannel returns a buffered error channel sized for the given number
// of workers, so that every worker can report its error without blocking.
func NewErrorChannel(workers int) chan error {
	if workers < 1 {
		workers = 1
	}
	return make(chan error, workers)
}
