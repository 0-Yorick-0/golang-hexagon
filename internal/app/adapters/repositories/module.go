package repositories

import "go.uber.org/fx"

var Module = fx.Options(
	mongodb.Module,
)
