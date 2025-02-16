package profile

import (
	"context"
	"fmt"
	"time"

	"github.com/charitan-go/auth-server/pkg/discovery"
	"github.com/charitan-go/auth-server/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const PROFILE_GRPC_SERVICE_NAME = "profile-server-grpc"

type ProfileGrpcClient interface {
	CreateDonorProfile(reqDto *proto.CreateDonorProfileRequestDto) (*proto.CreateDonorProfileResponseDto, error)

	GetDonorProfile(reqDto *proto.GetDonorProfileRequestDto) (*proto.GetDonorProfileResponseDto, error)
}

type profileProtoClientImpl struct{}

func NewProfileGrpcClient() ProfileGrpcClient {
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
	profileServerAddress := discovery.DiscoverService(PROFILE_GRPC_SERVICE_NAME)

	// Connect to the gRPC server
	conn, err := grpc.NewClient(profileServerAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("connection failed: %v", err)
	}
	defer conn.Close()

	// Create a client
	client := proto.NewProfileGrpcServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	responseDto, err := client.CreateDonorProfile(ctx, reqDto)
	if err != nil {
		return nil, fmt.Errorf("CreateDonorProfile failed: %v", err)
	}

	return responseDto, nil
}

// GetDonorProfile implements ProfileGrpcClient.
func (c *profileProtoClientImpl) GetDonorProfile(reqDto *proto.GetDonorProfileRequestDto) (*proto.GetDonorProfileResponseDto, error) {
	profileServerAddress := discovery.DiscoverService(PROFILE_GRPC_SERVICE_NAME)

	// Connect to the gRPC server
	conn, err := grpc.NewClient(profileServerAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("connection failed: %v", err)
	}
	defer conn.Close()

	// Create a client
	client := proto.NewProfileGrpcServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	responseDto, err := client.GetDonorProfile(ctx, reqDto)
	if err != nil {
		return nil, fmt.Errorf("GetDonorProfile failed: %v", err)
	}

	return responseDto, nil
}
