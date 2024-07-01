package app

import (
	"golang-hexagon/internal/app/infrastructure/config"

	"go.uber.org/fx"
)

func newApp(cfg *config.AppConfig) *fx.App {
	return fx.New(
		fx.Provide(func() *config.AppConfig {
			return cfg
		}),
		adapters.Module,
		core.Module,
		server.Module,
		fx.Invoke(server.StartServer),
	)
}
