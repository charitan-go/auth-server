package service

import (
	"fmt"
	"log"
	"net/http"

	"github.com/charitan-go/auth-server/external/profile"
	"github.com/charitan-go/auth-server/internal/auth/dto"
	"github.com/charitan-go/auth-server/internal/auth/model"
	"github.com/charitan-go/auth-server/internal/auth/repository"
	"github.com/charitan-go/auth-server/pkg/proto"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	LoginUser(req *dto.LoginUserRequestDto) (*dto.LoginUserResponseDto, *dto.ErrorResponseDto)
	RegisterDonor(req *dto.RegisterDonorRequestDto) (*dto.RegisterResponseDto, *dto.ErrorResponseDto)
}

type authServiceImpl struct {
	r                 repository.AuthRepository
	profileGrpcClient profile.ProfileGrpcClient
}

func NewAuthService(r repository.AuthRepository, profileProtoClient profile.ProfileGrpcClient) AuthService {
	return &authServiceImpl{r: r, profileGrpcClient: profileProtoClient}
}

func (svc *authServiceImpl) RegisterDonor(req *dto.RegisterDonorRequestDto) (*dto.RegisterResponseDto, *dto.ErrorResponseDto) {
	// Check does email existed
	existedEmailDonor, _ := svc.r.FindOneByEmail(req.Email)

	if existedEmailDonor != nil {
		return nil, &dto.ErrorResponseDto{Message: "Email already existed", StatusCode: http.StatusBadRequest}
	}

	createDonorProfileRequestDto := &proto.CreateDonorProfileRequestDto{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Address:   req.Address,
	}
	// createDonorProfileResponseDto, err := protoclient.ProfileClient.CreateDonorProfile(*protoclient.ProfileCtx, createDonorProfile)
	createDonorProfileResponseDto, err := svc.profileGrpcClient.CreateDonorProfile(createDonorProfileRequestDto)
	if err != nil {
		errorMessage := fmt.Sprintf("Cannot send to profile-server: %v\n", err)
		log.Fatalln(errorMessage)
		return nil, &dto.ErrorResponseDto{Message: errorMessage, StatusCode: http.StatusInternalServerError}
	}

	// Parse profileId
	profileReadableIdStr := createDonorProfileResponseDto.GetProfileReadableId()
	profileReadableId, err := uuid.Parse(profileReadableIdStr)
	if err != nil {
		errorMessage := fmt.Sprintf("Cannot parse profileReadableId: %v", err)
		log.Fatalln(errorMessage)
		return nil, &dto.ErrorResponseDto{Message: errorMessage, StatusCode: http.StatusInternalServerError}
	}

	// Hash password
	hashedPasswordByte, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, &dto.ErrorResponseDto{Message: "Error in hashedPassword", StatusCode: http.StatusInternalServerError}
	}
	hashedPassword := string(hashedPasswordByte)
	_ = hashedPassword

	// Save to repo
	authModel := model.NewAuth(
		req,
		hashedPassword,
		dto.RoleDonor,
		profileReadableId)
	_, err = svc.r.Save(authModel)
	if err != nil {
		return nil, &dto.ErrorResponseDto{Message: "Failed to save to database", StatusCode: http.StatusInternalServerError}
	}

	// Return response
	return &dto.RegisterResponseDto{Message: "Register successfully"}, nil
}

// LoginUser implements AuthService.
func (svc *authServiceImpl) LoginUser(req *dto.LoginUserRequestDto) (*dto.LoginUserResponseDto, *dto.ErrorResponseDto) {
	// TODO: Implements
	return &dto.LoginUserResponseDto{Token: "312321312312"}, nil
}
