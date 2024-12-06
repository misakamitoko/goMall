package main

import (
	"api-gateway/internal/config"
	"api-gateway/internal/svc"
	"context"

	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

func main() {
	var conf config.Config
	config.LoadConfig(&conf, "etc/config.yaml")
	ctx := svc.NewServiceContext(conf)
	router := routing.New()

	Route(router, ctx, context.Background())
	fasthttp.ListenAndServe(":8080", router.HandleRequest)
}
