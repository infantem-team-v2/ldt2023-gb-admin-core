package accountInterface

import "gb-admin-core/internal/account/model"

type UseCase interface {
	GetCommonInfo(userId int64) (*model.GetCommonInfoResponse, error)
	UpdateUserInfo(userId int64, params *model.UpdateUserInfoRequest) (*model.GetCommonInfoResponse, error)
	GetResultsByUser(userId int64) (interface{}, uint16, error)
	UpdateConstants(params *model.ChangeConstantsRequest) (interface{}, uint16, error)
	InsertConstants(params *model.ChangeConstantsRequest) (interface{}, uint16, error)
	GetConstants() (interface{}, uint16, error)
}
