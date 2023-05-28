package accountInterface

import (
	"context"
	"gb-admin-core/internal/account/model"
)

type RelationalRepository interface {
	GetPersonalUserInfo(userId int64) (*model.UserDAO, error)
	GetBusinessInfo(userId int64) (*model.BusinessDAO, error)

	UpdatePersonalInfo(ctx context.Context, params *model.UpdateUserDataDAO) error
}
