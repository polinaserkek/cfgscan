package parser

import (
	"testing"
)

func TestYAMLParser_CorrectData(t *testing.T) {
	data := []byte(`
		log: 
			level: debug
	`)

	parser := YAMLParser{}

	cfg, err := parser.Parse(data)

	if err != nil {
		t.Fatal(err)
	}

	if cfg.Log.Level != "debug" {
		t.Errorf("ожидался уровень debug, получено %s", cfg.Log.Level)
	}

}

func TestYAMLParser_IncorrectData(t *testing.T) {
	data := []byte(`log: `)

	parser := YAMLParser{}

	_, err := parser.Parse(data)

	if err == nil {
		t.Fatal("ожидалась ошибка при некорректном YAML")
	}
}

func TestYAMLParser_EmptyData(t *testing.T) {
	data := []byte{}

	parser := JSONParser{}

	_, err := parser.Parse(data)
	if err == nil {
		t.Fatal("ожидалась ошибка при пустом YAML")
	}
}
