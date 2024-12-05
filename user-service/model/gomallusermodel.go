package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ GomallUserModel = (*customGomallUserModel)(nil)

type (
	// GomallUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGomallUserModel.
	GomallUserModel interface {
		gomallUserModel
		withSession(session sqlx.Session) GomallUserModel
	}

	customGomallUserModel struct {
		*defaultGomallUserModel
	}
)

// NewGomallUserModel returns a model for the database table.
func NewGomallUserModel(conn sqlx.SqlConn) GomallUserModel {
	return &customGomallUserModel{
		defaultGomallUserModel: newGomallUserModel(conn),
	}
}

func (m *customGomallUserModel) withSession(session sqlx.Session) GomallUserModel {
	return NewGomallUserModel(sqlx.NewSqlConnFromSession(session))
}
