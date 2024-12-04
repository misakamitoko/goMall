package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	Jwt Jwt `json:"jwt"`
}

type Jwt struct {
	SecretKey string
	Expire    int64
}
