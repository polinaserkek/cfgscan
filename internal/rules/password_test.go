package rules

import (
	"cfgscan/internal/config"
	"cfgscan/internal/issue"
	"testing"
)

func TestPasswordRule(t *testing.T) {
	cfg := config.Config{
		Auth: config.AuthConfig{
			Password: "admin111",
		},
	}

	rule := PasswordRule{}
	issues := rule.Check(cfg)

	if len(issues) != 1 {
		t.Fatalf("ожидалась 1 проблема, получено %d", len(issues))

	}

	if issues[0].Severity != issue.HIGH {
		t.Errorf("ожидался уровень HIGH, получен %s", issues[0].Severity)
	}

	if issues[0].Message != "пароль хранится в открытом виде." {
		t.Errorf("неверное сообщение: %s", issues[0].Message)
	}

	if issues[0].Recommendation != "Используйте переменные окружения, менеджер секретов или хеш." {
		t.Errorf("неверная рекомендация: %s", issues[0].Recommendation)
	}
}

func TestPasswordRule_Check_NoIssue(t *testing.T) {

	cfg := config.Config{
		Auth: config.AuthConfig{
			Password: "",
		},
	}

	rule := PasswordRule{}

	issues := rule.Check(cfg)

	if len(issues) != 0 {
		t.Fatalf("ожидалось отсутствие проблем, получено %d", len(issues))
	}
}
