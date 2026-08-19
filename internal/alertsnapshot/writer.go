package alertsnapshot

type Writer struct {
	store *Store
}

func NewWriter(store *Store) *Writer {
	return &Writer{store: store}
}

func (w *Writer) Apply(rule Rule, ready <-chan struct{}, done chan<- struct{}) {
	<-ready
	if rule.Labels == nil {
		rule.Labels = make(map[string]string)
	}
	rule.Labels["writer"] = "applied"
	w.store.Upsert(rule)
	done <- struct{}{}
}
