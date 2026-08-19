package sweepcontext

import "context"

type Repository struct {
	gateway *Gateway
	ctx     context.Context
}

func NewRepository(gateway *Gateway) *Repository {
	return &Repository{gateway: gateway}
}

func (r *Repository) Load(ctx context.Context) error {
	if r.ctx == nil {
		r.ctx = ctx
	}
	return r.gateway.Fetch(r.ctx)
}
