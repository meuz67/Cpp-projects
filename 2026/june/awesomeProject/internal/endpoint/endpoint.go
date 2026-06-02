package endpoint

import (
	"fmt"
	"june/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Endpoint struct {
	s *service.Service
}

func New(svc *service.Service) *Endpoint {
	return &Endpoint{s: svc}
}
func (e *Endpoint) Status(ctx echo.Context) error {
	days := e.s.Daysuntil()
	value := fmt.Sprintf("Days until 2027: %d", days)
	err := ctx.String(http.StatusOK, value)
	if err != nil {
		return err
	}
	return nil
}
