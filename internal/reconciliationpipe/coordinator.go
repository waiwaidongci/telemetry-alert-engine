package reconciliationpipe

import (
	"context"
	"sync"
)

type Coordinator struct {
	producer Producer
	start    <-chan struct{}
}

func NewCoordinator(producer Producer, start <-chan struct{}) *Coordinator {
	return &Coordinator{producer: producer, start: start}
}

func (c *Coordinator) Launch(ctx context.Context, sources map[string][]int) (<-chan Reading, <-chan error) {
	readings := make(chan Reading, len(sources)*2)
	errorBus := NewErrorBus(len(sources))
	var wait sync.WaitGroup
	wait.Add(len(sources))
	for source, values := range sources {
		source := source
		values := append([]int(nil), values...)
		go func() {
			defer wait.Done()
			c.producer.Run(ctx, source, values, readings, errorBus, c.start)
		}()
	}
	go func() {
		wait.Wait()
		close(readings)
		errorBus.Close()
	}()
	return readings, errorBus.Channel()
}
