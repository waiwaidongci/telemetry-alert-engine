package notificationerrors

type Kind string

const (
	KindMissing Kind = "missing"
	KindSystem  Kind = "system"
)

func Classify(err error) Kind {
	if err == nil {
		return KindSystem
	}
	message := err.Error()
	if message == ErrMissingAttempt.Error() {
		return KindMissing
	}
	return KindSystem
}
