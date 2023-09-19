package logic

import (
	"context"

	"catroom/app/usercenter/cmd/rpc/internal/svc"
	"catroom/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CatOrderCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCatOrderCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CatOrderCreateLogic {
	return &CatOrderCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Order rpc doc
func (l *CatOrderCreateLogic) CatOrderCreate(in *pb.CreateCatOrderReq) (*pb.CreateCatOrderResp, error) {
	// todo: add your logic here and delete this line
	
	return &pb.CreateCatOrderResp{}, nil
}
