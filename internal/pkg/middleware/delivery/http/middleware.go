package http

import (
	"fmt"
	"gb-auth-gate/config"
	authInterface "gb-auth-gate/internal/auth/interface"
	"gb-auth-gate/internal/pkg/middleware/model"
	"gb-auth-gate/pkg/tlogger"
	"gb-auth-gate/pkg/tutils/etc"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/redis/v2"
	"github.com/sarulabs/di"
	"time"
)

type MiddlewareManager struct {
	sessionStorage *session.Store
	Config         *config.Config        `di:"config"`
	Logger         tlogger.ILogger       `di:"logger"`
	AuthUC         authInterface.UseCase `di:"authUC"`
}

func BuildMiddlewareManager(ctn di.Container) (interface{}, error) {
	cfg := ctn.Get("config").(*config.Config)
	logger := ctn.Get("logger").(tlogger.ILogger)
	authUC := ctn.Get("authUC").(authInterface.UseCase)

	s := session.New(session.Config{
		CookieHTTPOnly: true,
		Storage: redis.New(redis.Config{
			Reset:    false,
			Host:     cfg.StorageConfig.Cache.Redis.Host,
			Port:     etc.MustParseToInt(cfg.StorageConfig.Cache.Redis.Port),
			Password: cfg.StorageConfig.Cache.Redis.Password,
			Database: etc.MustParseToInt(cfg.StorageConfig.Cache.Redis.DB),
		}),
		Expiration:   time.Duration(cfg.MiddlewareConfig.Auth.ExpirationTime) * time.Minute,
		KeyLookup:    fmt.Sprintf("cookie:%s", model.SessionKey),
		CookieDomain: cfg.BaseConfig.Service.URL,
		CookieSecure: true,
	})

	return &MiddlewareManager{
		sessionStorage: s,
		Config:         cfg,
		Logger:         logger,
		AuthUC:         authUC,
	}, nil
}
