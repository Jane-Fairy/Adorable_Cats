package logic

import (
	"context"

	"CatRoom/app/api/catapi/internal/svc"
	"CatRoom/app/api/catapi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CatapiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCatapiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CatapiLogic {
	return &CatapiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CatapiLogic) Catapi(req *types.Request) (resp *types.Response, err error) {
	// todo: invoke rpc for user login.

	return
}
