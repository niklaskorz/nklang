package parser

import (
	"fmt"

	"niklaskorz.de/nklang/lexer"
)

type UnexpectedTokenError struct {
	Token *lexer.Token
}

func (e UnexpectedTokenError) Error() string {
	return fmt.Sprintf("Unexpected token at line %d, column %d: %s", e.Token.Line, e.Token.Column, e.Token)
}

func unexpectedToken(token *lexer.Token) UnexpectedTokenError {
	return UnexpectedTokenError{Token: token}
}
