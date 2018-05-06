package ast

import (
	"fmt"
)

type Statement interface{}

type IfStatement struct {
	Condition  Expression
	Statements []Statement
	ElseBranch *IfStatement
}

func (n *IfStatement) String() string {
	return fmt.Sprintf("IfStatement{Condition: %s, Statements: %s, ElseBranch: %s}", n.Condition, n.Statements, n.ElseBranch)
}

type WhileStatement struct {
	Condition  Expression
	Statements []Statement
}

func (n *WhileStatement) String() string {
	return fmt.Sprintf("WhileStatement{Condition: %s, Statements: %s}", n.Condition, n.Statements)
}

type ExpressionStatement struct {
	Expression Expression
}

func (n *ExpressionStatement) String() string {
	return fmt.Sprintf("ExpressionStatement{Expression: %s}", n.Expression)
}

type DeclarationStatement struct {
	Identifier string
	Value      Expression
}

func (n *DeclarationStatement) String() string {
	return fmt.Sprintf("DeclarationStatement{Identifier: %s, Value: %s}", n.Identifier, n.Value)
}

type AssignmentStatement struct {
	Identifier string
	ScopeIndex int
	Value      Expression
}

func (n *AssignmentStatement) String() string {
	return fmt.Sprintf("AssignmentStatement{Identifier: %s, ScopeIndex: %d, Value: %s}", n.Identifier, n.ScopeIndex, n.Value)
}

type ReturnStatement struct {
	Expression Expression
}

func (n *ReturnStatement) String() string {
	return fmt.Sprintf("ReturnStatement{Expression: %s}", n.Expression)
}

type ContinueStatement struct{}

func (n *ContinueStatement) String() string {
	return "ContinueStatement"
}

type BreakStatement struct{}

func (n *BreakStatement) String() string {
	return "BreakStatement"
}
