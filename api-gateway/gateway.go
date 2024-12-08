package main

import (
	"api-gateway/internal/config"
	"api-gateway/internal/service"
	"api-gateway/internal/svc"
	"api-gateway/internal/trie"
	"context"
	"fmt"

	"github.com/valyala/fasthttp"
)

type Gateway struct {
	conf   *config.Config
	l      *svc.ServiceContext
	client *fasthttp.Client
}

var routeTrie *trie.PathTrie

func NewGateWay(conf *config.Config, l *svc.ServiceContext) *Gateway {
	return &Gateway{
		conf:   conf,
		l:      l,
		client: &fasthttp.Client{},
	}
}

func (g *Gateway) DiscoveryAllService() {
	if routeTrie == nil {
		routeTrie = trie.CreateNewPathTrie()
	}

	for _, route := range g.conf.Routes {
		for _, path := range route.Path {
			routeTrie.Insert(path, route.ServiceName)
		}
	}
}

func (g *Gateway) ServeHttp(ctx *fasthttp.RequestCtx) {
	path := string(ctx.Path())
	targetServer, ok := routeTrie.Search(path)
	service.NewEtcdService(context.Background(), g.l, targetServer)
	if !ok || targetServer == "" {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}
	if targetServer == "" {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}
	targetUri := service.GetOneNodeByParent(targetServer) + path

	// RedirectReqeust
	req := &ctx.Request
	resp := &ctx.Response
	req.SetRequestURI("http://" + targetUri)
	req.Header.SetMethod(string(ctx.Method()))
	req.Header.SetContentType("application/json")
	err := g.client.Do(req, resp)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadGateway)
		fmt.Fprintf(ctx, "Failed to proxy request: %v", err)
		return
	}
	ctx.SetStatusCode(resp.StatusCode())
	ctx.SetBody(resp.Body())
}
