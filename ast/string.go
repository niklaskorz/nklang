package ast

type String struct {
	Value string
}

func (s String) String() string {
	return "\"" + s.Value + "\""
}

func (s String) Evaluate() Node {
	return s
}

func (s String) IsTrue() bool {
	return s.Value != ""
}
