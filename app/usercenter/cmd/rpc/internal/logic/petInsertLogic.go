package logic

import (
	"context"

	"catroom/app/usercenter/cmd/rpc/internal/svc"
	"catroom/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PetInsertLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPetInsertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PetInsertLogic {
	return &PetInsertLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PetInsertLogic) PetInsert(in *pb.PetReq) (*pb.PetResp, error) {
	// todo: add your logic here and delete this line

	return &pb.PetResp{}, nil
}
