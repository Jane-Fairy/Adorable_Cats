package pet

import (
	"context"

	"catroom/app/usercenter/cmd/api/internal/svc"
	"catroom/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindPetOneLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindPetOneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindPetOneLogic {
	return &FindPetOneLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindPetOneLogic) FindPetOne(req *types.PetReq) (resp *types.PetResp, err error) {
	// todo: add your logic here and delete this line
	
	return
}
