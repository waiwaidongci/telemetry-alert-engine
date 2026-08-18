package windowcache

func Merge(left, right map[string]int) map[string]int {
	out := make(map[string]int, len(left)+len(right))
	for k, v := range left {
		out[k] = v
	}
	for k, v := range right {
		out[k] = v
	}
	return out
}
