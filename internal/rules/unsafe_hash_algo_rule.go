package rules

import (
	"cfgscan/internal/config"
	"cfgscan/internal/issue"
	"strings"
)

type UnsafeHashAlgorithmRule struct{}

func (uha UnsafeHashAlgorithmRule) Check(cfg config.Config) []issue.Issue {
	unsafeHashAlgos := map[string]bool{
		"md5":   true,
		"md-5":  true,
		"sha1":  true,
		"sha-1": true,
	}

	algorithm := strings.ToLower(cfg.Storage.DigestAlgorithm)

	if unsafeHashAlgos[algorithm] {
		return []issue.Issue{
			{
				Severity:       issue.HIGH,
				Message:        "слишком слабый алгоритм - " + cfg.Storage.DigestAlgorithm + ".",
				Recommendation: "Замените его на более безопасный.",
			},
		}
	}
	return nil
}
