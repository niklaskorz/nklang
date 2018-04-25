package ast

type Assignment struct {
	Identifier string
	Value      Node
}

func (a Assignment) String() string {
	return a.Identifier + " = " + a.Value.String()
}

func (a Assignment) Evaluate() Node {
	// TODO: Implement assignment
	return a.Value.Evaluate()
}

func (a Assignment) IsTrue() bool {
	return a.Evaluate().IsTrue()
}
