package middleware

import (
	"app/internal/auth"
	"app/internal/repository"

	"github.com/labstack/echo/v4"
)

func AttachUser(users *repository.UserRepository, sessions *auth.SessionStore) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userID, ok := sessions.UserID(c)
			if !ok {
				return next(c)
			}
			user, err := users.GetByID(c.Request().Context(), userID)
			if err != nil {
				_ = sessions.Clear(c)
				return next(c)
			}
			c.Set("userID", user.ID)
			c.Set("username", user.Username)
			return next(c)
		}
	}
}
