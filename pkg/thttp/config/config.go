package config

type ThttpConfig struct {
	TimeOut           uint32 `json:"timeOut"`           // request timeout in seconds
	Accept            string `json:"accept"`            // what types of content accepts
	DoLogRequests     uint8  `json:"doLogRequests"`     // do logging of requests and responses
	JWTSalt           string `json:"jwtSalt"`           // string for signature salt
	AccessExpireTime  uint16 `json:"accessExpireTime"`  // access token expiration time in seconds
	RefreshExpireTime uint16 `json:"refreshExpireTime"` // refresh token expiration time in minutes
}

func (thc *ThttpConfig) ParseConfig() error {
	return nil
}
