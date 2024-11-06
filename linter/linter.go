package linter

import (
	"fmt"
	"pati/patistructs"
	"pati/tokenizer"
)

// Linter struct to hold linter information and warnings
type Linter struct {
	warnings         []string
	declaredVars     map[string]bool // Track declared variables
	usedVars         map[string]bool // Track used variables
	procedureNames   map[string]bool // Track declared procedure names
	calledProcedures map[string]bool // Track called procedure names
}

// NewLinter creates a new instance of the Linter
func NewLinter() *Linter {
	return &Linter{
		warnings:         []string{},
		declaredVars:     make(map[string]bool),
		usedVars:         make(map[string]bool),
		procedureNames:   make(map[string]bool),
		calledProcedures: make(map[string]bool),
	}
}

// Lint checks the program content for issues and returns warnings
func (l *Linter) Lint(content string) []string {
	tokens := tokenizer.Tokenize(content)
	l.checkSyntax(tokens)
	l.checkVariableUsage()
	l.checkProcedureDeclarations()
	l.checkKeywordUsage(tokens)
	l.checkUnreachableCode(tokens)
	l.checkTypeMismatch(tokens)

	return l.warnings
}

// checkSyntax analyzes the tokens for syntax issues and captures variables and procedures
func (l *Linter) checkSyntax(tokens []*patistructs.Token) {
	var braceCount int
	var lastToken *patistructs.Token

	for _, token := range tokens {
		switch token.Class {
		case patistructs.TOKEN_LEFT_BRACE:
			braceCount++
		case patistructs.TOKEN_RIGHT_BRACE:
			braceCount--
			if braceCount < 0 {
				l.warnings = append(l.warnings, fmt.Sprintf("Unmatched '}' at line %d", token.Line))
				braceCount = 0
			}
		case patistructs.TOKEN_LET:
			// Next token should be a variable
			lastToken = token
		case patistructs.TOKEN_VARIABLE:
			if lastToken != nil && lastToken.Class == patistructs.TOKEN_LET {
				// Mark variable as declared
				l.declaredVars[token.Content] = true
			} else {
				// Mark variable as used
				l.usedVars[token.Content] = true
			}
			lastToken = nil
		case patistructs.TOKEN_WORD:
			if token.Content == "PROC" {
				lastToken = token
			} else if lastToken != nil && lastToken.Content == "PROC" {
				// Capture procedure name
				l.procedureNames[token.Content] = true
			} else if token.Content == "CALL" {
				lastToken = token
			} else if lastToken != nil && lastToken.Content == "CALL" {
				// Capture called procedure name
				l.calledProcedures[token.Content] = true
			}
		}
	}

	if braceCount > 0 {
		l.warnings = append(l.warnings, "Unmatched '{' found in the code")
	}
}

// checkVariableUsage checks for undeclared or unused variables
func (l *Linter) checkVariableUsage() {
	// Check for undeclared variables
	for variable := range l.usedVars {
		if !l.declaredVars[variable] {
			l.warnings = append(l.warnings, fmt.Sprintf("Variable '%s' is used but not declared", variable))
		}
	}

	// Check for unused variables
	for variable := range l.declaredVars {
		if !l.usedVars[variable] {
			l.warnings = append(l.warnings, fmt.Sprintf("Variable '%s' is declared but not used", variable))
		}
	}
}

// checkProcedureDeclarations checks for undeclared or misused procedures
func (l *Linter) checkProcedureDeclarations() {
	// Check for calls to undeclared procedures
	for procedure := range l.calledProcedures {
		if !l.procedureNames[procedure] {
			l.warnings = append(l.warnings, fmt.Sprintf("Procedure '%s' is called but not declared", procedure))
		}
	}

	// Check for unused procedures
	for procedure := range l.procedureNames {
		if !l.calledProcedures[procedure] {
			l.warnings = append(l.warnings, fmt.Sprintf("Procedure '%s' is declared but never called", procedure))
		}
	}
}

