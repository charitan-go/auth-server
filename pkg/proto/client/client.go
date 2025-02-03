package proto

import "github.com/charitan-go/auth-server/pkg/proto/client/profile"

func SetupGrpcServiceClient() {
	profile.SetupProfileGrpcServiceClient()
}
