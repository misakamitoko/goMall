package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	Auth Auth `json:"auth"`
}

type Auth struct {
	SecretKey string
	Expire    int64
}
