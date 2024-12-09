package svc

import (
	"product-service/api/internal/config"
	"product-service/api/internal/db"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	Conn   *sqlx.SqlConn
}

func NewServiceContext(c config.Config) *ServiceContext {
	Conn := db.NewMySql(c)
	return &ServiceContext{
		Config: c,
		Conn:   &Conn,
	}
}
