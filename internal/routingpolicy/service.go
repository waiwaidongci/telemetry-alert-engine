package routingpolicy

type Service struct {
	config Config
}

func NewService(config Config) *Service {
	return &Service{config: Normalize(config)}
}

func (s *Service) Apply(name, endpoint string) Config {
	s.config = Normalize(s.config)
	s.config.Routes[name] = endpoint
	s.config.Labels["last_route"] = name
	return Normalize(s.config)
}
