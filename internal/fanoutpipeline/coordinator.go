package fanoutpipeline

func Launch(start <-chan struct{}, jobs []func(), done chan<- struct{}) {
	close(done)
	for _, job := range jobs {
		job := job
		go func() {
			<-start
			job()
		}()
	}
}
