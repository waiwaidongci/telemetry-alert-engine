package ruleconfig

func Apply(c *Config, name string, value float64) { c.Rules[name] = value }
