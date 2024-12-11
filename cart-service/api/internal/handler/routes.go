// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package handler

import (
	"net/http"

	"cart-service/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/cart/addItem",
				Handler: AddCartItemHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/cart/emptyCart",
				Handler: EmptyCartHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/cart/getCart",
				Handler: GetCartHandler(serverCtx),
			},
		},
	)
}
