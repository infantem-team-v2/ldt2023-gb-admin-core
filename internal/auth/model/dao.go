package model

type AuthServiceDAO struct {
	Id         uint64 `db:"id"`
	Name       string `db:"name"`
	PublicKey  string `db:"public_key"`
	PrivateKey string `db:"private_key"`
	URL        string `db:"url"`
}

type UserShortDAO struct {
	Id       int64  `db:"id"`
	FullName string `db:"full_name"`
	Email    string `db:"email"`
}

type AuthUserDAO struct {
	UserId   int64  `db:"id"`
	Email    string `db:"email"`
	Password string `db:"password"`
}

type CreateUserDAO struct {
	UserId *int64 `db:"user_id"`

	FullName    *string `db:"full_name"`
	Email       *string `db:"email"`
	Password    *string `db:"password"`
	JobPosition *string `db:"job_position"`

	City    *string `db:"city"`
	Country *string `db:"country"`

	Inn              *string `db:"inn"`
	BusinessName     *string `db:"business_name"`
	EconomicActivity *string `db:"economic_activity"`
	Website          *string `db:"website"`
}
