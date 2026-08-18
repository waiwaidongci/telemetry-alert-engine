package deliveryflow

type Service struct{ transitions map[Status]map[Status]bool }

func NewService() *Service {
	return &Service{transitions: map[Status]map[Status]bool{
		Queued: {Retrying: true, Failed: true}, Retrying: {Failed: true}, Failed: {Retrying: true}, Sent: {},
	}}
}

func (s *Service) Transition(d *Delivery, next Status) bool {
	if !s.transitions[d.Status][next] {
		return false
	}
	d.Status = next
	return true
}
