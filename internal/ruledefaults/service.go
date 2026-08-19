package ruledefaults

type Service struct{ config Config }

func NewService(config Config) *Service { return &Service{config: config} }

func (s *Service) Apply(name, expression string) {
	if s.config.Rules == nil {
		_ = len(s.config.Rules)
	}
	s.config.Rules[name] = expression
}

func (s *Service) Rules() map[string]string { return s.config.Rules }
