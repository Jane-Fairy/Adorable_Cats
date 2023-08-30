package pet

import (
	"context"

	"catroom/app/usercenter/cmd/api/internal/svc"
	"catroom/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PetEditLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPetEditLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PetEditLogic {
	return &PetEditLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PetEditLogic) PetEdit(req *types.PetReq) (resp int, err error) {
	// todo: add your logic here and delete this line

	return
}
