package pet

import (
	"context"

	"catroom/app/usercenter/cmd/api/internal/svc"
	"catroom/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PetDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPetDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PetDeleteLogic {
	return &PetDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PetDeleteLogic) PetDelete(req *types.PetReq) (resp int, err error) {
	// todo: add your logic here and delete this line

	return
}
