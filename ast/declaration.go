package ast

type Declaration struct {
	Identifier string
	Value      Node
}

func (d Declaration) String() string {
	return d.Identifier + " := " + d.Value.String()
}

func (d Declaration) Evaluate() Node {
	// TODO: Implement declaration
	return d.Value.Evaluate()
}

func (d Declaration) IsTrue() bool {
	return d.Evaluate().IsTrue()
}
