package devicequery

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("device not found")

func RepositoryError(err error) error { return fmt.Errorf("load device: %v", err) }
