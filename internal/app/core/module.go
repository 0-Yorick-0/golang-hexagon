package core

import (
	"golang-hexagon/internal/app/core/ports"
	"golang-hexagon/internal/app/core/services"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		fx.Annotate(services.NewUrlService, fx.As(new(ports.UrlService))),
		fx.Annotate(services.NewUrlScraper, fx.As(new(ports.ScraperService))),
	),
)
