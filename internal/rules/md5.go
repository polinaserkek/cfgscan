package rules

import (
	"cfgscan/internal/config"
	"cfgscan/internal/issue"
	"strings"
)

type MD5Rule struct{}

func (m MD5Rule) Check(cfg config.Config) []issue.Issue {
	if strings.EqualFold(cfg.Storage.DigestAlgorithm, "md5") {
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
