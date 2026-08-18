package exportbatch

type Tracker struct{ Open, Peak int }
type Resource struct {
	t      *Tracker
	closed bool
}

func (t *Tracker) Acquire() *Resource {
	t.Open++
	if t.Open > t.Peak {
		t.Peak = t.Open
	}
	return &Resource{t: t}
}
func (r *Resource) Close() {
	if r.closed {
		return
	}
	r.closed = true
	r.t.Open--
}
