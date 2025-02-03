package profile

import (
	context "context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	Client ProfileGrpcServiceClient
	Ctx    context.Context
)

func SetupProfileGrpcServiceClient() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("connection failed: %v", err)
	}
	defer conn.Close()

	Client = NewProfileGrpcServiceClient(conn)

	Ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	_ = Ctx
	defer cancel()
}
