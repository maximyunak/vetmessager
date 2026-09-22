package auth

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Secret    []byte        `envconfig:"SECRET" required:"true"`
	AccessTTL time.Duration `envconfig:"ACCESS_TTL" default:"168h"`
}

func NewConfig() (Config, error) {
	var config Config
	if err := envconfig.Process("JWT", &config); err != nil {
		return config, err
	}
	return config, nil
}

func NewConfigMust() Config {
	config, err := NewConfig()
	if err != nil {
		err = fmt.Errorf("get jwt config: %w", err)
		panic(err)
	}

	return config
}
