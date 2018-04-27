package lexer

import (
	"fmt"
)

type UnexpectedSymbolError struct {
	Line, Column int
	Symbol       rune
}

func (e UnexpectedSymbolError) Error() string {
	return fmt.Sprintf("Unexpected symbol at line %d, column %d: %c", e.Line, e.Column, e.Symbol)
}

func (s *Scanner) unexpectedSymbol(symbol rune) UnexpectedSymbolError {
	return UnexpectedSymbolError{Line: s.line, Column: s.column, Symbol: symbol}
}
