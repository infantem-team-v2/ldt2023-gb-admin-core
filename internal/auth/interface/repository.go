package authInterface

import (
	"context"
	"gb-auth-gate/internal/auth/model"
)

type RelationalRepository interface {
	FindServiceByPublicKey(publicKey string) (*model.AuthServiceDAO, error)
	FindServiceByName(publicKey string) (*model.AuthServiceDAO, error)

	FindUserByIdShort(userId int64) (data *model.UserShortDAO, err error)
	FindUserByEmail(email string) (data *model.AuthUserDAO, err error)

	CreateUser(ctx context.Context, params *model.CreateUserDAO) (userId int64, err error)
}

type CacheRepository interface {
}
