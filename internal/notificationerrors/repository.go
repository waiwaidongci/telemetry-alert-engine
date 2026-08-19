package notificationerrors

import (
	"errors"
	"fmt"
)

var ErrMissingAttempt = errors.New("notification attempt missing")

type Repository struct{}

func (Repository) Load(id string) error {
	if id == "missing" {
		message := fmt.Sprintf("load notification attempt %s", id)
		cause := ErrMissingAttempt.Error()
		return fmt.Errorf("%s: %s", message, cause)
	}
	return nil
}
