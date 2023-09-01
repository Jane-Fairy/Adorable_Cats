package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	DB struct {
		DataSource string
	}
	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}
	Cache cache.CacheConf
	zrpc.RpcServerConf
	Gorm GormConfig
}

type GormConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
	Charset  string
}
