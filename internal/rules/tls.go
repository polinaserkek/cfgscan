package rules

import (
	"cfgscan/internal/config"
	"cfgscan/internal/issue"
)

type TLSRule struct{}

func (t TLSRule) Check(cfg config.Config) []issue.Issue {
	if cfg.TLS.Enabled != nil && !*cfg.TLS.Enabled {
		return []issue.Issue{
			{
				Severity:       issue.HIGH,
				Message:        "TLS отключен.",
				Recommendation: "Включите TLS для защиты передаваемых данных.",
			},
		}
	}
	return nil
}
