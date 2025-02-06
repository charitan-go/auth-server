package profile

import "go.uber.org/fx"

var ProfileModule = fx.Module("profile",
	fx.Provide(
		NewProfileGrpcClient,
	),
)
