package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	DB DB `json:"DB"`
}

type DB struct {
	DSN            string `json:"dsn"`
	ConnectTimeout int64  `json:"connectTimeout"`
}
