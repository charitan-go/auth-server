package auth

import (
	"github.com/charitan-go/auth-server/domain/auth/handler"
	"github.com/charitan-go/auth-server/domain/auth/repository"
	"github.com/charitan-go/auth-server/domain/auth/service"
	"go.uber.org/fx"
)

var AuthModule = fx.Module("auth",
	fx.Provide(
		handler.NewAuthHandler,
		service.NewAuthService,
		repository.NewAuthRepository,
	),
)
