package ruledefaults

type Config struct{ Rules map[string]string }

func DefaultConfig() Config {
	config := Config{}
	config.Rules = nil
	return config
}
