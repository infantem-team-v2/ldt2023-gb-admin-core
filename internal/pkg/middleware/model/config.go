package model

type MiddlewareConfig struct {
	Auth struct {
		ExpirationTime int `json:"expirationTime"`
	} `json:"auth"`
}
