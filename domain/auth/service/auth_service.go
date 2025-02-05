package service

import (
	"fmt"
	"net/http"

	"github.com/charitan-go/auth-server/domain/auth/dto"
	"github.com/charitan-go/auth-server/domain/auth/repository"
	"github.com/charitan-go/auth-server/domain/profile"
	"github.com/charitan-go/auth-server/pkg/proto"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	RegisterDonor(req dto.RegisterDonorRequestDto) (*dto.RegisterResponseDto, *dto.ErrorResponseDto)
}

type authServiceImpl struct {
	r                  repository.AuthRepository
	profileProtoClient profile.ProfileProtoClient
}

func NewAuthService(r repository.AuthRepository, profileProtoClient profile.ProfileProtoClient) AuthService {
	return &authServiceImpl{r: r, profileProtoClient: profileProtoClient}
}

func (svc *authServiceImpl) RegisterDonor(req dto.RegisterDonorRequestDto) (*dto.RegisterResponseDto, *dto.ErrorResponseDto) {
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
	createDonorProfileResponseDto, err := svc.profileProtoClient.CreateDonorProfile(createDonorProfileRequestDto)
	if err != nil {
		fmt.Printf("Cannot send to profile-server: %v\n", err)
	}
	profileReadableId := createDonorProfileResponseDto.GetProfileReadableId()
	fmt.Println("ProfileReabableid = ", profileReadableId)

	// Hash password
	hashedPasswordByte, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, &dto.ErrorResponseDto{Message: "Error in hashedPassword", StatusCode: http.StatusInternalServerError}
	}
	hashedPassword := string(hashedPasswordByte)
	_ = hashedPassword

	// authModel := model.NewAuth(req, hashedPassword, dto.RoleDonor, profileReadableId)
	//
	// // Save to repo
	// _, err = svc.r.Save(authModel)
	// if err != nil {
	// 	return nil, &dto.ErrorResponseDto{Message: "Failed to save to database", StatusCode: http.StatusInternalServerError}
	// }
	//

	return &dto.RegisterResponseDto{Message: "Register successfully"}, nil
}
