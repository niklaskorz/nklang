package ast

type Program struct {
	Statements []Statement
}

func (p Program) Execute() error {
	for _, s := range p.Statements {
		if err := s.Evaluate(); err != nil {
			return err
		}
	}
	return nil
}
