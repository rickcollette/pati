// /home/megalith/pati/patistructs/patistructs.go
package patistructs

// ErrorHandler interface to handle errors
type ErrorHandler interface {
	SetCode(errorCode int, line int)
	GetCode() int
}

// LanguageOptions struct for compiler options
type LanguageOptions struct {
	CommentsEnabled bool
	GosubLimit      int
}

// TokenClass enumerates the types of tokens
type TokenClass int

const (
	// List of TokenClass constants
	TOKEN_NONE TokenClass = iota
	TOKEN_EOF
	TOKEN_EOL
	TOKEN_WORD
	TOKEN_STRING
	TOKEN_LET
	TOKEN_IF
	TOKEN_THEN
	TOKEN_RETURN
	TOKEN_END
	TOKEN_PRINT
	TOKEN_INPUT
	TOKEN_REM
	TOKEN_VARIABLE
	TOKEN_PLUS
	TOKEN_MINUS
	TOKEN_MULTIPLY
	TOKEN_NUMBER
	TOKEN_SEMICOLON
	TOKEN_DIVIDE
	TOKEN_LEFT_PARENTHESIS
	TOKEN_RIGHT_PARENTHESIS
	TOKEN_LEFT_BRACE
	TOKEN_RIGHT_BRACE
	TOKEN_EQUAL
	TOKEN_UNEQUAL
	TOKEN_LESSTHAN
	TOKEN_LESSOREQUAL
	TOKEN_GREATERTHAN
	TOKEN_GREATEROREQUAL
	TOKEN_COMMA
	TOKEN_ILLEGAL
)

// Token struct
type Token struct {
	Class   TokenClass
	Line    int
	Pos     int
	Content string
}

// NewToken creates a new Token without initialization
func NewToken() *Token {
	return &Token{
		Class: TOKEN_NONE,
		Line:  0,
		Pos:   0,
	}
}

// NewTokenWithValues creates a new Token with initialization
func NewTokenWithValues(class TokenClass, line, pos int, content string) *Token {
	return &Token{
		Class:   class,
		Line:    line,
		Pos:     pos,
		Content: content,
	}
}

// RelationalOperator enumerates the types of relational operators
type RelationalOperator int

const (
	RELOP_EQUAL RelationalOperator = iota
	RELOP_UNEQUAL
	RELOP_LESSTHAN
	RELOP_LESSOREQUAL
	RELOP_GREATERTHAN
	RELOP_GREATEROREQUAL
)

// FactorClass enumerates the types of factors
type FactorClass int

const (
	FACTOR_NONE FactorClass = iota
	FACTOR_VARIABLE
	FACTOR_VALUE
	FACTOR_EXPRESSION
)

// FactorNode struct
type FactorNode struct {
	Class      FactorClass
	Sign       int
	Variable   int
	Value      int
	Expression *ExpressionNode
}

// ExpressionNode struct
type ExpressionNode struct {
	Term *TermNode
	Next *RightHandTerm
}

// TermNode struct
type TermNode struct {
	Factor *FactorNode
	Next   *RightHandFactor
}

// RightHandFactor struct
type RightHandFactor struct {
	Op     TermOperator
	Factor *FactorNode
	Next   *RightHandFactor
}

// RightHandTerm struct
type RightHandTerm struct {
	Op   ExpressionOperator
	Term *TermNode
	Next *RightHandTerm
}

// TermOperator enumerates the types of term operators
type TermOperator int

const (
	TERM_OPERATOR_NONE TermOperator = iota
	TERM_OPERATOR_MULTIPLY
	TERM_OPERATOR_DIVIDE
)

// ExpressionOperator enumerates the types of expression operators
type ExpressionOperator int

const (
	EXPRESSION_OPERATOR_NONE ExpressionOperator = iota
	EXPRESSION_OPERATOR_PLUS
	EXPRESSION_OPERATOR_MINUS
)

// StatementClass enumerates the types of statements
type StatementClass int

const (
	STATEMENT_NONE StatementClass = iota
	STATEMENT_LET
	STATEMENT_IF
	STATEMENT_RETURN
	STATEMENT_END
	STATEMENT_PRINT
	STATEMENT_INPUT
	STATEMENT_CALL
)

// LetStatementNode struct
type LetStatementNode struct {
	Variable   int
	Expression *ExpressionNode
}

// IfStatementNode struct
type IfStatementNode struct {
	Left      *ExpressionNode
	Op        RelationalOperator
	Right     *ExpressionNode
	Statement *StatementNode
}

// PrintStatementNode struct
type PrintStatementNode struct {
	First *OutputNode
}

// InputStatementNode struct
type InputStatementNode struct {
	First *VariableListNode
}

// ArgumentNode struct for procedure arguments
type ArgumentNode struct {
	Name  string // Name of the argument
	Type  string // Type of the argument (e.g., "string", "int")
	Value interface{} // Value of the argument
}

// StatementNode struct
type StatementNode struct {
	Class     StatementClass
	LetNode   *LetStatementNode
	IfNode    *IfStatementNode
	PrintNode *PrintStatementNode
	InputNode *InputStatementNode
	CallName  string          // Name of the procedure to CALL
	Arguments []*ArgumentNode // Arguments passed to the procedure
}

// ProgramNode struct
type ProgramNode struct {
	Procedures map[string]*ProgramLineNode // Map of named procedures
	Main       *ProgramLineNode            // Entry point of the main program
}

// ProgramLineNode struct
type ProgramLineNode struct {
	ProcedureName string            // Name of the procedure (if applicable)
	Statement     *StatementNode    // Statement in the line
	Next          *ProgramLineNode  // Next line in the procedure or main program
}

// OutputNode struct
type OutputNode struct {
	Value string // Output value (e.g., variables or strings)
}

// VariableListNode struct
type VariableListNode struct {
	Variables []int // List of variables
}
