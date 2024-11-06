// /home/megalith/pati/interpreter/interpreter.go
package interpreter

import (
	"fmt"
	"pati/patistructs"
	"strconv" // Added for converting int to string
)

// Interpreter struct to maintain the state of the interpreter
type Interpreter struct {
	variables    map[string]interface{}             // Use a map for variables to support different types
	errors       patistructs.ErrorHandler           // Error handler
	currentLine  *patistructs.ProgramLineNode       // Current line for RETURN
	lineStack    []*patistructs.ProgramLineNode     // Stack for GOSUB and RETURN
	procedures   map[string]*patistructs.ProgramLineNode // Map of procedure names to nodes
}

// NewInterpreter creates a new instance of the interpreter
func NewInterpreter(errors patistructs.ErrorHandler) *Interpreter {
	return &Interpreter{
		variables:    make(map[string]interface{}),
		errors:       errors,
		lineStack:    []*patistructs.ProgramLineNode{},
		procedures:   make(map[string]*patistructs.ProgramLineNode),
	}
}

func (i *Interpreter) RunProgram(program *patistructs.ProgramNode) {
	// Initialize procedures
	i.procedures = program.Procedures
	i.currentLine = program.Main

	// Execute the main program
	for i.currentLine != nil {
		i.executeStatement(i.currentLine.Statement)
		i.currentLine = i.currentLine.Next
	}
}

// Execute a single statement
func (i *Interpreter) executeStatement(statement *patistructs.StatementNode) {
	if statement == nil {
		return
	}

	switch statement.Class {
	case patistructs.STATEMENT_LET:
		i.executeLet(statement.LetNode)
	case patistructs.STATEMENT_IF:
		i.executeIf(statement.IfNode)
	case patistructs.STATEMENT_PRINT:
		i.executePrint(statement.PrintNode)
	case patistructs.STATEMENT_INPUT:
		i.executeInput(statement.InputNode)
	case patistructs.STATEMENT_CALL:
		i.executeCall(statement.CallName, statement.Arguments)
	case patistructs.STATEMENT_RETURN:
		i.executeReturn()
	case patistructs.STATEMENT_END:
		i.executeEnd()
	default:
		i.errors.SetCode(7, 0) // Unrecognized statement error
	}
}

// Expanded evaluateExpression: Handles operators and nested expressions
func (i *Interpreter) evaluateExpression(expr *patistructs.ExpressionNode) int {
	if expr == nil {
		return 0
	}

	termValue := i.evaluateTerm(expr.Term)
	currentTerm := expr.Next
	for currentTerm != nil {
		rightValue := i.evaluateTerm(currentTerm.Term)
		switch currentTerm.Op {
		case patistructs.EXPRESSION_OPERATOR_PLUS:
			termValue += rightValue
		case patistructs.EXPRESSION_OPERATOR_MINUS:
			termValue -= rightValue
		default:
			i.errors.SetCode(9, 0) // Error: Unknown expression operator
		}
		currentTerm = currentTerm.Next
	}
	return termValue
}

// Evaluate a term: Handles multiplication and division
func (i *Interpreter) evaluateTerm(term *patistructs.TermNode) int {
	if term == nil {
		return 0
	}

	factorValue := i.evaluateFactor(term.Factor)
	currentFactor := term.Next
	for currentFactor != nil {
		rightValue := i.evaluateFactor(currentFactor.Factor)
		switch currentFactor.Op {
		case patistructs.TERM_OPERATOR_MULTIPLY:
			factorValue *= rightValue
		case patistructs.TERM_OPERATOR_DIVIDE:
			if rightValue == 0 {
				i.errors.SetCode(10, 0) // Error: Division by zero
				return 0
			}
			factorValue /= rightValue
		default:
			i.errors.SetCode(11, 0) // Error: Unknown term operator
		}
		currentFactor = currentFactor.Next
	}
	return factorValue
}

