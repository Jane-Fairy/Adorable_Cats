package logic

import (
	"catroom/app/usercenter/model"
	"catroom/common/xerr"
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

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
	petInfo, err := l.svcCtx.PetModel.FindOne(l.ctx, in.Id)

	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "mobile:%s,err:%v", in.Id, err)
	}
	petResp := new(pb.PetResp)
	_ = copier.Copy(&petResp, petInfo)

	petResp.CreateString = petInfo.CreateTime.String()
	petResp.UpdateString = petInfo.UpdateTime.String()
	petResp.DeleteString = petInfo.DeleteTime.String()
	petResp.DateOfBirth = petInfo.DateOfBirth.Time.String()
	return petResp, nil
}
