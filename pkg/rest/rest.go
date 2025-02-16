package restpkg

import (
	"errors"

	"github.com/labstack/echo/v4"
)

type JwtPayload struct {
	ReadableId string
	Role       string
}

func GetJwtPayload(c echo.Context) (*JwtPayload, error) {
	readableId := c.Request().Header.Get("X-User-Id")
	role := c.Request().Header.Get("X-User-Role")

	if readableId == "" || role == "" {
		return nil, errors.New("Miss info of header payload")
	}

	return &JwtPayload{ReadableId: readableId, Role: role}, nil
}
