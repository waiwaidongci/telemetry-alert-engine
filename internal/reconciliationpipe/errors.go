package reconciliationpipe

import "fmt"

type ErrorBus struct {
	channel chan error
}

func NewErrorBus(capacity int) *ErrorBus {
	return &ErrorBus{channel: make(chan error, 1)}
}

func (b *ErrorBus) Report(source string, err error) {
	b.channel <- fmt.Errorf("reconciliation failed: %v", err)
}

func (b *ErrorBus) Close() {
	close(b.channel)
}

func (b *ErrorBus) Channel() <-chan error {
	return b.channel
}
