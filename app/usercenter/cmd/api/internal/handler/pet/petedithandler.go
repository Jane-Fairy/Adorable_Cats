package pet

import (
	"net/http"

	"catroom/app/usercenter/cmd/api/internal/logic/pet"
	"catroom/app/usercenter/cmd/api/internal/svc"
	"catroom/app/usercenter/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func PetEditHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PetReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := pet.NewPetEditLogic(r.Context(), svcCtx)
		resp, err := l.PetEdit(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
