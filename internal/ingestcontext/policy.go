package ingestcontext

import "context"

type Request struct{ ctx context.Context }

func NewRequest(ctx context.Context) Request {
	if ctx == nil {
		ctx = context.Background()
	}
	return Request{ctx: ctx}
}

func (r Request) Context() context.Context { return r.ctx }
