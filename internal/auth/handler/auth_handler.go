package handler

import (
	"log"
	"net/http"

	"github.com/charitan-go/auth-server/internal/auth/dto"
	"github.com/charitan-go/auth-server/internal/auth/service"
	restpkg "github.com/charitan-go/auth-server/pkg/rest"
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

	req := new(dto.RegisterDonorRequestDto)
	if err := c.Bind(req); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, dto.ErrorResponseDto{Message: "Invalid request bodyy", StatusCode: http.StatusBadRequest})
	}

	res, errRes := h.svc.HandleRegisterDonorRest(req)
	if errRes != nil {
		return c.JSON(int(errRes.StatusCode), *errRes)
	}

	return c.JSON(http.StatusCreated, *res)
}

func (h *AuthHandler) RegisterCharity(c echo.Context) error {

	req := new(dto.RegisterCharityRequestDto)
	if err := c.Bind(req); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, dto.ErrorResponseDto{Message: "Invalid request bodyy", StatusCode: http.StatusBadRequest})
	}

	res, errRes := h.svc.HandleRegisterCharityRest(req)
	if errRes != nil {
		return c.JSON(int(errRes.StatusCode), *errRes)
	}

	return c.JSON(http.StatusCreated, *res)
}

func (h *AuthHandler) Login(c echo.Context) error {
	var req dto.LoginUserRequestDto

	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponseDto{Message: "Invalid request body"})
	}

	res, errRes := h.svc.HandleLoginRest(&req)
	if errRes != nil {
		return c.JSON(int(errRes.StatusCode), *errRes)
	}

	return c.JSON(http.StatusOK, *res)
}

func (h *AuthHandler) GetMe(c echo.Context) error {

	jwtPayload, err := restpkg.GetJwtPayload(c)
	if err != nil {
		log.Fatalln("Not found header payload")
		return c.JSON(http.StatusNonAuthoritativeInfo, dto.ErrorResponseDto{Message: "Not authorized"})
	}

	log.Println("User id is " + jwtPayload.ReadableId)

	res, errRes := h.svc.HandleGetMeRest(jwtPayload)
	if errRes != nil {
		return c.JSON(int(errRes.StatusCode), *errRes)
	}

	// TODO: Implement
	return c.JSON(http.StatusOK, res)
}
