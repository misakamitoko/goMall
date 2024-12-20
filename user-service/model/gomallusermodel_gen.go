// Code generated by goctl. DO NOT EDIT.
// versions:
//  goctl version: 1.7.3

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	gomallUserFieldNames          = builder.RawFieldNames(&GomallUser{})
	gomallUserRows                = strings.Join(gomallUserFieldNames, ",")
	gomallUserRowsExpectAutoSet   = strings.Join(gomallUserFieldNames, ",")
	gomallUserRowsWithPlaceHolder = strings.Join(stringx.Remove(gomallUserFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	gomallUserModel interface {
		Insert(ctx context.Context, data *GomallUser) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*GomallUser, error)
		FindOneByEmail(ctx context.Context, email string) (*GomallUser, error)
		Update(ctx context.Context, data *GomallUser) error
		Delete(ctx context.Context, id int64) error
	}

	defaultGomallUserModel struct {
		conn  sqlx.SqlConn
		table string
	}

	GomallUser struct {
		Id          uint32     `db:"id"`
		Name        string    `db:"name"`         // 用户名
		Gender      int64     `db:"gender"`       // 用户性别
		Mobile      string    `db:"mobile"`       // 电话号码
		Email       string    `db:"email"`		  //电子邮箱
		Password    string    `db:"password"`     // 用户密码
		CreateTime  time.Time `db:"create_time"`  // 创建时间
		UpdateTime time.Time `db:"updated_time"` // 修改时间
		
	}
)

func newGomallUserModel(conn sqlx.SqlConn) *defaultGomallUserModel {
	return &defaultGomallUserModel{
		conn:  conn,
		table: "`gomall_user`",
	}
}

func (m *defaultGomallUserModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultGomallUserModel) FindOne(ctx context.Context, id int64) (*GomallUser, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", gomallUserRows, m.table)
	var resp GomallUser
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func(m *defaultGomallUserModel) FindOneByEmail(ctx context.Context, email string) (*GomallUser, error) {
	query := fmt.Sprintf("select %s from %s where email = ? limit 1", gomallUserRows, m.table)
	var resp GomallUser
	err := m.conn.QueryRowCtx(ctx, &resp, query, email)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultGomallUserModel) Insert(ctx context.Context, data *GomallUser) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, gomallUserRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Id, data.Name, data.Gender, data.Mobile, data.Email, data.Password, data.CreateTime, data.UpdateTime)
	return ret, err
}

func (m *defaultGomallUserModel) Update(ctx context.Context, data *GomallUser) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, gomallUserRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Name, data.Gender, data.Mobile, data.Password, data.UpdateTime, data.Id)
	return err
}

func (m *defaultGomallUserModel) tableName() string {
	return m.table
}
