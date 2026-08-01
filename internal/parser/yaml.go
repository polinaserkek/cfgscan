package parser

import (
	"cfgscan/internal/config"

	"gopkg.in/yaml.v3"
)

type YAMLParser struct {
}

func (yp YAMLParser) Parse(data []byte) (config.Config, error) {
	// cfg = config.Config{
	// 	Log: config.LogConfig{
	// 		Level: "debug",
	// 	},
	// }

	cfg := config.Config{}

	err := yaml.Unmarshal(data, &cfg)
	if err != nil {
		return config.Config{}, err
	}

	return cfg, nil

}
