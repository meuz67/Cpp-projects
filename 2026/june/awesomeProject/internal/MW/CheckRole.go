package MW

import (
	"log"
	"strings"

	"github.com/labstack/echo/v4"
)

func CheckRole(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		value := c.Request().Header.Get("user-role")
		if strings.ToLower(value) == "admin" {
			log.Println("Red button user detected")
		}
		return next(c)
	}
}
