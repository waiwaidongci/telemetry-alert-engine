package exportlease

import "errors"

var ErrRejectedPayload = errors.New("export payload rejected")

type Item struct {
	ID     string
	Reject bool
}

type Batch struct {
	pool       *Pool
	repository *Repository
}

func NewBatch(pool *Pool, repository *Repository) *Batch {
	return &Batch{pool: pool, repository: repository}
}

func (b *Batch) process(item Item) (err error) {
	lease, err := b.pool.Acquire()
	if err != nil {
		return err
	}
	defer lease.Close()
	tx := b.repository.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()
	if item.Reject {
		return ErrRejectedPayload
	}
	tx.Add(item.ID)
	return tx.Commit()
}

func (b *Batch) Run(items []Item) error {
	for _, item := range items {
		if err := b.process(item); err != nil {
			return err
		}
	}
	return nil
}
