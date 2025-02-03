package service

import (
	"net/http"

	"github.com/charitan-go/auth-server/domain/auth/dto"
	"github.com/charitan-go/auth-server/domain/auth/model"
	"github.com/charitan-go/auth-server/domain/auth/repository"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	RegisterDonor(req dto.RegisterDonorRequestDto) (*dto.RegisterResponseDto, *dto.ErrorResponseDto)
}

type authServiceImpl struct {
	r repository.AuthRepository
}

func NewAuthService(r repository.AuthRepository) AuthService {
	return &authServiceImpl{r: r}
}

func (svc *authServiceImpl) RegisterDonor(req dto.RegisterDonorRequestDto) (*dto.RegisterResponseDto, *dto.ErrorResponseDto) {
	// Check does email existed
	existedEmailDonor, _ := svc.r.FindOneByEmail(req.Email)

	if existedEmailDonor != nil {
		return nil, &dto.ErrorResponseDto{Message: "Email already existed", StatusCode: http.StatusBadRequest}
	}

	// TODO Send kafka topic
	profileReadableId := uuid.New()

	// Hash password
	hashedPasswordByte, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, &dto.ErrorResponseDto{Message: "Error in hashedPassword", StatusCode: http.StatusInternalServerError}
	}
	hashedPassword := string(hashedPasswordByte)

	authModel := model.NewAuth(req, hashedPassword, dto.RoleDonor, profileReadableId)

	// Save to repo
	_, err = svc.r.Save(authModel)
	if err != nil {
		return nil, &dto.ErrorResponseDto{Message: "Failed to save to database", StatusCode: http.StatusInternalServerError}
	}

	return &dto.RegisterResponseDto{Message: "Register successfully"}, nil
}
