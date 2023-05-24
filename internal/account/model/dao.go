package model

type UserDAO struct {
	FullName    string `db:"full_name"`
	Email       string `db:"email"`
	JobPosition string `db:"job_position"`
	Geography   string `db:"geo"`
}

type BusinessDAO struct {
	Inn              string `db:"inn"`
	Name             string `db:"name"`
	EconomicActivity string `db:"economic_activity"`
	Website          string `db:"website"`
}

type UpdateUserDataDAO struct {
	UserId int64 `db:"user_id"`
	GeoId  int64 `db:"geo_id"`

	FullName    string `db:"full_name"`
	Email       string `db:"email"`
	JobPosition string `db:"job_position"`

	City    string `db:"city"`
	Country string `db:"country"`

	Inn              string `db:"inn"`
	Name             string `db:"name"`
	EconomicActivity string `db:"economic_activity"`
	Website          string `db:"website"`
}
