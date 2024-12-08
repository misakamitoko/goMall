package main

import (
	"api-gateway/internal/config"
	"api-gateway/internal/svc"
	"fmt"

	"github.com/valyala/fasthttp"
)

func main() {
	var conf config.Config
	config.LoadConfig(&conf, "etc/config.yaml")
	fmt.Println(conf.Routes)
	ctx := svc.NewServiceContext(conf)
	gateway := NewGateWay(&conf, ctx)
	gateway.DiscoveryAllService()
	fasthttp.ListenAndServe(":8080", gateway.ServeHttp)
}
