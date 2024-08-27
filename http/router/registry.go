package router

import (
	"codnect.io/procyon-web/http"
	"codnect.io/procyon-web/http/middleware"
)

type MappingRegistry struct {
}

/*type SimpleMappingRegistry struct {
	//routes         []*tree
	returnValueHandlers []http.ReturnValueHandler
}*/

func newMappingRegistry(returnValueHandlers []http.ReturnValueHandler) *MappingRegistry {
	return nil
}

func (r *MappingRegistry) GetHandler(ctx http.Context) (http.HandlerChain, bool) {
	return nil, false
}
func (r *MappingRegistry) Handlers() map[*Mapping]any {
	return nil
}

func (r *MappingRegistry) Register(mapping *Mapping, handler http.Handler, middlewares ...*middleware.Middleware) {
}

func (r *MappingRegistry) Unregister(mapping *Mapping) {

}

type Registry interface {
	Handler(ctx http.Context) (http.HandlerChain, bool)
	Register(mapping Mapping, handler http.Handler) error
	Unregister(mapping Mapping)
}

type SimpleRegistry struct {
	tree []*routingTree
}

func NewSimpleRegistry() *SimpleRegistry {
	registry := &SimpleRegistry{
		make([]*routingTree, 9),
	}

	registry.createTree(http.MethodGet)
	registry.createTree(http.MethodHead)
	registry.createTree(http.MethodPost)
	registry.createTree(http.MethodPut)
	registry.createTree(http.MethodPatch)
	registry.createTree(http.MethodDelete)
	registry.createTree(http.MethodConnect)
	registry.createTree(http.MethodOptions)
	registry.createTree(http.MethodTrace)
	return registry
}

func (r *SimpleRegistry) createTree(method http.Method) {
	r.tree[method.IntValue()] = &routingTree{}
}

func (r *SimpleRegistry) Handler(ctx http.Context) (http.HandlerChain, bool) {
	return r.findHandler(ctx)
}

func (r *SimpleRegistry) Register(mapping Mapping, handler http.Handler) error {
	methods := mapping.Methods()

	for _, method := range methods {
		methodTree := r.tree[method.IntValue()]

		if methodTree.children == nil {
			methodTree.children = &routingNode{}
		}

		methodTree.add(mapping, nil)
	}

	return nil
}

func (r *SimpleRegistry) Unregister(mapping Mapping) {

}

func (r *SimpleRegistry) findHandler(ctx http.Context) (http.HandlerChain, bool) {
	var (
		request = ctx.Request()
		path    = request.Path()
		method  = request.Method()
	)

	methodTree := r.tree[method.IntValue()]

	if methodTree.staticRoutes != nil {
		if _, ok := methodTree.staticRoutes[path]; ok {
			return nil, true
		} else {
			return nil, false
		}
	}

	chain := methodTree.match(ctx)
	return chain, true
}
