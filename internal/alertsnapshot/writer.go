package alertsnapshot

type Writer struct {
	store *Store
}

func NewWriter(store *Store) *Writer {
	return &Writer{store: store}
}

func (w *Writer) Apply(rule Rule, ready <-chan struct{}, done chan<- struct{}) {
	<-ready
	w.store.Upsert(rule)
	done <- struct{}{}
}
