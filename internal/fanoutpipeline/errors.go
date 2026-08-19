package fanoutpipeline

func NewErrorChannel(workers int) chan error {
	if workers < 1 {
		workers = 1
	}
	return make(chan error, workers)
}
