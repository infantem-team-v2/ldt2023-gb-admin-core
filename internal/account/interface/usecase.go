package accountInterface

import "gb-auth-gate/internal/account/model"

type UseCase interface {
	GetCommonInfo(userId int64) (*model.GetCommonInfoResponse, error)
	UpdateUserInfo(userId int64, params *model.UpdateUserInfoRequest) (*model.GetCommonInfoResponse, error)
	GetResultsByUser(userId int64) (interface{}, uint16, error)
}
