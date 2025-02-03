package handler

import (
	"net/http"

	"github.com/charitan-go/auth-server/domain/auth/dto"
	"github.com/charitan-go/auth-server/domain/auth/service"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	svc service.AuthService
}

func (h *AuthHandler) CheckHealth() string {
	return "OK"
}

func NewAuthHandler(svc service.AuthService) *AuthHandler {
	return &AuthHandler{svc: svc}
}

func (h *AuthHandler) RegisterDonor(c echo.Context) error {
	var req dto.RegisterDonorRequestDto

	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponseDto{Message: "Invalid request body"})
	}

	res, errRes := h.svc.RegisterDonor(req)
	if errRes != nil {
		return c.JSON(int(errRes.StatusCode), *errRes)
	}

	return c.JSON(http.StatusCreated, *res)
}
