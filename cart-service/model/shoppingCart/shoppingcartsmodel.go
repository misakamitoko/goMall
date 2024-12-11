package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ ShoppingCartsModel = (*customShoppingCartsModel)(nil)

type (
	// ShoppingCartsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customShoppingCartsModel.
	ShoppingCartsModel interface {
		shoppingCartsModel
		withSession(session sqlx.Session) ShoppingCartsModel
	}

	customShoppingCartsModel struct {
		*defaultShoppingCartsModel
	}
)

// NewShoppingCartsModel returns a model for the database table.
func NewShoppingCartsModel(conn sqlx.SqlConn) ShoppingCartsModel {
	return &customShoppingCartsModel{
		defaultShoppingCartsModel: newShoppingCartsModel(conn),
	}
}

func (m *customShoppingCartsModel) withSession(session sqlx.Session) ShoppingCartsModel {
	return NewShoppingCartsModel(sqlx.NewSqlConnFromSession(session))
}
