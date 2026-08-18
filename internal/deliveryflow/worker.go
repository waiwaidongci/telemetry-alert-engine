package deliveryflow

func RetrySucceeded(s *Service, d *Delivery) bool { return s.Transition(d, Sent) }
