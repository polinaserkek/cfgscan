package parser

import (
	"cfgscan/internal/config"
)

type Parser interface {
	Parse([]byte) (config.Config, error)
}

func ParseBytes(p Parser, file []byte) {
	p.Parse(file)
}
