package repositories

import (
	"golang-hexagon/internal/app/adapters/repositories/mongodb"

	"go.uber.org/fx"
)

var Module = fx.Options(
	mongodb.Module,
)
