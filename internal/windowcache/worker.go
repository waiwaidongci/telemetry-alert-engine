package windowcache

func Merge(left, right map[string]int) map[string]int {
	for k, v := range right {
		left[k] = v
	}
	return left
}
