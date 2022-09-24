package main

import (
	_ "embed"
	"fmt"

	"github.com/yngvark.com/kjoremiljo-convert-csv-to-issues/pkg/convert"
)

//go:embed issues.csv
var csv string

func main() {
	err := run()
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
}

func run() error {
	_, err := convert.CreateGhCommands(csv, "yngvark-temp")
	if err != nil {
		return fmt.Errorf("creating gh commands: %w", err)
	}

	return nil
}
