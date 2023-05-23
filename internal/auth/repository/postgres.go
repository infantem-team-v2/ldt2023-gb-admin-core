package repository

import (
	"context"
	"database/sql"
	"gb-auth-gate/internal/auth/model"
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

func (p PostgresRepository) FindServiceByPublicKey(publicKey string) (data *model.AuthServiceDAO, err error) {
	query := `
			 SELECT * FROM service.auth WHERE public_key = $1;
			 `
	data = &model.AuthServiceDAO{}
	err = p.db.Get(data, query, publicKey)
	if err != nil {
		return nil, terrors.Raise(err, 100002)
	}

	return data, nil
}

func (p PostgresRepository) FindServiceByName(name string) (data *model.AuthServiceDAO, err error) {
	query := `
			 SELECT * FROM service.auth WHERE name = $1;
			 `
	data = &model.AuthServiceDAO{}
	err = p.db.Get(data, query, name)
	if err != nil {
		return nil, terrors.Raise(err, 200002)
	}

	return data, nil
}

func (p PostgresRepository) FindUserByIdShort(userId int64) (data *model.UserShortDAO, err error) {
	query := `
			 SELECT id, full_name, email FROM personal."user" WHERE id = $1;
			 `
	data = &model.UserShortDAO{}
	err = p.db.Get(data, query, userId)
	if err != nil {
		return nil, terrors.Raise(err, 100010)
	}

	return data, nil
}

func (p PostgresRepository) FindUserByEmail(email string) (*model.AuthUserDAO, error) {
	query := `
			 SELECT id, email, password FROM personal.user WHERE email = $1;
			 `
	var data model.AuthUserDAO
	err := p.db.Get(&data, query, email)
	if err != nil {
		return nil, terrors.Raise(err, 100015)
	}

	return &data, err
}

func (p PostgresRepository) CreateUser(ctx context.Context, params *model.CreateUserDAO) (userId int64, err error) {
	tx, err := p.db.BeginTxx(ctx, &sql.TxOptions{
		Isolation: sql.LevelReadUncommitted,
		ReadOnly:  false,
	})
	if err != nil {
		return userId, terrors.Raise(err, 300006)
	}
	defer tx.Rollback()
	query := `
			 INSERT INTO personal."user"
			     (full_name, email, password, job_position) 
			 VALUES 
			 	(:full_name, :email, :password, :job_position)
			 RETURNING personal."user".id;
			 `
	namedQuery, err := tx.PrepareNamed(query)
	if err != nil {
		return userId, terrors.Raise(err, 300003)
	}
	err = namedQuery.QueryRowx(params).Scan(&userId)
	if err != nil {

		return userId, terrors.Raise(err, 300005)
	}
	params.UserId = &userId

	var (
		geoId      int64
		businessId int64
	)

	if params.City == nil && params.Country == nil {
		goto afterGeo
	}
	query = `
			SELECT 
			    id
			FROM
			    personal.geography
			WHERE
			    city = :city and country = :country;
			`
	namedQuery, err = tx.PrepareNamed(query)
	if err != nil {
		return userId, terrors.Raise(err, 300003)
	}
	err = namedQuery.QueryRowx(params).Scan(&geoId)
	if err != nil {
		if err != sql.ErrNoRows {
			return userId, terrors.Raise(err, 300005)
		}
	} else if err == nil && geoId == 0 {
		goto setGeo
	}

	query = `
			INSERT INTO personal.geography
				(city, country) 
			VALUES 
				(:city, :country)
			RETURNING personal.geography.id;
		  	`
	namedQuery, err = tx.PrepareNamed(query)
	if err != nil {
		return userId, terrors.Raise(err, 300003)
	}
	err = namedQuery.QueryRowx(params).Scan(&geoId)
	if err != nil {
		return userId, terrors.Raise(err, 300005)
	}

setGeo:
	query = `
			UPDATE
			    personal."user"
			SET
			    geo_position_id = $1
			WHERE
			    personal."user".id = $2
			RETURNING
			    geo_position_id;
			`
	err = tx.QueryRowx(query, geoId, userId).Scan(&geoId)
	if err != nil {
		return userId, terrors.Raise(err, 300005)
	}

afterGeo:

	query = `
			INSERT INTO
			    personal.business
			    (inn, name, economic_activity, website, user_id) 
			VALUES
			    (:inn, :business_name, :economic_activity, :website, :user_id) 
			RETURNING business.id;
			`
	namedQuery, err = tx.PrepareNamed(query)
	err = namedQuery.QueryRowx(params).Scan(&businessId)
	if err != nil {
		return userId, terrors.Raise(err, 300005)
	}

	err = tx.Commit()
	if err != nil {
		return userId, terrors.Raise(err, 300004)
	}

	return userId, err
}
