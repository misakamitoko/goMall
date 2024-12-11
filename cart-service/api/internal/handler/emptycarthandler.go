package handler

import (
	"net/http"

	"cart-service/api/internal/logic"
	"cart-service/api/internal/svc"
	"cart-service/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func EmptyCartHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EmptyCartReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewEmptyCartLogic(r.Context(), svcCtx)
		resp, err := l.EmptyCart(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
