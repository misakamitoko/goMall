package db

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"
	"user-service/api/internal/config"
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