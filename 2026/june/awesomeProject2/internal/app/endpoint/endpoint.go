package endpoint

import (
	"awesomeProject2/internal/app/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Endpoint struct {
	service *service.Service
}

func NewEndpoint() *Endpoint {
	return &Endpoint{
		service: service.NewService(),
	}
}
func (e *Endpoint) Status(ctx echo.Context) error {
	value := e.service.DaysLeft()
	return ctx.String(http.StatusOK, strconv.FormatInt(value, 10))
}
