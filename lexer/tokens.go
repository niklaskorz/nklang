package lexer

type TokenType int

const (
	DeclarationOperator TokenType = iota
	AssignmentOperator
	MultiplicationOperator
	DivisionOperator
	AdditionOperator
	SubstractionOperator
	Semicolon
	Whitespace
	LeftParen
	RightParen
	LeftBrace
	RightBrace
	LeftBracket
	RightBracket
	Identifier
	Integer
	Float
	String
	EOF
)

type Token struct {
	Type  TokenType
	Value string
}
