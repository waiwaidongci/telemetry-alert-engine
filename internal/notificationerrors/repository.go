package notificationerrors

import (
	"errors"
	"fmt"
)

var ErrMissingAttempt = errors.New("notification attempt missing")

type Repository struct{}

func (Repository) Load(id string) error {
	if id == "missing" {
		return fmt.Errorf("load notification attempt %s: %w", id, ErrMissingAttempt)
	}
	return nil
}
