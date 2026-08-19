package fanoutpipeline

import "sync"

func Launch(start <-chan struct{}, jobs []func(), done chan<- struct{}) {
	var workers sync.WaitGroup
	workers.Add(len(jobs))
	for _, job := range jobs {
		job := job
		go func() {
			defer workers.Done()
			<-start
			job()
		}()
	}
	go func() { workers.Wait(); close(done) }()
}
