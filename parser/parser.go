package parser

import (
	"pati/patistructs"
)
// Parser struct
type Parser struct {
	tokens     []*patistructs.Token
	currentPos int
	errors     patistructs.ErrorHandler
	options    *patistructs.LanguageOptions
}

// NewParser creates a new Parser instance
func NewParser(tokens []*patistructs.Token, errors patistructs.ErrorHandler, options *patistructs.LanguageOptions) *Parser {
	return &Parser{
		tokens:     tokens,
		currentPos: 0,
		errors:     errors,
		options:    options,
	}
}

// Helper function to get the current token
func (p *Parser) currentToken() *patistructs.Token {
	if p.currentPos < len(p.tokens) {
		return p.tokens[p.currentPos]
	}
	return &patistructs.Token{Class: patistructs.TOKEN_EOF}
}

// Helper function to advance to the next token
func (p *Parser) advance() {
	if p.currentPos < len(p.tokens) {
		p.currentPos++
	}
}

// ParseProgram parses an entire BASIC program
func (p *Parser) ParseProgram() *patistructs.ProgramNode {
	program := &patistructs.ProgramNode{
		Procedures: make(map[string]*patistructs.ProgramLineNode),
	}

	for p.currentToken().Class != patistructs.TOKEN_EOF {
		if p.currentToken().Class == patistructs.TOKEN_WORD && p.currentToken().Content == "PROC" {
			// Parse a named procedure
			p.advance() // Move past "PROC"
			nameToken := p.currentToken()
			if nameToken.Class != patistructs.TOKEN_WORD {
				p.errors.SetCode(16, nameToken.Line) // Error: Expected procedure name
				return nil
			}
			procedureName := nameToken.Content
			p.advance() // Move past the procedure name

			if p.currentToken().Class != patistructs.TOKEN_LEFT_BRACE {
				p.errors.SetCode(17, nameToken.Line) // Error: Expected '{'
				return nil
			}
			p.advance() // Move past '{'

			procedureLine := p.parseProgramLinesUntilRightBrace()
			if procedureLine != nil {
				program.Procedures[procedureName] = procedureLine
			}
		} else {
			// Parse the main program
			line := p.parseProgramLine()
			if line != nil {
				if program.Main == nil {
					program.Main = line
				} else {
					current := program.Main
					for current.Next != nil {
						current = current.Next
					}
					current.Next = line
				}
			}
		}
	}

	return program
}

// Helper to parse lines until a right brace is found
func (p *Parser) parseProgramLinesUntilRightBrace() *patistructs.ProgramLineNode {
	var head, current *patistructs.ProgramLineNode

	for p.currentToken().Class != patistructs.TOKEN_RIGHT_BRACE && p.currentToken().Class != patistructs.TOKEN_EOF {
		line := p.parseProgramLine()
		if line != nil {
			if head == nil {
				head = line
			} else {
				current.Next = line
			}
			current = line
		}
	}

	if p.currentToken().Class == patistructs.TOKEN_RIGHT_BRACE {
		p.advance() // Move past '}'
	} else {
		p.errors.SetCode(18, 0) // Error: Expected '}'
	}

	return head
}


func (p *Parser) parseProgramLine() *patistructs.ProgramLineNode {
	token := p.currentToken()
	if token.Class != patistructs.TOKEN_WORD {
		p.errors.SetCode(1, token.Line) // Error: Expected procedure name
		return nil
	}

	lineNode := &patistructs.ProgramLineNode{
		ProcedureName: token.Content,
	}
	p.advance() // Move past the procedure name

	statement := p.parseStatement()
	if statement != nil {
		lineNode.Statement = statement
	}
	return lineNode
}

// ParseStatement parses a single statement
func (p *Parser) parseStatement() *patistructs.StatementNode {
	token := p.currentToken()
	switch token.Class {
	case patistructs.TOKEN_LET:
		return p.parseLetStatement()
	case patistructs.TOKEN_IF:
		return p.parseIfStatement()
	case patistructs.TOKEN_PRINT:
		return p.parsePrintStatement()
	case patistructs.TOKEN_INPUT:
		return p.parseInputStatement()
	case patistructs.TOKEN_WORD:
		if token.Content == "CALL" {
			return p.parseCallStatement()
		}
	default:
		p.errors.SetCode(2, token.Line) // Example error code for unrecognized statement
		return nil
	}
		// Ensure a return value for all cases
		p.errors.SetCode(2, token.Line) // Example error code for unrecognized statement
		return nil
}

