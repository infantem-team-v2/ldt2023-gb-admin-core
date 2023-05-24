package accountInterface

import "gb-auth-gate/internal/account/model"

type RelationalRepository interface {
	GetPersonalUserInfo(userId int64) (*model.UserDAO, error)
	GetBusinessInfo(userId int64) (*model.BusinessDAO, error)
}
