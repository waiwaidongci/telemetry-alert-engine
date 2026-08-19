package ruledefaults

type Config struct{ Rules map[string]string }

func DefaultConfig() Config { return Config{Rules: map[string]string{}} }
