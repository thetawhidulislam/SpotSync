package middlewares

import (
	"net/http"
	"spotsync/internal/auth"
	"spotsync/internal/httpresponse"
	"strings"

	"github.com/labstack/echo/v5"
)

func AuthMiddleware(jwtService auth.JWTService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, httpresponse.Error{
					Code:    http.StatusUnauthorized,
					Message: "Unauthorized",
					Details: "missing authorization header",
				})
			}

			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				return c.JSON(http.StatusUnauthorized, httpresponse.Error{
					Code:    http.StatusUnauthorized,
					Message: "Unauthorized",
					Details: "invalid authorization header format",
				})
			}

			claims, err := jwtService.ValidateToken(parts[1])
			if err != nil {
				return c.JSON(http.StatusUnauthorized, httpresponse.Error{
					Code:    http.StatusUnauthorized,
					Message: "Unauthorized",
					Details: "invalid or expired token",
				})
			}

			// reject refresh tokens from being used as access tokens
			if claims.TokenType != auth.TokenTypeAccess {
				return c.JSON(http.StatusUnauthorized, httpresponse.Error{
					Code:    http.StatusUnauthorized,
					Message: "Unauthorized",
					Details: "invalid token type",
				})
			}

			c.Set("user_id", claims.UserID)
			c.Set("user_email", claims.Email)
			c.Set("role", claims.Role)
			c.Set("user_name", claims.Name)

			return next(c)
		}
	}
}

func AdminMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {

			role, ok := c.Get("role").(string)
			if !ok || role != "admin" {
				return c.JSON(http.StatusForbidden, httpresponse.Error{
					Code:    http.StatusForbidden,
					Message: "Forbidden",
					Details: "admin access required",
				})
			}

			return next(c)
		}
	}
}
