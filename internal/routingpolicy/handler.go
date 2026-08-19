package routingpolicy

type Handler struct {
	service   *Service
	validator Validator
}

func NewHandler(service *Service, validator Validator) *Handler {
	return &Handler{service: service, validator: validator}
}

func (h *Handler) Put(name, endpoint string) (Config, error) {
	config := h.service.Apply(name, endpoint)
	if h.validator == nil {
		return config, nil
	}
	if err := h.validator.Validate(config); err != nil {
		return Config{}, err
	}
	return config, nil
}
