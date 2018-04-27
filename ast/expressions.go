package ast

type Expression interface {
	Evaluate() Expression
	IsTrue() bool
}

type IfExpression struct {
	Condition  Expression
	Value      Expression
	ElseBranch *IfExpression
}

func (n IfExpression) Evaluate() Expression {
	if n.Condition != nil {
		if n.Condition.IsTrue() {
			return n.Value.Evaluate()
		}
		// Else branch must be set if condition is set
		return n.ElseBranch.Evaluate()
	}
	return n.Value.Evaluate()
}

type LookupExpression struct {
	Identifier string
}

func (n LookupExpression) Evaluate() Expression {
	// TODO: Implement lookup
	return IntegerExpression{Value: 0}
}

func (n LookupExpression) IsTrue() bool {
	return n.Evaluate().IsTrue()
}

type IntegerExpression struct {
	Value int64
}

func (n IntegerExpression) Evaluate() Expression {
	return n
}

func (n IntegerExpression) IsTrue() bool {
	return n.Value != 0
}

type StringExpression struct {
	Value string
}

func (n StringExpression) Evaluate() Expression {
	return n
}

func (n StringExpression) IsTrue() bool {
	return n.Value != ""
}

type BooleanExpression struct {
	Value bool
}

func (n BooleanExpression) Evaluate() Expression {
	return n
}

func (n BooleanExpression) IsTrue() bool {
	return n.Value
}