// checkKeywordUsage ensures proper usage of keywords (e.g., IF-THEN structure)
func (l *Linter) checkKeywordUsage(tokens []*patistructs.Token) {
	var expectThen bool

	for _, token := range tokens {
		if token.Class == patistructs.TOKEN_IF {
			expectThen = true
		} else if token.Class == patistructs.TOKEN_THEN {
			if !expectThen {
				l.warnings = append(l.warnings, fmt.Sprintf("'THEN' found without a preceding 'IF' at line %d", token.Line))
			}
			expectThen = false
		}
	}

	if expectThen {
		l.warnings = append(l.warnings, "Missing 'THEN' after 'IF' statement")
	}
}

// checkUnreachableCode analyzes the program flow for unreachable code
func (l *Linter) checkUnreachableCode(tokens []*patistructs.Token) {
	var endReached bool

	for _, token := range tokens {
		if token.Class == patistructs.TOKEN_END {
			endReached = true
		} else if endReached {
			l.warnings = append(l.warnings, fmt.Sprintf("Unreachable code detected at line %d", token.Line))
			break
		}
	}
}
// checkTypeMismatch checks for potential type mismatches in the program
func (l *Linter) checkTypeMismatch(tokens []*patistructs.Token) {
	// Map to keep track of variable types
	// The type could be "int", "string", or any other data type supported in the BASIC program
	variableTypes := make(map[string]string)

	var currentAssignmentType string
	var currentVariable string
	var lastToken *patistructs.Token

	for _, token := range tokens {
		switch token.Class {
		case patistructs.TOKEN_LET:
			// If we see a LET statement, the next TOKEN_VARIABLE will be the variable being assigned
			lastToken = token
		case patistructs.TOKEN_VARIABLE:
			if lastToken != nil && lastToken.Class == patistructs.TOKEN_LET {
				// This variable is being assigned a value
				currentVariable = token.Content
				currentAssignmentType = "" // Reset the current assignment type for this variable
				lastToken = nil
			} else {
				// This variable is being used in an expression
				varName := token.Content
				if expectedType, exists := variableTypes[varName]; exists {
					// Check if the current assignment type matches the expected type
					if currentAssignmentType != "" && currentAssignmentType != expectedType {
						l.warnings = append(l.warnings, fmt.Sprintf("Type mismatch: Variable '%s' is expected to be of type '%s' but is used as type '%s' at line %d", varName, expectedType, currentAssignmentType, token.Line))
					}
				}
			}
		case patistructs.TOKEN_NUMBER:
			// If we encounter a number, it implies an integer assignment
			currentAssignmentType = "int"
		case patistructs.TOKEN_STRING:
			// If we encounter a string, it implies a string assignment
			currentAssignmentType = "string"
		case patistructs.TOKEN_EQUAL:
			// If we see an equals sign, the current variable is being assigned a value
			lastToken = token
		case patistructs.TOKEN_PLUS, patistructs.TOKEN_MINUS, patistructs.TOKEN_MULTIPLY, patistructs.TOKEN_DIVIDE:
			// If we see an arithmetic operator, ensure the operands are integers
			if currentAssignmentType != "int" {
				l.warnings = append(l.warnings, fmt.Sprintf("Type mismatch: Arithmetic operations expect integer values, found '%s' at line %d", currentAssignmentType, token.Line))
			}
		}

		// If we are at the end of an assignment statement, store the type of the variable
		if token.Class == patistructs.TOKEN_EOL || token.Class == patistructs.TOKEN_SEMICOLON {
			if currentVariable != "" && currentAssignmentType != "" {
				if existingType, exists := variableTypes[currentVariable]; exists {
					// Check if the type is consistent with previous assignments
					if existingType != currentAssignmentType {
						l.warnings = append(l.warnings, fmt.Sprintf("Type mismatch: Variable '%s' was previously assigned as '%s' but now as '%s' at line %d", currentVariable, existingType, currentAssignmentType, token.Line))
					}
				} else {
					// Record the type of the variable if it's the first time it's assigned
					variableTypes[currentVariable] = currentAssignmentType
				}
			}
			// Reset for the next statement
			currentAssignmentType = ""
			currentVariable = ""
		}
	}
}
