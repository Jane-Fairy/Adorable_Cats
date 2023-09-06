package logic

import (
	"context"

	"catroom/app/usercenter/cmd/rpc/internal/svc"
	"catroom/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CatOrderDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCatOrderDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CatOrderDetailLogic {
	return &CatOrderDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// cat order detail information
func (l *CatOrderDetailLogic) CatOrderDetail(in *pb.CatOrderDetailReq) (*pb.CatOrderDetailResp, error) {
	// todo: add your logic here and delete this line

	return &pb.CatOrderDetailResp{}, nil
}
