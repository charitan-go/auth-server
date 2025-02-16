package service

import (
	"fmt"
	"log"
	"net/http"

	"github.com/charitan-go/auth-server/external/key"
	"github.com/charitan-go/auth-server/external/profile"
	"github.com/charitan-go/auth-server/internal/auth/dto"
	"github.com/charitan-go/auth-server/internal/auth/model"
	"github.com/charitan-go/auth-server/internal/auth/repository"
	"github.com/charitan-go/auth-server/pkg/proto"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(req *dto.LoginUserRequestDto) (*dto.LoginUserResponseDto, *dto.ErrorResponseDto)
	RegisterDonor(req *dto.RegisterDonorRequestDto) (*dto.RegisterResponseDto, *dto.ErrorResponseDto)

	HandleGetPrivateKeyRabbitmq() error
}

type authServiceImpl struct {
	passwordService   PasswordService
	jwtService        JwtService
	r                 repository.AuthRepository
	profileGrpcClient profile.ProfileGrpcClient
	keyGrpcClient     key.KeyGrpcClient
}

func verifyPassword(hashPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err == nil
}

func NewAuthService(passwordService PasswordService, jwtService JwtService, r repository.AuthRepository, profileGrpcClient profile.ProfileGrpcClient, keyGrpcClient key.KeyGrpcClient) AuthService {
	return &authServiceImpl{passwordService, jwtService, r, profileGrpcClient, keyGrpcClient}
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
	hashedPassword, err := svc.passwordService.HashPassword(req.Password)
	if err != nil {
		return nil, &dto.ErrorResponseDto{Message: "Error in hashedPassword", StatusCode: http.StatusInternalServerError}
	}

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
func (svc *authServiceImpl) Login(req *dto.LoginUserRequestDto) (*dto.LoginUserResponseDto, *dto.ErrorResponseDto) {
	// Check user existed or not
	existedUser, err := svc.r.FindOneByEmail(req.Email)
	if err != nil {
		return nil, &dto.ErrorResponseDto{Message: "Invalid credentials", StatusCode: http.StatusBadRequest}
	}

	// Verify password
	if !svc.passwordService.VerifyPassword(existedUser.HashedPassword, req.Password) {
		return nil, &dto.ErrorResponseDto{Message: "Invalid credentials", StatusCode: http.StatusBadRequest}
	}

	// Sign JWT
	token, err := svc.jwtService.SignToken(existedUser)
	if err != nil {
		return nil, &dto.ErrorResponseDto{Message: "Error happen in sign token", StatusCode: http.StatusInternalServerError}
	}

	return &dto.LoginUserResponseDto{Token: token}, nil
}

func (svc *authServiceImpl) HandleGetPrivateKeyRabbitmq() error {
	getPrivateKeyRequestDto := &proto.GetPrivateKeyRequestDto{}
	// createDonorProfileRequestDto := &proto.CreateDonorProfileRequestDto{
	// 	FirstName: req.FirstName,
	// 	LastName:  req.LastName,
	// 	Address:   req.Address,
	// }

	getPrivateKeyResponseDto, err := svc.keyGrpcClient.GetPrivateKey(getPrivateKeyRequestDto)
	if err != nil {
		log.Fatalf("Cannot get private key from key-server: %v\n", err)
		return err
	}

	err = svc.jwtService.UpdatePrivateKey(getPrivateKeyResponseDto.PrivateKey)
	if err != nil {
		log.Fatalf("Cannot update private key")
		return err
	}

	return nil
}
