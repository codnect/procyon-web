package server

import (
	"codnect.io/procyon-web/http"
)

type RequestDelegate struct {
	ctx *Context
}

func (d RequestDelegate) Invoke(ctx http.Context) {
	d.ctx.Invoke(ctx)
}
