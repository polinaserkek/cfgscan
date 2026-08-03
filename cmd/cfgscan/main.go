package main

import (
	"cfgscan/internal/parser"
	"cfgscan/internal/scanner"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	path, data, err := readInput()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(data))

	p, err := parser.SelectParser(path)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(p)

	cfg, err := p.Parse(data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", cfg)
	fmt.Printf("%v\n", cfg)

	sc := scanner.New()

	issues := sc.Scan(cfg)

	fmt.Printf("%v\n", issues)
}

func readInput() (string, []byte, error) {
	if len(os.Args) < 2 {
		return "", nil, errors.New("no file")
	}

	if os.Args[1] == "--stdin" {
		data, err := io.ReadAll(os.Stdin)
		if err != nil {
			return "", nil, err
		}
		return "", data, nil

	}
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		return "", nil, err
	}
	return os.Args[1], data, nil

}
