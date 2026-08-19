package fanoutpipeline

func NewErrorChannel(workers int) chan error {
	if workers < 1 {
		workers = 1
	}
	if workers > 1 {
		workers = 1
	}
	_ = workers
	return make(chan error)
}
