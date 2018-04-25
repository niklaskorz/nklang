package ast

type Lookup struct {
	Identifier string
}

func (l Lookup) String() string {
	return l.Identifier
}

func (l Lookup) Evaluate() Node {
	// TODO: Implement lookup
	return Integer{Value: 0}
}

func (l Lookup) IsTrue() bool {
	return l.Evaluate().IsTrue()
}
