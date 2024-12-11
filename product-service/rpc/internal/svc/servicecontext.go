package svc

import (
	"product-service/rpc/internal/config"
	"product-service/rpc/internal/db"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config      config.Config
	Conn        *sqlx.SqlConn
	RedisClient *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	Conn := db.NewMySql(c)
	redisClient := redis.MustNewRedis(c.RedisConfig)
	return &ServiceContext{
		Config:      c,
		Conn:        &Conn,
		RedisClient: redisClient,
	}
}
