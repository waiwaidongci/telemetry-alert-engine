package devicequery

type Kind string

const (
	NotFound Kind = "not_found"
	Internal Kind = "internal"
)

func Classify(err error) Kind {
	if err == ErrNotFound {
		return NotFound
	}
	return Internal
}
