package rules

import (
	"cfgscan/internal/config"
	"cfgscan/internal/issue"
	"strings"
)

type DebugRule struct{}

func (d DebugRule) Check(cfg config.Config) []issue.Issue {
	if strings.EqualFold(cfg.Log.Level, "debug") {
		return []issue.Issue{
			{
				Severity:       issue.LOW,
				Message:        "логирование в debug-режиме.",
				Recommendation: "Поменяйте режим на более избирательный (info+).",
			},
		}
	}

	return nil
}
