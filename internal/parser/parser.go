package parser

import (
	"cfgscan/internal/config"
)

type Parser interface {
	Parse([]byte) (config.Config, error)
}
