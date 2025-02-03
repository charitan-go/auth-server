package service

import (
	"fmt"
	"net/http"

	"github.com/charitan-go/auth-server/domain/auth/dto"
	"github.com/charitan-go/auth-server/domain/auth/repository"
	"github.com/charitan-go/auth-server/pkg/proto"
	protoclient "github.com/charitan-go/auth-server/pkg/proto/client"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	RegisterDonor(req dto.RegisterDonorRequestDto) (*dto.RegisterResponseDto, *dto.ErrorResponseDto)
}

type authServiceImpl struct {
	r repository.AuthRepository
	// profileProtoClient proto.ProfileServiceClient
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

	// TODO Send GRPC to profile to have profileReadableId
	createDonorProfile := &proto.CreateDonorProfileRequestDto{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Address:   req.Address,
	}
	createDonorProfileResponseDto, err := protoclient.ProfileClient.CreateDonorProfile(protoclient.ProfileCtx, createDonorProfile)
	if err != nil {
		fmt.Println("Cannot send to profile-server")
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
