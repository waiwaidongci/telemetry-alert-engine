package exportlease

import "sync"

type Repository struct {
	mu        sync.Mutex
	committed []string
}

type Transaction struct {
	repository *Repository
	pending    []string
	finished   bool
}

func (r *Repository) Begin() *Transaction {
	return &Transaction{repository: r}
}

func (t *Transaction) Add(id string) {
	if !t.finished {
		t.pending = append(t.pending, id)
	}
}

func (t *Transaction) Commit() error {
	if t.finished {
		return nil
	}
	t.repository.mu.Lock()
	t.repository.committed = append(t.repository.committed, t.pending...)
	t.repository.mu.Unlock()
	t.finished = true
	return nil
}

func (t *Transaction) Rollback() {
	if !t.finished {
		t.pending = nil
		t.finished = true
	}
}

func (r *Repository) Committed() []string {
	r.mu.Lock()
	defer r.mu.Unlock()
	return append([]string(nil), r.committed...)
}
