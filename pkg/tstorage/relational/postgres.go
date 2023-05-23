package relational

import (
	"fmt"
	"gb-auth-gate/config"
	sConfig "gb-auth-gate/pkg/tstorage/config"
	_ "github.com/jackc/pgx/stdlib" // pgx driver
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sarulabs/di"
	"time"
)

func InitPsqlDB(cfg *sConfig.TStorageConfig) (*sqlx.DB, error) {
	connectionUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Relational.Postgres.Host,
		cfg.Relational.Postgres.Port,
		cfg.Relational.Postgres.User,
		cfg.Relational.Postgres.Password,
		cfg.Relational.Postgres.DBName,
		cfg.Relational.Postgres.SSLMode)
	database, err := sqlx.Connect(cfg.Relational.Postgres.PgDriver, connectionUrl)
	if err != nil {
		return nil, err
	}
	database.DB.SetConnMaxIdleTime(time.Duration(cfg.Relational.Postgres.ConnMaxIdleTime) * time.Second)
	database.DB.SetMaxOpenConns(cfg.Relational.Postgres.MaxOpenConns)
	return database, nil
}

func BuildPostgres(ctn di.Container) (interface{}, error) {
	cfg := ctn.Get("config").(*config.Config)

	connectionUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.StorageConfig.Relational.Postgres.Host,
		cfg.StorageConfig.Relational.Postgres.Port,
		cfg.StorageConfig.Relational.Postgres.User,
		cfg.StorageConfig.Relational.Postgres.Password,
		cfg.StorageConfig.Relational.Postgres.DBName,
		cfg.StorageConfig.Relational.Postgres.SSLMode)
	database, err := sqlx.Connect(cfg.StorageConfig.Relational.Postgres.PgDriver, connectionUrl)
	if err != nil {
		return nil, err
	}
	database.DB.SetConnMaxIdleTime(time.Duration(cfg.StorageConfig.Relational.Postgres.ConnMaxIdleTime) * time.Second)
	database.DB.SetMaxOpenConns(cfg.StorageConfig.Relational.Postgres.MaxOpenConns)
	return database, nil
}
