package rules

import (
	"cfgscan/internal/config"
	"cfgscan/internal/issue"
)

func CheckDebug(cfg config.Config) []issue.Issue {
	if cfg.Log.Level == "debug" {
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
