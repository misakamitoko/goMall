package svc

import (
	"cart-service/api/internal/config"
	"cart-service/api/internal/db"
	"cart-service/api/internal/product"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        config.Config
	ProductClient product.ProductCatalogServiceClient
	Conn          *sqlx.SqlConn
}

func NewServiceContext(c config.Config) *ServiceContext {
	DBConn := db.NewMySql(c)
	var clientConf zrpc.RpcClientConf
	conf.MustLoad("etc/client.yaml", &clientConf)
	conn := zrpc.MustNewClient(clientConf)
	return &ServiceContext{
		Config:        c,
		ProductClient: product.NewProductCatalogServiceClient(conn.Conn()),
		Conn:          &DBConn,
	}
}
