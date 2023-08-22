package svc

import (
	"catroom/app/usercenter/cmd/api/internal/config"
	"catroom/app/usercenter/cmd/rpc/usercenter"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        config.Config
	UserCenterRpc usercenter.Usercenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		UserCenterRpc: usercenter.NewUsercenter(zrpc.MustNewClient(c.UsercenterRpcConf)),
	}
}
