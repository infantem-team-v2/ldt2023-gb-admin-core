package http

import (
	"gb-auth-gate/internal/auth/model"
	mdwModel "gb-auth-gate/internal/pkg/middleware/model"
	"gb-auth-gate/pkg/terrors"
	"gb-auth-gate/pkg/thttp/server"
	"github.com/gofiber/fiber/v2"
	"strings"
	"time"
)

// SignatureMiddleware Validates request by HMAC512
func (mdw *MiddlewareManager) SignatureMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		headers := c.GetReqHeaders()
		body := c.Body()
		ok, err := mdw.AuthUC.ValidateService(&model.AuthHeadersLogic{
			Signature: headers[mdwModel.SignatureKey],
			PublicKey: headers[mdwModel.PublicKey],
			Body:      body,
		})
		if !ok {
			if err == nil {
				return terrors.Raise(nil, 300001)
			}
			return err
		}
		return c.Next()
	}
}

// JWTMiddleware Validates request by access token
func (mdw *MiddlewareManager) JWTMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		path := c.Path()
		if (strings.Contains(path, "/auth/") &&
			!strings.Contains(path, "sign/out") &&
			!strings.Contains(path, "check")) ||
			strings.Contains(path, "/base") ||
			strings.Contains(path, "/docs") ||
			strings.Contains(path, "/ui/") {
			return c.Next()
		}
		var err error
		accessToken := c.Cookies(mdwModel.AccessKey)
		if accessToken == "" {
			refreshToken := c.Cookies(mdwModel.RefreshKey)
			if refreshToken == "" {
				return terrors.Raise(nil, 100004)
			}
			claims, err := server.ParseJwtToken(refreshToken, mdw.Config.HttpConfig.JWTSalt)
			if err != nil {
				return err
			}
			if claims == nil {
				return terrors.Raise(nil, 100011)
			}
			userId := int64(claims["userId"].(float64))

			accessToken, err = mdw.AuthUC.GenerateAccessToken(refreshToken, &model.CreateAuthTokensLogic{
				UserId: userId,
			})
			if err != nil {
				return err
			}

			ok, err := mdw.AuthUC.ValidateUser(userId)
			if err != nil {
				return err
			}
			if !ok {
				return terrors.Raise(nil, 100011)
			}

			c.Cookie(&fiber.Cookie{
				Name:     mdwModel.AccessKey,
				Value:    accessToken,
				Domain:   mdw.Config.BaseConfig.Service.URL,
				MaxAge:   int((time.Duration(mdw.Config.HttpConfig.AccessExpireTime) * time.Second).Seconds()),
				Secure:   true,
				Path:     "/",
				HTTPOnly: true,
				SameSite: "None",
			})
			c.Locals(mdwModel.UserIdLocals, userId)
			return c.Next()
		}
		claims, err := server.ParseJwtToken(accessToken, mdw.Config.HttpConfig.JWTSalt)
		if err != nil {
			return err
		}
		userId := int64(claims["userId"].(float64))
		ok, err := mdw.AuthUC.ValidateUser(userId)
		if err != nil {
			return err
		}
		if !ok {
			return terrors.Raise(nil, 100011)
		}

		c.Locals(mdwModel.UserIdLocals, userId)
		return c.Next()
	}
}

// FiberSessionMiddleware Validates request by fiber's session are saved in redis' storage
func (mdw *MiddlewareManager) FiberSessionMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {

		return c.Next()
	}
}
