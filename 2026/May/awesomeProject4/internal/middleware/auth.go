package middleware

import (
	"net/http"

	"app/internal/auth"

	"github.com/labstack/echo/v4"
)

func RequireAuth(sessions *auth.SessionStore) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userID, ok := sessions.UserID(c)
			if !ok {
				return c.Redirect(http.StatusSeeOther, "/login")
			}
			c.Set("userID", userID)
			return next(c)
		}
	}
}

func RedirectIfAuth(sessions *auth.SessionStore) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if _, ok := sessions.UserID(c); ok {
				return c.Redirect(http.StatusSeeOther, "/")
			}
			return next(c)
		}
	}
}
