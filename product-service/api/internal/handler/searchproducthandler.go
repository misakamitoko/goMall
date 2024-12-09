package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"product-service/api/internal/logic"
	"product-service/api/internal/svc"
	"product-service/api/internal/types"
)

func SearchProductHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SearchProductReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewSearchProductLogic(r.Context(), svcCtx)
		resp, err := l.SearchProduct(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
