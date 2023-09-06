package CatOrder

import (
	"context"

	"catroom/app/usercenter/cmd/api/internal/svc"
	"catroom/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserCatOrderListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserCatOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCatOrderListLogic {
	return &UserCatOrderListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserCatOrderListLogic) UserCatOrderList(req *types.UserCatOrderListReq) (resp *types.UserCatOrderListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
