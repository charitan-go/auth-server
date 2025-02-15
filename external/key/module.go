package key

import "go.uber.org/fx"

var KeyModule = fx.Module("key",
	fx.Provide(
		NewKeyGrpcClient,
	),
)
