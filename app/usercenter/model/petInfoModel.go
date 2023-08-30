package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ PetInfoModel = (*customPetInfoModel)(nil)

type (
	// PetInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPetInfoModel.
	PetInfoModel interface {
		petInfoModel
	}

	customPetInfoModel struct {
		*defaultPetInfoModel
	}
)

// NewPetInfoModel returns a model for the database table.
func NewPetInfoModel(conn sqlx.SqlConn) PetInfoModel {
	return &customPetInfoModel{
		defaultPetInfoModel: newPetInfoModel(conn),
	}
}
