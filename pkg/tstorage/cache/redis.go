package cache

import (
	"fmt"
	mainConfig "gb-auth-gate/config"
	"gb-auth-gate/pkg/tutils/etc"
	"github.com/fiorix/go-redis/redis"
	"github.com/sarulabs/di"
)

func InitRedis(cfg *mainConfig.Config) (*redis.Client, error) {
	url := fmt.Sprintf(
		"%s:%s db=%d passwd=%s",
		cfg.StorageConfig.Cache.Redis.Host,
		cfg.StorageConfig.Cache.Redis.Port,
		etc.MustParseToInt(cfg.StorageConfig.Cache.Redis.DB),
		cfg.StorageConfig.Cache.Redis.Password)
	client, err := redis.NewClient(url)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func BuildRedis(ctn di.Container) (interface{}, error) {
	cfg := ctn.Get("config").(*mainConfig.Config)

	return InitRedis(cfg)
}
