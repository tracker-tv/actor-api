package testutils

import (
	"github.com/caarlos0/env/v11"
	"github.com/tracker-tv/actor-api/internal/config"
)

func SetupConfig() (*config.Config, error) {
	var cfg config.Config
	err := env.Parse(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
