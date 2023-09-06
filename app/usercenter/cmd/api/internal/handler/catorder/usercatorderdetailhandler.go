package catorder

import (
	"net/http"

	"catroom/app/usercenter/cmd/api/internal/logic/CatOrder"
	"catroom/app/usercenter/cmd/api/internal/svc"
	"catroom/app/usercenter/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserCatOrderDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserCatOrderDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := CatOrder.NewUserCatOrderDetailLogic(r.Context(), svcCtx)
		resp, err := l.UserCatOrderDetail(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
