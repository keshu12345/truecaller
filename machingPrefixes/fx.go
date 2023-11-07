package machingPrefixes

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewGetMatchingPrefixesService),
	fx.Provide(NewHealthCheckService),
	fx.Invoke(RegisterMatchingPrefixesformEndPoint),
)
