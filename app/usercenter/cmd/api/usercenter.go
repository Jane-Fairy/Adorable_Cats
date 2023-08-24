package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"log"
	"os"
	"path/filepath"

	"catroom/app/usercenter/cmd/api/internal/config"
	"catroom/app/usercenter/cmd/api/internal/handler"
	"catroom/app/usercenter/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "app/usercenter/cmd/api/etc/usercenter.yaml", "the config file")

func main() {
	flag.Parse()

	filePath, err := os.Getwd()
	if err != nil {
		fmt.Println("can not get current file path：", err)
		return
	}

	x := 1 << 10
	fmt.Println("===", x)
	configFilePath := filepath.Join(filePath, *configFile)

	var c config.Config
	conf.MustLoad(configFilePath, &c)

	marshal, err := json.Marshal(c)
	log.Println(string(marshal))

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	logx.DisableStat()
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
