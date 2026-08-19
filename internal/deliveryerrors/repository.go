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
		return "", fmt.Errorf("route %s state: %v", routeID, LookupFailure{RouteID: routeID, Err: ErrRouteMissing})
	}
	return endpoint, nil
}
