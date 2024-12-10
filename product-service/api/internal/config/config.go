package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	DB          DB              `json:"DB"`
	EtcdConfig  EtcdConfig      `json:"Etcd"`
	RedisConfig redis.RedisConf `json:"RedisConfig"`
	RedisPrefix string          `json:"RedisPrefix"`
}

type DB struct {
	DSN            string `json:"dsn"`
	ConnectTimeout int64  `json:"connectTimeout"`
}

type EtcdConfig struct {
	Hosts       []string
	Key         string
	DialTimeout uint32
}
