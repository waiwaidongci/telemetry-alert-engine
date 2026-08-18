package exportbatch

import "errors"

var ErrBusiness = errors.New("export rejected")

type Tx struct{ CommitErr error }

func (tx *Tx) Commit() error { return tx.CommitErr }
func Persist(tx *Tx, businessErr error) (err error) {
	defer func() { err = tx.Commit() }()
	return businessErr
}
