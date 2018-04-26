package ast

type Program struct {
	Statements []Statement
}

func (p Program) Execute() {
	for _, s := range p.Statements {
		s.Evaluate()
	}
}
