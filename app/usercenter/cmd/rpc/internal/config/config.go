package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	DB      DateGroup
	JwtAuth Jwt
	Cache   cache.CacheConf
	zrpc.RpcServerConf
}

type Jwt struct {
	AccessSecret string
	AccessExpire int64
}

type DateGroup struct {
	DataSource string
}

type Gorm struct {
}
