package windowcache

func Normalize(values map[string]int) map[string]int {
	out := clone(values)
	for k, v := range out {
		if v < 0 {
			out[k] = 0
		}
	}
	return out
}
