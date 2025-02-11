package auth

import (
	"github.com/charitan-go/auth-server/internal/auth/handler"
	"github.com/charitan-go/auth-server/internal/auth/repository"
	"github.com/charitan-go/auth-server/internal/auth/service"
	"go.uber.org/fx"
)

var AuthModule = fx.Module("auth",
	fx.Provide(
		handler.NewAuthHandler,
		service.NewAuthService,
		service.NewPasswordService,
		service.NewJwtService,
		repository.NewAuthRepository,
		// profile.NewProfileProtoClient,
	),
)
