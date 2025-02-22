package rabbitmq

import (
	"github.com/charitan-go/auth-server/rabbitmq/service"
	"go.uber.org/fx"
)

var RabbitmqModule = fx.Module("rabbitmq",
	fx.Provide(
		NewRabbitmqServer,
		service.NewRabbitmqService,
	),
)
