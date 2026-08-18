package ruleconfig

type Config struct{ Rules map[string]float64 }

func LoadDefaults() Config { return Config{} }
