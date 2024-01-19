package server

import (
	"codnect.io/procyon-web/http"
)

type ContextDelegate struct {
	ctx *Context
}

func (d ContextDelegate) Invoke(ctx http.Context) {
	d.ctx.Invoke(ctx)
}
