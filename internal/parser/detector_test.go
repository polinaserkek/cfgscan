package parser

import "testing"

func TestDetectParser_JSON(t *testing.T) {
	data := []byte(`{
		"log": {
			"level": "debug"
		}
	}`)

	cfg, err := DetectParser(data)

	if err != nil {
		t.Fatal(err)
	}

	if cfg.Log.Level != "debug" {
		t.Errorf("ожидался уровень debug, получено %s", cfg.Log.Level)
	}
}

func TestDetectParser_YAML(t *testing.T) {
	data := []byte(`
log:
  level: debug
`)

	cfg, err := DetectParser(data)

	if err != nil {
		t.Fatal(err)
	}

	if cfg.Log.Level != "debug" {
		t.Errorf("ожидался уровень debug, получено %s", cfg.Log.Level)
	}
}
