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

func (p *Program) Execute() error {
	for _, s := range p.Statements {
		if err := s.Evaluate(); err != nil {
			return err
		}
	}
	return nil
}
