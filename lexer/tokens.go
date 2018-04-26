package lexer

type TokenType int

const (
	EOF                 TokenType = iota
	ReturnKeyword                 // return
	ContinueKeyword               // continue
	BreakKeyword                  // break
	FunctionKeyword               // func
	IfKeyword                     // if
	ElseKeyword                   // else
	WhileKeyword                  // while
	DeclarationOperator           // :=
	AssignmentOperator            // =
	MulOperator                   // *
	DivOperator                   // /
	AddOperator                   // +
	SubOperator                   // -
	Semicolon                     // ;
	Comma                         // ,
	LeftParen                     // (
	RightParen                    // )
	LeftBrace                     // {
	RightBrace                    // }
	LeftBracket                   // [
	RightBracket                  // ]
	ID                            // Unicode letter followed by unicde letters or digits
	Integer                       // Digits
	Float
	String // Arbitrary characters enclosed by quotation marks
)

type Token struct {
	Type  TokenType
	Value string
}