// Parse a CALL statement
func (p *Parser) parseCallStatement() *patistructs.StatementNode {
	p.advance() // Move past the CALL token

	nameToken := p.currentToken()
	if nameToken.Class != patistructs.TOKEN_WORD {
		p.errors.SetCode(19, nameToken.Line) // Error: Expected procedure name
		return nil
	}
	callName := nameToken.Content
	p.advance() // Move past the procedure name

	// Parse arguments (if any)
	var arguments []*patistructs.ArgumentNode
	if p.currentToken().Class == patistructs.TOKEN_LEFT_PARENTHESIS {
		p.advance() // Move past '('
		for p.currentToken().Class != patistructs.TOKEN_RIGHT_PARENTHESIS {
			argToken := p.currentToken()
			argName := argToken.Content
			var argValue interface{} = argToken.Content // Simplified parsing, handle types appropriately
			arguments = append(arguments, &patistructs.ArgumentNode{Name: argName, Value: argValue})
			p.advance() // Move past the argument

			if p.currentToken().Class == patistructs.TOKEN_COMMA {
				p.advance() // Move past ','
			}
		}
		p.advance() // Move past ')'
	}

	return &patistructs.StatementNode{
		Class:     patistructs.STATEMENT_CALL,
		CallName:  callName,
		Arguments: arguments,
	}
}

// Parse a LET statement
func (p *Parser) parseLetStatement() *patistructs.StatementNode {
	p.advance() // Move past the LET token

	token := p.currentToken()
	if token.Class != patistructs.TOKEN_VARIABLE {
		p.errors.SetCode(3, token.Line) // Error: Expected variable
		return nil
	}

	letNode := &patistructs.LetStatementNode{
		Variable: int(token.Content[0] - 'A'), // Convert 'A' to 0, 'B' to 1, etc.
	}
	p.advance() // Move past the variable

	if p.currentToken().Class != patistructs.TOKEN_EQUAL {
		p.errors.SetCode(4, token.Line) // Error: Expected '='
		return nil
	}
	p.advance() // Move past the '='

	letNode.Expression = p.parseExpression()
	return &patistructs.StatementNode{
		Class:   patistructs.STATEMENT_LET,
		LetNode: letNode,
	}
}

// Parse an IF statement
func (p *Parser) parseIfStatement() *patistructs.StatementNode {
	p.advance() // Move past the IF token

	ifNode := &patistructs.IfStatementNode{
		Left: p.parseExpression(),
	}
	if ifNode.Left == nil {
		return nil
	}

	token := p.currentToken()
	if !p.isRelationalOperator(token.Class) {
		p.errors.SetCode(5, token.Line) // Error: Expected relational operator
		return nil
	}
	ifNode.Op = p.getRelationalOperator(token.Class)
	p.advance() // Move past the operator

	ifNode.Right = p.parseExpression()
	if ifNode.Right == nil {
		return nil
	}

	if p.currentToken().Class != patistructs.TOKEN_THEN {
		p.errors.SetCode(6, token.Line) // Error: Expected THEN
		return nil
	}
	p.advance() // Move past the THEN token

	ifNode.Statement = p.parseStatement()
	return &patistructs.StatementNode{
		Class:    patistructs.STATEMENT_IF,
		IfNode:   ifNode,
	}
}

// Helper function to check for relational operators
func (p *Parser) isRelationalOperator(class patistructs.TokenClass) bool {
	return class == patistructs.TOKEN_EQUAL || class == patistructs.TOKEN_UNEQUAL || class == patistructs.TOKEN_LESSTHAN ||
		class == patistructs.TOKEN_LESSOREQUAL || class == patistructs.TOKEN_GREATERTHAN || class == patistructs.TOKEN_GREATEROREQUAL
}

// Helper function to map tokens to RelationalOperator
func (p *Parser) getRelationalOperator(class patistructs.TokenClass) patistructs.RelationalOperator {
	switch class {
	case patistructs.TOKEN_EQUAL:
		return patistructs.RELOP_EQUAL
	case patistructs.TOKEN_UNEQUAL:
		return patistructs.RELOP_UNEQUAL
	case patistructs.TOKEN_LESSTHAN:
		return patistructs.RELOP_LESSTHAN
	case patistructs.TOKEN_LESSOREQUAL:
		return patistructs.RELOP_LESSOREQUAL
	case patistructs.TOKEN_GREATERTHAN:
		return patistructs.RELOP_GREATERTHAN
	case patistructs.TOKEN_GREATEROREQUAL:
		return patistructs.RELOP_GREATEROREQUAL
	}
	return patistructs.RELOP_EQUAL // Default case, though it shouldn't occur
}

// Parse an expression (placeholder for now)
func (p *Parser) parseExpression() *patistructs.ExpressionNode {
	// Implement expression parsing logic
	return &patistructs.ExpressionNode{}
}

// Parse a PRINT statement (simplified for now)
func (p *Parser) parsePrintStatement() *patistructs.StatementNode {
	p.advance() // Move past the PRINT token

	printNode := &patistructs.PrintStatementNode{}
	// Logic to parse expressions or strings to print
	return &patistructs.StatementNode{
		Class:     patistructs.STATEMENT_PRINT,
		PrintNode: printNode,
	}
}

// Parse an INPUT statement (simplified for now)
func (p *Parser) parseInputStatement() *patistructs.StatementNode {
	p.advance() // Move past the INPUT token

	inputNode := &patistructs.InputStatementNode{}
	// Logic to parse variable list for input
	return &patistructs.StatementNode{
		Class:     patistructs.STATEMENT_INPUT,
		InputNode: inputNode,
	}
}
