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
	TrueKeyword                   // true
	FalseKeyword                  // false
	NilKeyword                    // nil
	DeclarationOperator           // :=
	AssignmentOperator            // =
	MulOperator                   // *
	DivOperator                   // /
	AddOperator                   // +
	SubOperator                   // -
	LogicalOr                     // ||
	LogicalAnd                    // &&
	EqOperator                    // ==
	LtOperator                    // <
	LeOperator                    // <=
	GtOperator                    // >
	GeOperator                    // >=
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
	Line, Column int
	Type         TokenType
	Value        string
}

func (t Token) String() string {
	if t.Type == String {
		return "\"" + t.Value + "\""
	}
	if t.Type == EOF {
		return "EOF"
	}
	return t.Value
}
