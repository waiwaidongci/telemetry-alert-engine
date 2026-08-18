package ruleconfig

func Apply(c *Config, name string, value float64) {
	if c.Rules == nil {
		c.Rules = make(map[string]float64)
	}
	c.Rules[name] = value
}
