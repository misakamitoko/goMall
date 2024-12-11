package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DB          DB              `json:"DB"`
	RedisConfig redis.RedisConf `json:"RedisConfig"`
	RedisPrefix string          `json:"RedisPrefix"`
}

type DB struct {
	DSN            string `json:"dsn"`
	ConnectTimeout int64  `json:"connectTimeout"`
}

type RedisConfig struct {
	Host        string `json:"host"`
	Type        string `json:"type"`
	NonBlock    bool   `json:"nonBlock"`
	Tls         bool   `json:"tls"`
	PingTimeout int    `json:"pingTimeout"`
}
