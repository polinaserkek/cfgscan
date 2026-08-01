package main

import (
	"errors"
	"io"
	"log"
	"os"
)

func main() {
	if _, err := readInput(); err != nil {
		log.Fatal(err)
	}
}

func readInput() ([]byte, error) {
	if len(os.Args) < 2 {
		return nil, errors.New("config file path is required")
	}

	if os.Args[1] == "--stdin" {
		data, err := io.ReadAll(os.Stdin)
		if err != nil {
			return nil, err
		}
		return data, nil

	}
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		return nil, err
	}
	return data, nil

}
