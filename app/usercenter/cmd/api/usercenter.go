package main

import (
	_ "net/http/pprof"

	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/zeromicro/go-zero/core/logx"

	"catroom/app/usercenter/cmd/api/internal/config"
	"catroom/app/usercenter/cmd/api/internal/handler"
	"catroom/app/usercenter/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "../app/usercenter/cmd/api/etc/usercenter.yaml", "the config file")

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

	marshal, err := json.Marshal(c)
	log.Println(string(marshal))

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	logx.DisableStat()
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)

	//fmt.Println(http.ListenAndServe(":6060", nil)) //性能分析

	server.Start()
}