// Evaluate a factor: Handles variables and values
func (i *Interpreter) evaluateFactor(factor *patistructs.FactorNode) int {
	if factor == nil {
		return 0
	}

	switch factor.Class {
	case patistructs.FACTOR_VALUE:
		return factor.Value
	case patistructs.FACTOR_VARIABLE:
		key := strconv.Itoa(factor.Variable) // Convert int to string for map key
		value, ok := i.variables[key]
		if !ok {
			i.errors.SetCode(13, 0) // Error: Variable not found
			return 0
		}
		intValue, ok := value.(int) // Assert interface{} to int
		if !ok {
			i.errors.SetCode(14, 0) // Error: Variable type mismatch
			return 0
		}
		return intValue
	case patistructs.FACTOR_EXPRESSION:
		return i.evaluateExpression(factor.Expression)
	default:
		i.errors.SetCode(12, 0) // Error: Unknown factor class
		return 0
	}
}

// Execute a LET statement
func (i *Interpreter) executeLet(letNode *patistructs.LetStatementNode) {
	if letNode == nil {
		return
	}

	value := i.evaluateExpression(letNode.Expression)
	key := strconv.Itoa(letNode.Variable) // Convert int to string for map key
	i.variables[key] = value
}

// Execute an IF statement
func (i *Interpreter) executeIf(ifNode *patistructs.IfStatementNode) {
	if ifNode == nil {
		return
	}

	leftValue := i.evaluateExpression(ifNode.Left)
	rightValue := i.evaluateExpression(ifNode.Right)

	conditionMet := false
	switch ifNode.Op {
	case patistructs.RELOP_EQUAL:
		conditionMet = (leftValue == rightValue)
	case patistructs.RELOP_UNEQUAL:
		conditionMet = (leftValue != rightValue)
	case patistructs.RELOP_LESSTHAN:
		conditionMet = (leftValue < rightValue)
	case patistructs.RELOP_LESSOREQUAL:
		conditionMet = (leftValue <= rightValue)
	case patistructs.RELOP_GREATERTHAN:
		conditionMet = (leftValue > rightValue)
	case patistructs.RELOP_GREATEROREQUAL:
		conditionMet = (leftValue >= rightValue)
	}

	if conditionMet {
		i.executeStatement(ifNode.Statement)
	}
}

// Execute a PRINT statement
func (i *Interpreter) executePrint(printNode *patistructs.PrintStatementNode) {
	if printNode == nil || printNode.First == nil {
		return
	}

	fmt.Println(printNode.First.Value) // Simplified to print a single value for now
}

// Execute an INPUT statement
func (i *Interpreter) executeInput(inputNode *patistructs.InputStatementNode) {
	if inputNode == nil || inputNode.First == nil {
		return
	}

	for _, variableIndex := range inputNode.First.Variables {
		var input int
		fmt.Printf("Enter value for variable %c: ", 'A'+variableIndex)
		_, err := fmt.Scan(&input)
		if err != nil {
			i.errors.SetCode(8, 0) // Error handling input
			return
		}
		key := strconv.Itoa(variableIndex) // Convert int to string for map key
		i.variables[key] = input
	}
}

// Execute a CALL statement
func (i *Interpreter) executeCall(name string, arguments []*patistructs.ArgumentNode) {
	procedure, exists := i.procedures[name]
	if !exists {
		i.errors.SetCode(20, 0) // Error: Procedure not found
		return
	}

	// Set arguments as local variables
	for _, arg := range arguments {
		i.variables[arg.Name] = arg.Value
	}

	// Execute the procedure
	i.lineStack = append(i.lineStack, i.currentLine)
	i.currentLine = procedure
}

// Execute a RETURN statement
func (i *Interpreter) executeReturn() {
	if len(i.lineStack) == 0 {
		i.errors.SetCode(15, 0) // Error: No line to return to
		return
	}
	i.currentLine = i.lineStack[len(i.lineStack)-1]
	i.lineStack = i.lineStack[:len(i.lineStack)-1]
}

// Execute an END statement
func (i *Interpreter) executeEnd() {
	i.currentLine = nil // End program execution
}
