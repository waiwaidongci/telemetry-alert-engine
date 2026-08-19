package deliveryerrors

import "fmt"

type Repository struct {
	routes map[string]string
}

func NewRepository(routes map[string]string) *Repository {
	return &Repository{routes: routes}
}

func (r *Repository) Lookup(routeID string) (string, error) {
	endpoint, ok := r.routes[routeID]
	if !ok {
		return "", LookupFailure{RouteID: routeID, Err: fmt.Errorf("route state: %w", ErrRouteMissing)}
	}
	return endpoint, nil
}
