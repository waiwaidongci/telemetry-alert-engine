package sweepcontext

import "context"

type Repository struct {
	gateway *Gateway
}

func NewRepository(gateway *Gateway) *Repository {
	return &Repository{gateway: gateway}
}

func (r *Repository) Load(ctx context.Context) error {
	return r.gateway.Fetch(ctx)
}
