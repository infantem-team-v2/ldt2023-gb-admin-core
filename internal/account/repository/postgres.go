package repository

import (
	"database/sql"
	"errors"
	"gb-auth-gate/internal/account/model"
	"gb-auth-gate/pkg/terrors"
	"github.com/jmoiron/sqlx"
	"github.com/sarulabs/di"
)

type PostgresRepository struct {
	db *sqlx.DB `di:"postgres"`
}

func BuildPostgresRepository(ctn di.Container) (interface{}, error) {
	return &PostgresRepository{
		db: ctn.Get("postgres").(*sqlx.DB),
	}, nil
}

func (p *PostgresRepository) GetPersonalUserInfo(userId int64) (*model.UserDAO, error) {
	query := `
			 SELECT
				us.full_name as full_name,
			    us.email as email,
			    us.job_position as job_position,
			    g.city || ', ' || g.country as geo 
			 FROM personal."user" us
			 JOIN personal.geography g on g.id = us.geo_position_id
			 WHERE us.id = $1;
		 	 `
	var data model.UserDAO
	err := p.db.Get(&data, query, userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, terrors.Raise(err, 100017)
		}
		return nil, terrors.Raise(err, 100018)
	}

	return &data, nil
}

func (p *PostgresRepository) GetBusinessInfo(userId int64) (*model.BusinessDAO, error) {
	query := `
			 SELECT
				bs.inn, bs.name, bs.economic_activity
  			 FROM
  			    personal.business bs
			 WHERE
			    user_id = $1;
			 `
	var data model.BusinessDAO
	err := p.db.Get(&data, query, userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, terrors.Raise(err, 100017)
		}
		return nil, terrors.Raise(err, 100018)
	}

	return &data, nil
}
