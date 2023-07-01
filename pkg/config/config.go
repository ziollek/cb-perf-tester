package config

import (
	"github.com/spf13/viper"
)

type Couchbase struct {
	URI      string `yaml:"uri"`
	Bucket   string `yaml:"bucket"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type Config struct {
	Couchbase *Couchbase `yaml:"couchbase"`
}

func GetConfig() (*Config, error) {
	var config Config
	err := viper.Unmarshal(&config)
	return &config, err
}
