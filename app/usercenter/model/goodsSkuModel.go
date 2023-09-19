package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ GoodsSkuModel = (*customGoodsSkuModel)(nil)

type (
	// GoodsSkuModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGoodsSkuModel.
	GoodsSkuModel interface {
		goodsSkuModel
	}

	customGoodsSkuModel struct {
		*defaultGoodsSkuModel
	}
)

// NewGoodsSkuModel returns a model for the database table.
func NewGoodsSkuModel(conn sqlx.SqlConn) GoodsSkuModel {
	return &customGoodsSkuModel{
		defaultGoodsSkuModel: newGoodsSkuModel(conn),
	}
}
