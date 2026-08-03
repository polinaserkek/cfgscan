package parser

import (
	"fmt"
	"path/filepath"
)

func SelectParser(path string) (Parser, error) {
	ext := filepath.Ext(path)

	switch ext {
	case ".json":
		return JSONParser{}, nil
	case ".yaml", ".yml":
		return YAMLParser{}, nil
	default:
		return nil, fmt.Errorf("такой формат на поддерживается: %s", ext)
	}

}
