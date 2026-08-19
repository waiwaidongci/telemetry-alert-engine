package escalationstate

type Query struct{}

func (Query) InProgress(incidents []Incident) []Incident {
	result := make([]Incident, 0)
	for _, incident := range incidents {
		if incident.State == StateQueued {
			result = append(result, incident)
		}
	}
	return result
}

func (Query) Delivered(incidents []Incident) []Incident {
	result := make([]Incident, 0)
	for _, incident := range incidents {
		if incident.State == StateDelivered || incident.State == StateRetrying {
			result = append(result, incident)
		}
	}
	return result
}
