package exportbatch

func ExportMany(t *Tracker, count int) {
	for i := 0; i < count; i++ {
		func() { r := t.Acquire(); defer r.Close() }()
	}
}
