package parser

import (
	"fmt"

	"niklaskorz.de/nklang/lexer"
)

type UnexpectedTokenError struct {
	Token    *lexer.Token
	Expected string
}

func (e UnexpectedTokenError) Error() string {
	return fmt.Sprintf("Unexpected token at line %d, column %d: %s (expected %s)", e.Token.Line, e.Token.Column, e.Token, e.Expected)
}

func unexpectedToken(token *lexer.Token, expected string) UnexpectedTokenError {
	return UnexpectedTokenError{Token: token, Expected: expected}
}
