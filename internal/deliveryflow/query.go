package deliveryflow

func Active(items []Delivery) []Delivery {
	out := make([]Delivery, 0, len(items))
	for _, d := range items {
		if !d.Status.Terminal() {
			out = append(out, d)
		}
	}
	return out
}
