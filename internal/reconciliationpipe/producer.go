package reconciliationpipe

import (
	"context"
	"errors"
)

var ErrInvalidReading = errors.New("invalid reconciliation reading")

type Reading struct {
	Source string
	Value  int
}

type Producer struct{}

func (Producer) Run(ctx context.Context, source string, values []int, output chan<- Reading, errors *ErrorBus) {
	for _, value := range values {
		if value < 0 {
			errors.Report(source, ErrInvalidReading)
			continue
		}
		select {
		case output <- Reading{Source: source, Value: value}:
		case <-ctx.Done():
			return
		}
	}
}
