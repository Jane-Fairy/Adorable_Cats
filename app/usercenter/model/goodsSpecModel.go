package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ GoodsSpecModel = (*customGoodsSpecModel)(nil)

type (
	// GoodsSpecModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGoodsSpecModel.
	GoodsSpecModel interface {
		goodsSpecModel
	}

	customGoodsSpecModel struct {
		*defaultGoodsSpecModel
	}
)

// NewGoodsSpecModel returns a model for the database table.
func NewGoodsSpecModel(conn sqlx.SqlConn) GoodsSpecModel {
	return &customGoodsSpecModel{
		defaultGoodsSpecModel: newGoodsSpecModel(conn),
	}
}
