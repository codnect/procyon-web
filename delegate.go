package web

import "codnect.io/procyon-web/http"

type defaultRequestDelegate struct {
	ctx *defaultHttpServerContext
}

func (d defaultRequestDelegate) Invoke(ctx http.Context) {
	d.ctx.Invoke(ctx)
}
