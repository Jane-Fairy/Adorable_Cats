package handler

import (
	"net/http"

	"CatRoom/app/api/catapi/internal/logic"
	"CatRoom/app/api/catapi/internal/svc"
	"CatRoom/app/api/catapi/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CatapiHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewCatapiLogic(r.Context(), svcCtx)
		resp, err := l.Catapi(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
