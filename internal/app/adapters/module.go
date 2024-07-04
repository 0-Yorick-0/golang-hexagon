package adapters

import (
	"golang-hexagon/internal/app/adapters/repositories"

	"go.uber.org/fx"
)

var Module = fx.Options(
	handlers.Module,
	repositories.Module,
)
