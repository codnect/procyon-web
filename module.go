package web

import (
	"codnect.io/procyon-core/component"
	"codnect.io/procyon-core/component/condition"
	"codnect.io/procyon-web/http"
)

type Module struct {
}

func (m Module) InitModule() {
	component.Register(NewDefaultHttpServer, component.Named("procyonDefaultHttpServer")).
		ConditionalOn(condition.OnMissingType[http.Server]())
	component.Register(newServerLifecycle, component.Named("procyonHttpServerLifecycle"))
}
