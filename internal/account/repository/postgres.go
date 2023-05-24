package repository

import (
	"context"
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
				bs.inn, bs.name, bs.economic_activity, bs.website
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

func (p *PostgresRepository) UpdatePersonalInfo(ctx context.Context, params *model.UpdateUserDataDAO) error {
	tx, err := p.db.BeginTxx(ctx, &sql.TxOptions{
		Isolation: sql.LevelReadUncommitted,
		ReadOnly:  false,
	})
	if err != nil {
		return terrors.Raise(err, 300006)
	}
	defer tx.Rollback()

	query := `
			 UPDATE
				personal.user
			 SET
			    full_name = :full_name,
			    email = :email,
			    job_position = :job_position
			 WHERE
			     id = :user_id
			 RETURNING geo_position_id;
			 `
	namedQuery, err := tx.PrepareNamed(query)
	if err != nil {
		return terrors.Raise(err, 300007)
	}

	var geoId int64

	err = namedQuery.QueryRowx(params).Scan(&geoId)
	if err != nil {
		return terrors.Raise(err, 300008)
	}
	params.GeoId = geoId

	query = `
		 	UPDATE
				personal.geography
		 	SET 
		 	    city = :city,
				country = :country
		 	WHERE
		 	    id = :geo_id
			RETURNING id;
			
		 	`
	namedQuery, err = tx.PrepareNamed(query)
	if err != nil {
		return terrors.Raise(err, 300007)
	}
	err = namedQuery.QueryRowx(params).Scan(&geoId)
	if err != nil {
		return terrors.Raise(err, 300008)
	}

	query = `
			UPDATE
				personal.business
			SET
				inn = :inn,
				name = :name,
				economic_activity = :economic_activity,
				website = :website
			WHERE
			    user_id = :user_id
			RETURNING id;
			`
	namedQuery, err = tx.PrepareNamed(query)
	if err != nil {
		return terrors.Raise(err, 300007)
	}
	err = namedQuery.QueryRowx(params).Scan(&geoId)
	if err != nil {
		return terrors.Raise(err, 300008)
	}
	err = tx.Commit()
	if err != nil {
		return terrors.Raise(err, 300009)
	}
	return nil
}
