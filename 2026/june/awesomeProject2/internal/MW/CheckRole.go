package MW

import (
	"log"
	"strings"

	"github.com/labstack/echo/v4"
)

func CheckRole(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		value := ctx.Request().Header.Get("user-role")
		if strings.ToLower(value) == "admin" {
			log.Println("Admin role")
		}
		return next(ctx)
	}
}
