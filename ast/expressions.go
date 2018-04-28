package ast

import (
	"fmt"
)

type Expression interface {
	Evaluate() (Object, error)
}

type IfExpression struct {
	Condition  Expression
	Value      Expression
	ElseBranch *IfExpression
}

func (n *IfExpression) String() string {
	return fmt.Sprintf("IfExpression{Condition: %s, Value: %s, ElseBranch: %s}", n.Condition, n.Value, n.ElseBranch)
}

func (n *IfExpression) Evaluate() (Object, error) {
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

type BinaryOperator int

const (
	BinaryOperatorEq BinaryOperator = iota
	BinaryOperatorLt
	BinaryOperatorLe
	BinaryOperatorGt
	BinaryOperatorGe
	BinaryOperatorLand
	BinaryOperatorLor
	BinaryOperatorAdd
	BinaryOperatorSub
	BinaryOperatorMul
	BinaryOperatorDiv
)

type BinaryOperationExpression struct {
	Operator BinaryOperator
	A        Expression
	B        Expression
}

func (n *BinaryOperationExpression) String() string {
	return fmt.Sprintf("BinaryOperationExpression{Operator: %d, A: %s, B: %s}", n.Operator, n.A, n.B)
}

func (n *BinaryOperationExpression) Evaluate() (Object, error) {
	// TODO: Implement
	return &Integer{Value: 0}, nil
}

type LookupExpression struct {
	Identifier string
	ScopeIndex int
}

func (n *LookupExpression) String() string {
	return fmt.Sprintf("LookupExpression{Identifier: %s, ScopeIndex: %d}", n.Identifier, n.ScopeIndex)
}

func (n *LookupExpression) Evaluate() (Object, error) {
	// TODO: Implement lookup
	return &Integer{Value: 0}, nil
}

type CallExpression struct {
	Callee     Expression
	Parameters []Expression
}

func (n *CallExpression) String() string {
	return fmt.Sprintf("CallExpresion{Callee: %s, Parameters: %d}", n.Callee, n.Parameters)
}

func (n *CallExpression) Evaluate() (Object, error) {
	// TODO: Implement
	return &Integer{Value: 0}, nil
}
