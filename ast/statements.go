package ast

type Statement interface {
	Evaluate()
}

type Declaration struct {
	Identifier string
	Value      Expression
}

func (d Declaration) Evaluate() {
	// TODO: Implement declaration
	d.Value.Evaluate()
}

type Assignment struct {
	Identifier string
	Value      Expression
}

func (a Assignment) Evaluate() {
	// TODO: Implement assignment
	a.Value.Evaluate()
}

type ReturnStatement struct {
	Expression Expression
}

func (n ReturnStatement) Evaluate() {
	// TODO: Implement
	n.Expression.Evaluate()
}
