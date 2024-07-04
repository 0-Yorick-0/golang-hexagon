package config

import "github.com/jinzhu/configor"

type Kafka struct {
	Broker  string
	Topic   string
	GroupId string
}

type Mongo struct {
	ConnectionString string
	Collection       string
}

type AppConfig struct {
	Environment string
	Port        string
	Kafka       Kafka
	Mongo       Mongo
}

func NewAppConfig(path string) (*AppConfig, error) {
	config := new(AppConfig)
	// configor is a package that allow to automatically inject config
	// into a struct with a path
	err := configor.
		New(&configor.Config{ErrorOnUnmatchedKeys: true}).
		Load(config, path)
	if err != nil {
		return nil, err
	}
	return config, nil
}
