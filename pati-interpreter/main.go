// main.go
package main

import (
	"fmt"
	"os"
	"pati/interpreter"
	"pati/parser"
	"pati/patistructs"
	"pati/tokenizer"
)

// ErrorHandler implementation for the interpreter
type SimpleErrorHandler struct {
	errorCode int
	line      int
}

func (eh *SimpleErrorHandler) SetCode(errorCode int, line int) {
	eh.errorCode = errorCode
	eh.line = line
}

func (eh *SimpleErrorHandler) GetCode() int {
	return eh.errorCode
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: pati <file.bas>")
		return
	}

	fileName := os.Args[1]
	content, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Error reading file: %s\n", err)
		return
	}

	// Tokenize the content of the BASIC file
	tokens := tokenizer.Tokenize(string(content))

	// Set up the error handler
	errorHandler := &SimpleErrorHandler{}

	// Parse the tokens to create a ProgramNode
	programParser := parser.NewParser(tokens, errorHandler, &patistructs.LanguageOptions{})
	program := programParser.ParseProgram()

	if errorHandler.GetCode() != 0 {
		fmt.Printf("Parsing error at line %d: error code %d\n", errorHandler.line, errorHandler.GetCode())
		return
	}

	// Create a new instance of the interpreter
	basicInterpreter := interpreter.NewInterpreter(errorHandler)

	// Run the parsed program
	basicInterpreter.RunProgram(program)

	if errorHandler.GetCode() != 0 {
		fmt.Printf("Runtime error at line %d: error code %d\n", errorHandler.line, errorHandler.GetCode())
	}
}
