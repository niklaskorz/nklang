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

type ComparisonOperator int

const (
	ComparisonOperatorEq ComparisonOperator = iota
	ComparisonOperatorLt
	ComparisonOperatorLe
	ComparisonOperatorGt
	ComparisonOperatorGe
)

type ComparisonExpression struct {
	Operator ComparisonOperator
	A        Expression
	B        Expression
}

func (n ComparisonExpression) Evaluate() Object {
	// TODO: Implement
	return Boolean{Value: true}
}

type AdditionExpression struct {
	A Expression
	B Expression
}

func (n AdditionExpression) Evaluate() Object {
	// TODO: Implement
	return Integer{Value: 0}
}

type SubstractionExpression struct {
	A Expression
	B Expression
}

func (n SubstractionExpression) Evaluate() Object {
	// TODO: Implement
	return Integer{Value: 0}
}

type MultiplicationExpression struct {
	A Expression
	B Expression
}

func (n MultiplicationExpression) Evaluate() Object {
	// TODO: Implement
	return Integer{Value: 0}
}

type DivisionExpression struct {
	A Expression
	B Expression
}

func (n DivisionExpression) Evaluate() Object {
	// TODO: Implement
	return Integer{Value: 0}
}

type LookupExpression struct {
	Identifier string
}

func (n LookupExpression) Evaluate() Object {
	// TODO: Implement lookup
	return Integer{Value: 0}
}

type CallExpression struct {
	Callee     Expression
	Parameters []Expression
}

func (n CallExpression) Evaluate() Object {
	// TODO: Implement
	return Integer{Value: 0}
}
