package logic

import (
	"context"

	"catroom/app/usercenter/cmd/rpc/internal/svc"
	"catroom/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CatOrderTradeStateUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCatOrderTradeStateUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CatOrderTradeStateUpdateLogic {
	return &CatOrderTradeStateUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// update cat order detail information
func (l *CatOrderTradeStateUpdateLogic) CatOrderTradeStateUpdate(in *pb.UpdateCatOrderTradeStateReq) (*pb.UpdateCatOrderTradeStateResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateCatOrderTradeStateResp{}, nil
}
