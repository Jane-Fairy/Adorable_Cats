package logic

import (
	"catroom/app/usercenter/cmd/rpc/usercenter"
	"catroom/app/usercenter/model"
	"catroom/common/xerr"
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"catroom/app/usercenter/cmd/rpc/internal/svc"
	"catroom/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {
	// todo: add your logic here and delete this line
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "GetUserInfo find user db err , id:%d , err:%v", in.Id, err)
	}
	if user == nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.UserNotExitsError), "id:%d", in.Id)
	}
	var respUser usercenter.User
	_ = copier.Copy(&respUser, user)

	return &usercenter.GetUserInfoResp{
		User: &respUser,
	}, nil
}
