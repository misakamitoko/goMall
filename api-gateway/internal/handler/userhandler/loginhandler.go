package userhandler

import (
	"api-gateway/internal/service"
	"api-gateway/internal/svc"
	"context"
	"log"

	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

func LoginHandler(svcCtx *svc.ServiceContext, c context.Context) routing.Handler {
	etcdClient := service.NewEtcdService(c, svcCtx)
	return func(ctx *routing.Context) error {
		endpoints := etcdClient.GetByParent("user-api")
		client := &fasthttp.Client{}

		// 构造请求
		req := fasthttp.AcquireRequest()
		defer fasthttp.ReleaseRequest(req)

		req.SetRequestURI("http://" + endpoints[1] + "/api/user/login")
		req.Header.SetMethod("POST")
		req.Header.SetContentType("application/json") // 设置请求体类型
		req.SetBody(ctx.Request.Body())
		// 构造响应
		resp := fasthttp.AcquireResponse()
		defer fasthttp.ReleaseResponse(resp)

		// 发送请求
		err := client.Do(req, resp)
		if err != nil {
			log.Fatal(err)
		}
		ctx.Response.SetStatusCode(resp.StatusCode())
		ctx.Response.SetBody(resp.Body())
		return nil
	}
}
