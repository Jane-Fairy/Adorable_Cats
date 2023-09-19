package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ GoodsUploadFileModel = (*customGoodsUploadFileModel)(nil)

type (
	// GoodsUploadFileModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGoodsUploadFileModel.
	GoodsUploadFileModel interface {
		goodsUploadFileModel
	}

	customGoodsUploadFileModel struct {
		*defaultGoodsUploadFileModel
	}
)

// NewGoodsUploadFileModel returns a model for the database table.
func NewGoodsUploadFileModel(conn sqlx.SqlConn) GoodsUploadFileModel {
	return &customGoodsUploadFileModel{
		defaultGoodsUploadFileModel: newGoodsUploadFileModel(conn),
	}
}
