package rules

import (
	"cfgscan/internal/config"
	"cfgscan/internal/issue"
	"testing"
)

func TestOpenBindAddressRule(t *testing.T) {
	cfg := config.Config{
		Network: config.NetworkConfig{
			Address: "0.0.0.0",
		},
	}

	rule := OpenBindAddressRule{}

	issues := rule.Check(cfg)

	if len(issues) != 1 {
		t.Fatalf("ожидалась 1 проблема, получено %d", len(issues))
	}

	if issues[0].Severity != issue.MEDIUM {
		t.Errorf("ожидался уровень MEDIUM, получен %s", issues[0].Severity)
	}

	if issues[0].Message != "приложение принимает подключения со всех сетевых интерфейсов." {
		t.Errorf("неверное сообщение: %s", issues[0].Message)
	}
	if issues[0].Recommendation != "Ограничьте адрес прослушивания или настройте фильтрацию доступа." {
		t.Errorf("неверная рекомендация: %s", issues[0].Recommendation)
	}
}

func TestOpenBindAddressRule_Check_NoIssue(t *testing.T) {

	cfg := config.Config{
		Network: config.NetworkConfig{
			Address: "192.168.0.120",
		},
	}

	rule := OpenBindAddressRule{}

	issues := rule.Check(cfg)

	if len(issues) != 0 {
		t.Fatalf("ожидалось отсутствие проблем, получено %d", len(issues))
	}
}
