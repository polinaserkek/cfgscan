package rules

import (
	"cfgscan/internal/config"
	"cfgscan/internal/issue"
)

type PasswordRule struct{}

func (p PasswordRule) Check(cfg config.Config) []issue.Issue {
	if cfg.Auth.Password != "" {
		return []issue.Issue{
			{
				Severity:       issue.HIGH,
				Message:        "пароль хранится в открытом виде.",
				Recommendation: "Используйте переменные окружения, менеджер секретов или хеш.",
			},
		}
	}

	return nil
}
