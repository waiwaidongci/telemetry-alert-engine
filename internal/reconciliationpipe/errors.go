package reconciliationpipe

import "fmt"

type ErrorBus struct {
	channel chan error
}

func NewErrorBus(capacity int) *ErrorBus {
	if capacity < 1 {
		capacity = 1
	}
	return &ErrorBus{channel: make(chan error, capacity)}
}

func (b *ErrorBus) Report(source string, err error) {
	b.channel <- fmt.Errorf("source %s: %w", source, err)
}

func (b *ErrorBus) Close() {
	close(b.channel)
}

func (b *ErrorBus) Channel() <-chan error {
	return b.channel
}
