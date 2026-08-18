package windowcache

func Normalize(values map[string]int) map[string]int {
	for k, v := range values {
		if v < 0 {
			values[k] = 0
		}
	}
	return values
}
