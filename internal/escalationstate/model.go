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
	StateFailed:   {StateRetrying: {}, StateDelivered: {}},
	StateRetrying: {StateFailed: {}},
}

func CanTransition(from, to State) bool {
	allowed, ok := transitions[from]
	if !ok {
		return false
	}
	_, exists := allowed[to]
	return ok && exists
}
