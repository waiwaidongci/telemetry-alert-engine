package escalationstate

import "fmt"

type Incident struct {
	ID      string
	State   State
	History []State
}

type Machine struct{}

func (Machine) Move(incident Incident, next State) (Incident, error) {
	if !CanTransition(incident.State, next) {
		return incident, fmt.Errorf("invalid escalation transition %s -> %s", incident.State, next)
	}
	incident.State = next
	if len(incident.History) == 0 {
		incident.History = []State{next}
	} else {
		incident.History[len(incident.History)-1] = next
	}
	return incident, nil
}
