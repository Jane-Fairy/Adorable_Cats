package logic

import (
	"context"

	"catroom/app/usercenter/cmd/rpc/internal/svc"
	"catroom/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PetDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPetDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PetDeleteLogic {
	return &PetDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PetDeleteLogic) PetDelete(in *pb.PetReq) (*pb.StatusInfo, error) {
	// todo: add your logic here and delete this line

	return &pb.StatusInfo{}, nil
}
