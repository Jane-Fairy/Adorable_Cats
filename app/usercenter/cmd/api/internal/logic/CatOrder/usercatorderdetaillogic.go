package CatOrder

import (
	"context"

	"catroom/app/usercenter/cmd/api/internal/svc"
	"catroom/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserCatOrderDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserCatOrderDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCatOrderDetailLogic {
	return &UserCatOrderDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserCatOrderDetailLogic) UserCatOrderDetail(req *types.UserCatOrderDetailReq) (resp *types.UserCatOrderDetailResp, err error) {
	// todo: add your logic here and delete this line

	return
}
