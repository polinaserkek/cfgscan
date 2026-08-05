package scanner

import (
	"cfgscan/internal/config"
	"testing"
)

func TestScanner(t *testing.T) {
	cfg := config.Config{
		Log: config.LogConfig{
			Level: "debug",
		},
	}

	sc := New()

	issues := sc.Scan(cfg)

	if len(issues) != 1 {
		t.Fatalf("ожидалась 1 проблема, получено %d", len(issues))
	}
}

func TestScanner_DoubleData(t *testing.T) {
	cfg := config.Config{
		Log: config.LogConfig{
			Level: "debug",
		},
		Storage: config.StorageConfig{
			DigestAlgorithm: "md5",
		},
	}

	sc := New()

	issues := sc.Scan(cfg)

	if len(issues) != 2 {
		t.Fatalf("ожидалось 2 проблемы, получено %d", len(issues))
	}
}
