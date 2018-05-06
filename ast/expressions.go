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
	return fmt.Sprintf("CallExpresion{Callee: %s, Parameters: %d}", n.Callee, n.Parameters)
}
