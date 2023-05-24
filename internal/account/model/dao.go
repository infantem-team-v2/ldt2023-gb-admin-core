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
}
