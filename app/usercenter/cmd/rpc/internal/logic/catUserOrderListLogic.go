package logic

import (
	"context"

	"catroom/app/usercenter/cmd/rpc/internal/svc"
	"catroom/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CatUserOrderListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCatUserOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CatUserOrderListLogic {
	return &CatUserOrderListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// oder detail list info
func (l *CatUserOrderListLogic) CatUserOrderList(in *pb.UserCatOrderListReq) (*pb.UserCatOrderListResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UserCatOrderListResp{}, nil
}
