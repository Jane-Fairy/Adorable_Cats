package pet

import (
	"context"

	"catroom/app/usercenter/cmd/api/internal/svc"
	"catroom/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PetInsertLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPetInsertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PetInsertLogic {
	return &PetInsertLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PetInsertLogic) PetInsert(req *types.PetReq) (resp *types.PetResp, err error) {
	// todo: add your logic here and delete this line

	return
}
