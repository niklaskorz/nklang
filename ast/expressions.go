package ast

import (
	"fmt"
)

type Expression interface{}

type IfExpression struct {
	Condition  Expression
	Value      Expression
	ElseBranch *IfExpression
}

func (n *IfExpression) String() string {
	return fmt.Sprintf("IfExpression{Condition: %s, Value: %s, ElseBranch: %s}", n.Condition, n.Value, n.ElseBranch)
}

type BinaryOperator int

const (
	BinaryOperatorEq BinaryOperator = iota
	BinaryOperatorNe
	BinaryOperatorLt
	BinaryOperatorLe
	BinaryOperatorGt
	BinaryOperatorGe
	BinaryOperatorAdd
	BinaryOperatorSub
	BinaryOperatorMul
	BinaryOperatorDiv
	BinaryOperatorLand
	BinaryOperatorLor
)

type BinaryOperationExpression struct {
	Operator BinaryOperator
	A        Expression
	B        Expression
}

func (n *BinaryOperationExpression) String() string {
	return fmt.Sprintf("BinaryOperationExpression{Operator: %d, A: %s, B: %s}", n.Operator, n.A, n.B)
}

type UnaryOperator int

const (
	UnaryOperatorLnot UnaryOperator = iota
	UnaryOperatorPos
	UnaryOperatorNeg
)

type UnaryOperationExpression struct {
	Operator UnaryOperator
	A        Expression
}

func (n *UnaryOperationExpression) String() string {
	return fmt.Sprintf("UnaryOperationExpression{Operator: %d, A: %s}", n.Operator, n.A)
}

type LookupExpression struct {
	Identifier string
	ScopeIndex int
}

func (n *LookupExpression) String() string {
	return fmt.Sprintf("LookupExpression{Identifier: %s, ScopeIndex: %d}", n.Identifier, n.ScopeIndex)
}

type CallExpression struct {
	Callee     Expression
	Parameters []Expression
}

func (n *CallExpression) String() string {
	return fmt.Sprintf("CallExpresion{Callee: %s, Parameters: %v}", n.Callee, n.Parameters)
}

type SubscriptExpression struct {
	Target Expression
	Index  Expression
}

func (n *SubscriptExpression) String() string {
	return fmt.Sprintf("SubscriptExpression{Target: %s, Index: %s}", n.Target, n.Index)
}
