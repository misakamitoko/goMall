package handler

import (
	"net/http"

	"cart-service/api/internal/logic"
	"cart-service/api/internal/svc"
	"cart-service/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AddCartItemHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddItemReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewAddCartItemLogic(r.Context(), svcCtx)
		resp, err := l.AddCartItem(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
