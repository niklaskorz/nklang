package ast

import (
	"fmt"
)

type Program struct {
	Statements []Statement
}

func (p *Program) String() string {
	return fmt.Sprintf("Program{Statements: %s}", p.Statements)
}
