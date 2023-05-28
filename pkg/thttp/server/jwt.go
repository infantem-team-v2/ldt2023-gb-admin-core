package server

import (
	"gb-admin-core/pkg/terrors"
	"github.com/golang-jwt/jwt/v5"
	"strings"
	"time"
)

type JwtParams struct {
	Salt *string `json:"-"`

	UserId    *int64                 `json:"userId"`
	Duration  *time.Time             `json:"exp"`
	Type      *string                `json:"type"`
	OtherData map[string]interface{} `json:"other_data"`
}

func ParseJwtToken(token, salt string) (jwt.MapClaims, error) {
	tokenParsed, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		claims := token.Claims.(jwt.MapClaims)
		if claims["userId"] == nil || claims["userId"] == 0 {
			return nil, terrors.Raise(nil, 100006)
		}
		if token.Method != jwt.SigningMethodHS512 {
			return nil, terrors.Raise(nil, 100008)
		}

		return []byte(salt), nil
	})
	if err != nil {
		if strings.Contains(err.Error(), "token is expired") {
			return nil, nil
		}
		return nil, terrors.Raise(err, 100009)
	}

	return tokenParsed.Claims.(jwt.MapClaims), nil
}

func CreateJwtToken(params *JwtParams) (string, error) {
	token := jwt.New(jwt.SigningMethodHS512)
	claims := token.Claims.(jwt.MapClaims)

	claims["userId"] = *params.UserId
	claims["exp"] = params.Duration.Unix()
	claims["type"] = *params.Type
	claims["iat"] = time.Now().Unix()
	claims["nbf"] = time.Now().Unix()

	if params.OtherData != nil {
		for k, v := range params.OtherData {
			claims[k] = v
		}
	}

	return token.SignedString([]byte(*params.Salt))
}
