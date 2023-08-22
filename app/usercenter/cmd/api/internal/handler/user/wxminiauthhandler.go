package user

import (
	"log"
	"net/http"

	"catroom/app/usercenter/cmd/api/internal/logic/user"
	"catroom/app/usercenter/cmd/api/internal/svc"
	"catroom/app/usercenter/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func WxMiniAuthHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {

	log.Println("hello world")
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.WXMiniAuthReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewWxMiniAuthLogic(r.Context(), svcCtx)
		resp, err := l.WxMiniAuth(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
