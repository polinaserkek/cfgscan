package parser

import (
	"cfgscan/internal/config"
	"encoding/json"
)

type JSONParser struct {
}

func (jp JSONParser) Parse(data []byte) (config.Config, error) {
	cfg := config.Config{}
	err := json.Unmarshal(data, &cfg)

	if err != nil {
		return config.Config{}, err
	}

	return cfg, nil
}
