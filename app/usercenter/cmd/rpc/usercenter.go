package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"os"
	"path/filepath"

	"catroom/app/usercenter/cmd/rpc/internal/config"
	"catroom/app/usercenter/cmd/rpc/internal/server"
	"catroom/app/usercenter/cmd/rpc/internal/svc"
	"catroom/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "app/usercenter/cmd/rpc/etc/usercenter.yaml", "the config file")

func main() {
	flag.Parse()

	filePath, err := os.Getwd()
	if err != nil {
		fmt.Println("can not get current file path：", err)
		return
	}

	configFilePath := filepath.Join(filePath, *configFile)

	var c config.Config
	conf.MustLoad(configFilePath, &c)
	ctx := svc.NewServiceContext(c)

	logx.DisableStat()

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterUsercenterServer(grpcServer, server.NewUsercenterServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
