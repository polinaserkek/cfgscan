package parser

import (
	"cfgscan/internal/config"
	"errors"
)

func DetectParser(data []byte) (config.Config, error) {
	jp := JSONParser{}
	yp := YAMLParser{}

	cfg, err := jp.Parse(data)
	if err == nil {
		return cfg, nil
	}
	cfg, err = yp.Parse(data)
	if err == nil {
		return cfg, nil
	}

	return config.Config{}, errors.New("поддерживаются только JSON и YAML")
}
