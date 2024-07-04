package mongodb

import (
	"golang-hexagon/internal/app/core/ports"
	"golang-hexagon/internal/app/infrastructure/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		func(config *config.AppConfig) MongoDBConfig {
			conn := config.Mongo.ConnectionString
			if conn == "" && config.Environment == "Development" { // for local development
				conn = "mongodb://localhost:27017"
			}

			return MongoDBConfig{URI: conn}
		},
		func(cfg MongoDBConfig) (*mongo.Client, error) {
			return ConntectToMongo(cfg)
		},
	),
	fx.Provide(
		fx.Annotate(NewUrlRepository, fx.As(new(ports.UrlRepository))),
	),
)
