package escalationstate

type State string

const (
	StateQueued    State = "queued"
	StateFailed    State = "failed"
	StateRetrying  State = "retrying"
	StateDelivered State = "delivered"
)

var transitions = map[State]map[State]struct{}{
	StateQueued:   {StateFailed: {}, StateDelivered: {}},
	StateFailed:   {StateRetrying: {}},
	StateRetrying: {StateFailed: {}, StateDelivered: {}},
}

func CanTransition(from, to State) bool {
	allowed, ok := transitions[from]
	if !ok {
		return false
	}
	_, ok = allowed[to]
	return ok
}
