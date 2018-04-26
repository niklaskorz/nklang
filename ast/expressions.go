package ast

type Expression interface {
	Evaluate() Expression
	IsTrue() bool
}

type Lookup struct {
	Identifier string
}

func (l Lookup) Evaluate() Expression {
	// TODO: Implement lookup
	return Integer{Value: 0}
}

func (l Lookup) IsTrue() bool {
	return l.Evaluate().IsTrue()
}

type Integer struct {
	Value int64
}

func (i Integer) Evaluate() Expression {
	return i
}

func (i Integer) IsTrue() bool {
	return i.Value != 0
}

type String struct {
	Value string
}

func (s String) Evaluate() Expression {
	return s
}

func (s String) IsTrue() bool {
	return s.Value != ""
}
