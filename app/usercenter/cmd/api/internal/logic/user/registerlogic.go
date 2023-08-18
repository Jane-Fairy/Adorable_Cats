package user

import (
	"CatRoom/app/usercenter/cmd/api/internal/svc"
	"CatRoom/app/usercenter/cmd/api/internal/types"
	"CatRoom/app/usercenter/cmd/rpc/usercenter"
	"CatRoom/app/usercenter/model"
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"log"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {

	log.Println("new account to register.")

	registerResp, err := l.svcCtx.UserCenterRpc.Register(l.ctx, &usercenter.RegisterReq{
		Mobile:   req.Mobile,
		Password: req.Password,
		AuthKey:  req.Mobile,
		AuthType: model.UserAuthTypeSystem,
	})

	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}

	var response types.RegisterResp
	_ = copier.Copy(&resp, registerResp)
	return &response, nil
}
