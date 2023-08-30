package logic

import (
	"context"

	"catroom/app/usercenter/cmd/rpc/internal/svc"
	"catroom/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindPetOneLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindPetOneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindPetOneLogic {
	return &FindPetOneLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindPetOneLogic) FindPetOne(in *pb.PetReq) (*pb.PetResp, error) {
	// todo: add your logic here and delete this line

	return &pb.PetResp{}, nil
}
