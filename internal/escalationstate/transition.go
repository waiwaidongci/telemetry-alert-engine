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
	incident.History = append(append([]State(nil), incident.History...), next)
	return incident, nil
}
