package rules

import (
	"cfgscan/internal/config"
	"cfgscan/internal/issue"
	"testing"
)

func TestTLSRule(t *testing.T) {
	enabled := false
	cfg := config.Config{
		TLS: config.TLSConfig{
			Enabled: &enabled,
		},
	}

	rule := TLSRule{}

	issues := rule.Check(cfg)

	if len(issues) != 1 {
		t.Fatalf("ожидалась 1 проблема, получено %d", len(issues))
	}

	if issues[0].Severity != issue.HIGH {
		t.Errorf("ожидался уровень HIGH, получен %s", issues[0].Severity)
	}

	if issues[0].Message != "TLS отключен." {
		t.Errorf("неверное сообщение: %s", issues[0].Message)
	}
	if issues[0].Recommendation != "Включите TLS для защиты передаваемых данных." {
		t.Errorf("неверная рекомендация: %s", issues[0].Recommendation)
	}
}

func TestTLSRule_Check_NoIssue(t *testing.T) {
	enabled := true

	cfg := config.Config{
		TLS: config.TLSConfig{
			Enabled: &enabled,
		},
	}

	rule := TLSRule{}

	issues := rule.Check(cfg)

	if len(issues) != 0 {
		t.Fatalf("ожидалось отсутствие проблем, получено %d", len(issues))
	}
}

func TestTLSRule_Check_NotConfigured(t *testing.T) {

	cfg := config.Config{}

	rule := TLSRule{}

	issues := rule.Check(cfg)

	if len(issues) != 0 {
		t.Fatalf("ожидалось отсутствие проблем, получено %d", len(issues))
	}
}
