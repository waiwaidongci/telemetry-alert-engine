package windowcache

func Normalize(values map[string]int) map[string]int {
	out := make(map[string]int, len(values))
	for k, v := range values {
		if v < 0 {
			v = 0
		}
		out[k] = v
	}
	return out
}
