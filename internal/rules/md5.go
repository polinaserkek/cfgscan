package rules

import (
	"cfgscan/internal/config"
	"cfgscan/internal/issue"
)

func CheckMd5(cfg config.Config) []issue.Issue {
	if cfg.Storage.DigestAlgorithm == "md5" {
		return []issue.Issue{
			{
				Severity:       issue.HIGH,
				Message:        "слишком слабый алгоритм - MD5.",
				Recommendation: "Замените его на более безопасный.",
			},
		}
	}
	return nil
}
