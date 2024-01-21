package web

import (
	"codnect.io/procyon-core/component"
	"codnect.io/procyon-web/http/server"
)

type Module struct {
}

func (m Module) InitModule() {
	component.Register(server.New, component.Name("defaultWebServer"))
	component.Register(newServerLifecycle, component.Name("httpServerLifecycle"))
}
