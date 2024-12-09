package db

import (
	"context"
	"product-service/api/internal/config"
	"time"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

func NewMySql(config config.Config) sqlx.SqlConn {
	conn := sqlx.NewMysql(config.DB.DSN)
	db, err := conn.RawDB()
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = db.PingContext(ctx)
	if err != nil {
		panic(err)
	}
	db.SetConnMaxIdleTime(10 * time.Minute)
	return conn
}
