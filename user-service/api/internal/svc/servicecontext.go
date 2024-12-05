package svc

import (
	"auth-service/auth"
	"user-service/api/db"
	"user-service/api/internal/config"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	AuthRpc     auth.AuthServiceClient
	RedisClient *redis.Redis
	Conn        sqlx.SqlConn
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := db.NewMySql(c)
	var clientConf zrpc.RpcClientConf
	conf.MustLoad("etc/client.yaml", &clientConf)
	conn := zrpc.MustNewClient(clientConf)
	redisClient := redis.MustNewRedis(c.RedisConfig)
	return &ServiceContext{
		Config:      c,
		AuthRpc:     auth.NewAuthServiceClient(conn.Conn()),
		RedisClient: redisClient,
		Conn:        sqlConn,
	}

}
