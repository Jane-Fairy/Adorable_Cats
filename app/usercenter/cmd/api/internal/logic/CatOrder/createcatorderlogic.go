package CatOrder

import (
	"context"

	"catroom/app/usercenter/cmd/api/internal/svc"
	"catroom/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCatOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateCatOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCatOrderLogic {
	return &CreateCatOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateCatOrderLogic) CreateCatOrder(req *types.CreateCatOrderReq) (resp *types.CreateCatOrderResp, err error) {
	// todo: add your logic here and delete this line

	return
}
