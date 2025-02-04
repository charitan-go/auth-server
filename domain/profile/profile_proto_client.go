package profile

import (
	"context"
	"fmt"
	"time"

	"github.com/charitan-go/auth-server/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var PROFILE_GRPC_SERVER_ADDRESS = "profile-server:50051"

type ProfileProtoClient interface {
	CreateDonorProfile(reqDto *proto.CreateDonorProfileRequestDto) (*proto.CreateDonorProfileResponseDto, error)
}

type profileProtoClientImpl struct{}

func NewProfileProtoClient() ProfileProtoClient {
	return &profileProtoClientImpl{}
}

// type AuthService interface {
// 	RegisterDonor(req dto.RegisterDonorRequestDto) (*dto.RegisterResponseDto, *dto.ErrorResponseDto)
// }
//
// type authServiceImpl struct {
// 	r                  repository.AuthRepository
// 	profileProtoClient profile.ProfileProtoClient
// }
//
// func NewAuthService(r repository.AuthRepository, profileProtoClient profile.ProfileProtoClient) AuthService {
// 	return &authServiceImpl{r: r, profileProtoClient: profileProtoClient}
// }

func (c *profileProtoClientImpl) CreateDonorProfile(reqDto *proto.CreateDonorProfileRequestDto) (*proto.CreateDonorProfileResponseDto, error) {
	// Connect to the gRPC server
	conn, err := grpc.NewClient(PROFILE_GRPC_SERVER_ADDRESS, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("connection failed: %v", err)
	}
	defer conn.Close()

	// Create a client
	client := proto.NewProfileServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	responseDto, err := client.CreateDonorProfile(ctx, reqDto)
	if err != nil {
		return nil, fmt.Errorf("CreateDonorProfile failed: %v", err)
	}

	return responseDto, nil
}
