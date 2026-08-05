package parser

import (
	"testing"
)

func TestJSONParser_CorrectData(t *testing.T) {
	data := []byte(`{
		"log": {
			"level": "debug"
		}
	}`)

	parser := JSONParser{}

	cfg, err := parser.Parse(data)

	if err != nil {
		t.Fatal(err)
	}

	if cfg.Log.Level != "debug" {
		t.Errorf("ожидался уровень debug, получено %s", cfg.Log.Level)
	}

}

func TestJSONParser_IncorrectData(t *testing.T) {
	data := []byte(`{"log": {`)

	parser := JSONParser{}

	_, err := parser.Parse(data)

	if err == nil {
		t.Fatal("ожидалась ошибка при некорректном JSON")
	}
}

func TestJSONParser_EmptyData(t *testing.T) {
	data := []byte{}

	parser := JSONParser{}

	_, err := parser.Parse(data)
	if err == nil {
		t.Fatal("ожидалась ошибка при пустом JSON")
	}
}
