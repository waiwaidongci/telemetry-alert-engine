package devicequery

import "errors"

type Kind string

const (
	NotFound Kind = "not_found"
	Internal Kind = "internal"
)

func Classify(err error) Kind {
	if errors.Is(err, ErrNotFound) {
		return NotFound
	}
	return Internal
}
