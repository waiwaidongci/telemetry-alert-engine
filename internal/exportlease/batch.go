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

func (b *Batch) process(item Item) (lease *Lease, err error) {
	lease, err = b.pool.Acquire()
	if err != nil {
		return nil, err
	}
	tx := b.repository.Begin()
	defer func() {
		err = tx.Commit()
	}()
	tx.Add(item.ID)
	if item.Reject {
		return lease, ErrRejectedPayload
	}
	return lease, nil
}

func (b *Batch) Run(items []Item) error {
	for _, item := range items {
		lease, err := b.process(item)
		if lease != nil {
			defer lease.Close()
		}
		if err != nil {
			return err
		}
	}
	return nil
}
