package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	DB struct {
		DataSource string
	}
	Cache cache.CacheConf
	zrpc.RpcServerConf
}
