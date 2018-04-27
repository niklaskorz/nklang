package ast

type Object interface {
	Evaluate() Object
	IsTrue() bool
}
type Function struct {
	Parameters []string
	Statements []Statement
}

func (n Function) Evaluate() Object {
	return n
}

func (n Function) IsTrue() bool {
	return true
}

type Integer struct {
	Value int64
}

func (n Integer) Evaluate() Object {
	return n
}

func (n Integer) IsTrue() bool {
	return n.Value != 0
}

type String struct {
	Value string
}

func (n String) Evaluate() Object {
	return n
}

func (n String) IsTrue() bool {
	return n.Value != ""
}

type Boolean struct {
	Value bool
}

func (n Boolean) Evaluate() Object {
	return n
}

func (n Boolean) IsTrue() bool {
	return n.Value
}
