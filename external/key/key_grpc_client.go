package key

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/charitan-go/auth-server/pkg/discovery"
	"github.com/charitan-go/auth-server/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const KEY_GRPC_SERVICE_NAME = "key-server-grpc"

type KeyGrpcClient interface {
	GetPrivateKey(reqDto *proto.GetPrivateKeyRequestDto) (*proto.GetPrivateKeyResponseDto, error)
}

type keyGrpcClientImpl struct{}

func NewKeyGrpcClient() KeyGrpcClient {
	return &keyGrpcClientImpl{}
}

func (*keyGrpcClientImpl) GetPrivateKey(reqDto *proto.GetPrivateKeyRequestDto) (*proto.GetPrivateKeyResponseDto, error) {
	keyServerAddress := discovery.DiscoverService(KEY_GRPC_SERVICE_NAME)

	// Connect to the gRPC server
	conn, err := grpc.NewClient(keyServerAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("connection failed: %v", err)
	}
	defer conn.Close()

	// Create a client
	client := proto.NewKeyGrpcServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	responseDto, err := client.GetPrivateKey(ctx, reqDto)
	if err != nil {
		msg := fmt.Sprintf("Failed to get private key: %v", err)
		log.Println(msg)
		return nil, errors.New(msg)
	}

	return responseDto, nil
}
