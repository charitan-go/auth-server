package api

import (
	"net/http"

	authhandler "github.com/charitan-go/auth-server/domain/auth/handler"

	"github.com/labstack/echo/v4"
)

type Api struct {
	AuthHandler *authhandler.AuthHandler
}

func NewApi(authHandler *authhandler.AuthHandler) *Api {
	return &Api{
		AuthHandler: authHandler,
	}
}

func (*Api) HealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
