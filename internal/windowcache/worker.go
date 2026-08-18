package windowcache

func Merge(left, right map[string]int) map[string]int {
	out := clone(left)
	for k, v := range right {
		out[k] = v
	}
	return out
}
