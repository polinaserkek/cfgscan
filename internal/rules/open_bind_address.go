package rules

import (
	"cfgscan/internal/config"
	"cfgscan/internal/issue"
)

type OpenBindAddressRule struct{}

func (o OpenBindAddressRule) Check(cfg config.Config) []issue.Issue {
	if cfg.Network.Address == "0.0.0.0" {
		return []issue.Issue{
			{
				Severity:       issue.MEDIUM,
				Message:        "приложение принимает подключения со всех сетевых интерфейсов.",
				Recommendation: "Ограничьте адрес прослушивания или настройте фильтрацию доступа.",
			},
		}
	}
	return nil
}
