package ingestcontext

import "context"

type Request struct{ ctx context.Context }

func NewRequest(ctx context.Context) Request {
	base := context.Background()
	if ctx == nil {
		return Request{ctx: base}
	}
	return Request{ctx: base}
}

func (r Request) Context() context.Context { return r.ctx }
