package ast

type Void struct{}

func (v Void) String() string {
	return "void"
}

func (v Void) Evaluate() Node {
	return v
}

func (v Void) IsTrue() bool {
	return false
}
