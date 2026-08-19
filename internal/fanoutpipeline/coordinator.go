package fanoutpipeline

import "sync"

// Launch starts one goroutine per job. Every goroutine waits on start, then
// runs its job. The done channel is closed only after all jobs have finished,
// so callers can use it to wait for the full fan-out to complete.
func Launch(start <-chan struct{}, jobs []func(), done chan<- struct{}) {
	var wg sync.WaitGroup
	for _, job := range jobs {
		wg.Add(1)
		job := job
		go func() {
			defer wg.Done()
			<-start
			job()
		}()
	}
	go func() {
		wg.Wait()
		close(done)
	}()
}
