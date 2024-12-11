package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	DB         DB         `json:"DB"`
	EtcdConfig EtcdConfig `json:"Etcd"`
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
