package web

import "codnect.io/procyon-web/http"

type ServerContextDelegate struct {
	ctx *ServerContext
}

func (d ServerContextDelegate) Invoke(ctx http.Context) {
	d.ctx.Invoke(ctx)
}
