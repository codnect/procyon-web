package main

import (
	"codnect.io/procyon-web"
	"codnect.io/procyon-web/http"
	"context"
	"log"
)

func main() {
	server := web.NewDefaultHttpServer(Handler{})
	log.Printf("starting server...")
	err := server.Start(context.Background())
	if err != nil {
		panic(err)
	}

	/*
		log.Printf("hello")

		registry := router.NewSimpleRegistry()
		mapping := router.NewMapping("/hello/:name")
		mapping2 := router.NewMapping("/hello/:test/name")
		err := registry.Register(mapping, nil)
		if err != nil {

		}
		err = registry.Register(mapping2, nil)
		if err != nil {

		}

		registry.Handler(nil)
	*/
}

type Handler struct {
}

func (h Handler) Invoke(ctx http.Context) (any, error) {
	req := ctx.Request()
	req.Cookie("")
	val, ok := req.QueryParameter("test")

	if ok {
		log.Printf("value %s", val)
	}

	resp := ctx.Response()
	resp.AddCookie(&http.Cookie{
		Name:  "test",
		Value: "val",
	})

	resp.Flush()

	return nil, nil
}
