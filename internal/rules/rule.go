package rules

import (
	"cfgscan/internal/config"
	"cfgscan/internal/issue"
)

type Rule interface {
	Check(cfg config.Config) []issue.Issue
}
