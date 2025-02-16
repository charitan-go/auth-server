package restpkg

import (
	"errors"

	"github.com/labstack/echo/v4"
)

type HeaderPayload struct {
	ReadableId string
	Role       string
}

func GetHeaderPayload(c echo.Context) (*HeaderPayload, error) {
	readableId := c.Request().Header.Get("X-User-Id")
	role := c.Request().Header.Get("X-User-Role")

	if readableId == "" || role == "" {
		return nil, errors.New("Miss info of header payload")
	}

	return &HeaderPayload{ReadableId: readableId, Role: role}, nil
}

