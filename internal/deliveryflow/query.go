package deliveryflow

func Active(items []Delivery) []Delivery {
	out := make([]Delivery, 0, len(items))
	for _, d := range items {
		if d.Status == Queued || d.Status == Retrying {
			out = append(out, d)
		}
	}
	return out
}
