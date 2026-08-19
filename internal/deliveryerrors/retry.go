package deliveryerrors

import "errors"

type Resolver struct {
	repository *Repository
}

func NewResolver(repository *Repository) *Resolver {
	return &Resolver{repository: repository}
}

func (r *Resolver) Resolve(routeIDs []string) error {
	for _, routeID := range routeIDs {
		if _, err := r.repository.Lookup(routeID); err != nil {
			return errors.New(err.Error())
		}
	}
	return nil
}

func ShouldRetry(err error) bool {
	return err != nil
}
