package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	DB          DB              `json:"DB"`
	RedisConfig redis.RedisConf `json:"RedisConfig"`
}

type DB struct {
	DSN            string `json:"dsn"`
	ConnectTimeout int64  `json:"connectTimeout"`
}
