package main

import (
	"api-gateway/internal/handler/userhandler"
	"api-gateway/internal/svc"
	"context"
	"log"

	routing "github.com/qiangxue/fasthttp-routing"
)

func Route(r *routing.Router, sctx *svc.ServiceContext, c context.Context) {
	userApi := r.Group("/api/user")
	userApi.Get("/test", func(ctx *routing.Context) error {
		log.Fatal(string(ctx.Request.Body()))
		ctx.Response.SetBody([]byte("test"))
		return nil
	})
	userApi.Post("/login", userhandler.LoginHandler(sctx, c))
}
