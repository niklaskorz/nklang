package ast

type Expression interface {
	Evaluate() Object
}

type IfExpression struct {
	Condition  Expression
	Value      Expression
	ElseBranch *IfExpression
}

func (n IfExpression) Evaluate() Object {
	if n.Condition != nil {
		if n.Condition.Evaluate().IsTrue() {
			return n.Value.Evaluate()
		}
		// Else branch must be set if condition is set
		return n.ElseBranch.Evaluate()
	}
	return n.Value.Evaluate()
}

type LogicalOrExpression struct {
	A Expression
	B Expression
}

func (n LogicalOrExpression) Evaluate() Object {
	a := n.A.Evaluate()
	if a.IsTrue() {
		return a
	}
	return n.B.Evaluate()
}

type LogicalAndExpression struct {
	A Expression
	B Expression
}

func (n LogicalAndExpression) Evaluate() Object {
	a := n.A.Evaluate()
	if a.IsTrue() {
		return n.B.Evaluate()
	}
	return a
}

type LookupExpression struct {
	Identifier string
}

func (n LookupExpression) Evaluate() Object {
	// TODO: Implement lookup
	return Integer{Value: 0}
}
