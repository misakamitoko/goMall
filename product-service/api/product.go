package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"product-service/api/internal/config"
	"product-service/api/internal/handler"
	"product-service/api/internal/register"
	"product-service/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/product.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)

	r := register.NewRegister(&c)
	go r.CreateKeyWithParent()

	go func() {
		sig := <-sigs
		fmt.Printf("signal %v\n", sig)
		r.UnRegister()
		os.Exit(0)
	}()


	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}