package adapters

import "go.uber.org/fx"

var Module = fx.Options(
	handlers.Module,
	repositories.Module,
)
