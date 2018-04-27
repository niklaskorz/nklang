package ast

type Expression interface {
	Evaluate() Expression
	IsTrue() bool
}

type Lookup struct {
	Identifier string
}

func (n Lookup) Evaluate() Expression {
	// TODO: Implement lookup
	return Integer{Value: 0}
}

func (n Lookup) IsTrue() bool {
	return n.Evaluate().IsTrue()
}

type Integer struct {
	Value int64
}

func (n Integer) Evaluate() Expression {
	return n
}

func (n Integer) IsTrue() bool {
	return n.Value != 0
}

type String struct {
	Value string
}

func (n String) Evaluate() Expression {
	return n
}

func (n String) IsTrue() bool {
	return n.Value != ""
}
