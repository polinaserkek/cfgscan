package rules

import (
	"cfgscan/internal/config"
	"cfgscan/internal/issue"
	"testing"
)

func TestUnsafeHashAlgorithmRule(t *testing.T) {
	cfg := config.Config{
		Storage: config.StorageConfig{
			DigestAlgorithm: "md5",
		},
	}

	rule := UnsafeHashAlgorithmRule{}

	issues := rule.Check(cfg)

	if len(issues) != 1 {
		t.Fatalf("ожидалась 1 проблема, получено %d", len(issues))
	}

	if issues[0].Severity != issue.HIGH {
		t.Errorf("ожидался уровень HIGH, получен %s", issues[0].Severity)
	}

	if issues[0].Message != "слишком слабый алгоритм - "+cfg.Storage.DigestAlgorithm+"." {
		t.Errorf("неверное сообщение: %s", issues[0].Message)
	}
	if issues[0].Recommendation != "Замените его на более безопасный." {
		t.Errorf("неверная рекомендация: %s", issues[0].Recommendation)
	}
}

func TestUnsafeHashAlgorithmRule_Check_NoIssue(t *testing.T) {

	cfg := config.Config{
		Storage: config.StorageConfig{
			DigestAlgorithm: "sha256",
		},
	}

	rule := UnsafeHashAlgorithmRule{}

	issues := rule.Check(cfg)

	if len(issues) != 0 {
		t.Fatalf("ожидалось отсутствие проблем, получено %d", len(issues))
	}
}
