package ast

type Expression interface {
	Evaluate() (Object, error)
}

type IfExpression struct {
	Condition  Expression
	Value      Expression
	ElseBranch *IfExpression
}

func (n IfExpression) Evaluate() (Object, error) {
	if n.Condition == nil {
		return n.Value.Evaluate()
	}

	c, err := n.Condition.Evaluate()
	if err != nil {
		return nil, err
	}
	if c.IsTrue() {
		return n.Value.Evaluate()
	}
	// Else branch must be set if condition is set
	return n.ElseBranch.Evaluate()
}

type LogicalOrExpression struct {
	A Expression
	B Expression
}

func (n LogicalOrExpression) Evaluate() (Object, error) {
	a, err := n.A.Evaluate()
	if err != nil {
		return nil, err
	}
	if a.IsTrue() {
		return a, nil
	}
	return n.B.Evaluate()
}

type LogicalAndExpression struct {
	A Expression
	B Expression
}

func (n LogicalAndExpression) Evaluate() (Object, error) {
	a, err := n.A.Evaluate()
	if err != nil {
		return nil, err
	}
	if a.IsTrue() {
		return n.B.Evaluate()
	}
	return a, nil
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

func (n ComparisonExpression) Evaluate() (Object, error) {
	// TODO: Implement
	return Boolean{Value: true}, nil
}

type AdditionExpression struct {
	A Expression
	B Expression
}

func (n AdditionExpression) Evaluate() (Object, error) {
	// TODO: Implement
	return Integer{Value: 0}, nil
}

type SubstractionExpression struct {
	A Expression
	B Expression
}

func (n SubstractionExpression) Evaluate() (Object, error) {
	// TODO: Implement
	return Integer{Value: 0}, nil
}

type MultiplicationExpression struct {
	A Expression
	B Expression
}

func (n MultiplicationExpression) Evaluate() (Object, error) {
	// TODO: Implement
	return Integer{Value: 0}, nil
}

type DivisionExpression struct {
	A Expression
	B Expression
}

func (n DivisionExpression) Evaluate() (Object, error) {
	// TODO: Implement
	return Integer{Value: 0}, nil
}

type LookupExpression struct {
	Identifier string
}

func (n LookupExpression) Evaluate() (Object, error) {
	// TODO: Implement lookup
	return Integer{Value: 0}, nil
}

type CallExpression struct {
	Callee     Expression
	Parameters []Expression
}

func (n CallExpression) Evaluate() (Object, error) {
	// TODO: Implement
	return Integer{Value: 0}, nil
}
