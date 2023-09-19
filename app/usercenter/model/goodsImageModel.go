package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ GoodsImageModel = (*customGoodsImageModel)(nil)

type (
	// GoodsImageModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGoodsImageModel.
	GoodsImageModel interface {
		goodsImageModel
	}

	customGoodsImageModel struct {
		*defaultGoodsImageModel
	}
)

// NewGoodsImageModel returns a model for the database table.
func NewGoodsImageModel(conn sqlx.SqlConn) GoodsImageModel {
	return &customGoodsImageModel{
		defaultGoodsImageModel: newGoodsImageModel(conn),
	}
}
