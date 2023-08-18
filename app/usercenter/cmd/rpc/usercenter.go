package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"CatRoom/app/usercenter/cmd/rpc/internal/config"
	"CatRoom/app/usercenter/cmd/rpc/internal/server"
	"CatRoom/app/usercenter/cmd/rpc/internal/svc"
	"CatRoom/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "app/usercenter/cmd/rpc/etc/usercenter.yaml", "the config file")

func main() {
	flag.Parse()

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
