// main.go
package main

import (
	"fmt"
	"os"
	"pati/linter"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: pati-linter <file.bas>")
		return
	}

	fileName := os.Args[1]
	content, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Error reading file: %s\n", err)
		return
	}

	// Create a new instance of the linter
	basicLinter := linter.NewLinter()

	// Run the linter on the file content
	warnings := basicLinter.Lint(string(content))

	// Output the warnings, if any
	if len(warnings) == 0 {
		fmt.Println("No issues found.")
	} else {
		fmt.Println("Linter Warnings:")
		for _, warning := range warnings {
			fmt.Println("- " + warning)
		}
	}
}
