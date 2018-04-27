package ast

type Statement interface {
	Evaluate()
}

type ExpressionStatement struct {
	Expression Expression
}

func (n ExpressionStatement) Evaluate() {
	n.Expression.Evaluate()
}

type Declaration struct {
	Identifier string
	Value      Expression
}

func (n Declaration) Evaluate() {
	// TODO: Implement declaration
	n.Value.Evaluate()
}

type Assignment struct {
	Identifier string
	Value      Expression
}

func (n Assignment) Evaluate() {
	// TODO: Implement assignment
	n.Value.Evaluate()
}

type ReturnStatement struct {
	Expression Expression
}

func (n ReturnStatement) Evaluate() {
	// TODO: Implement
	n.Expression.Evaluate()
}
