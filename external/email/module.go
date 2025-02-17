package email

import (
	"go.uber.org/fx"
)

var EmailModule = fx.Module("email",
	fx.Provide(
		NewEmailRabbitmqProducer,
	),
)
