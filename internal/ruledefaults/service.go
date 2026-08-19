package ruledefaults

type Service struct{ config Config }

func NewService(config Config) *Service { return &Service{config: config} }

func (s *Service) Apply(name, expression string) {
	if s.config.Rules == nil {
		s.config.Rules = map[string]string{}
	}
	s.config.Rules[name] = expression
}

func (s *Service) Rules() map[string]string { return s.config.Rules }
