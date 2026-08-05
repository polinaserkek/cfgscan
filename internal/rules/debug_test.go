package rules

import (
	"cfgscan/internal/config"
	"cfgscan/internal/issue"
	"testing"
)

func TestDebugRule(t *testing.T) {
	cfg := config.Config{
		Log: config.LogConfig{
			Level: "debug",
		},
	}

	rule := DebugRule{}

	issues := rule.Check(cfg)

	if len(issues) != 1 {
		t.Fatalf("ожидалась 1 проблема, получено %d", len(issues))
	}

	if issues[0].Severity != issue.LOW {
		t.Errorf("ожидался уровень LOW, получен %s", issues[0].Severity)
	}

	if issues[0].Message != "логирование в debug-режиме." {
		t.Errorf("неверное сообщение: %s", issues[0].Message)
	}
	if issues[0].Recommendation != "Поменяйте режим на более избирательный (info+)." {
		t.Errorf("неверная рекомендация: %s", issues[0].Recommendation)
	}
}

func TestDebugRule_Check_NoIssue(t *testing.T) {

	cfg := config.Config{
		Log: config.LogConfig{
			Level: "info",
		},
	}

	rule := DebugRule{}

	issues := rule.Check(cfg)

	if len(issues) != 0 {
		t.Fatalf("ожидалось отсутствие проблем, получено %d", len(issues))
	}
}
