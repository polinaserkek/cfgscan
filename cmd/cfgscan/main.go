package main

import (
	"cfgscan/internal/issue"
	"cfgscan/internal/parser"
	"cfgscan/internal/scanner"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	silent := flag.Bool("silent", false, "don't exit with error")
	stdin := flag.Bool("stdin", false, "read config from stdin")

	flag.Parse()

	args := flag.Args()

	if !*stdin && len(args) == 0 {
		log.Fatal("не указан конфигурационный файл")
	}

	var path string
	if len(args) > 0 {
		path = args[0]
	}

	data, err := readInput(path, *stdin)
	if err != nil {
		log.Fatal(err)
	}

	cfg, err := parser.DetectParser(data)
	if err != nil {
		log.Fatal(err)
	}

	sc := scanner.New()

	issues := sc.Scan(cfg)
	printIssues(issues, *silent)
}

func printIssues(issues []issue.Issue, silent bool) {
	if len(issues) > 0 {
		for _, issue := range issues {
			fmt.Printf("%s: %s %s\n",
				issue.Severity,
				issue.Message,
				issue.Recommendation,
			)
		}
		if !silent {
			os.Exit(1)
		}
	}
}

func readInput(path string, stdin bool) ([]byte, error) {
	if stdin {
		return io.ReadAll(os.Stdin)
	}
	return os.ReadFile(path)
}
