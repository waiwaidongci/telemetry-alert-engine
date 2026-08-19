package deliveryerrors

import "errors"

type Resolver struct {
	repository *Repository
}

func NewResolver(repository *Repository) *Resolver {
	return &Resolver{repository: repository}
}

func (r *Resolver) Resolve(routeIDs []string) error {
	errorsByRoute := make([]error, 0)
	for _, routeID := range routeIDs {
		if _, err := r.repository.Lookup(routeID); err != nil {
			errorsByRoute = append(errorsByRoute, err)
		}
	}
	return errors.Join(errorsByRoute...)
}

func ShouldRetry(err error) bool {
	return err != nil && !Permanent(err)
}
