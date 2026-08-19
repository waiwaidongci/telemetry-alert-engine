package routingpolicy

type Config struct {
	Routes map[string]string
	Labels map[string]string
}

func DefaultConfig() Config {
	return Config{}
}

func Normalize(config Config) Config {
	if config.Routes == nil {
		config.Routes = make(map[string]string)
	}
	return config
}
