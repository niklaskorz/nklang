package ast

import "fmt"

type ValueExpression interface {
	Expression
}

type Function struct {
	Parameters []string
	Statements []Statement
}

func (n *Function) String() string {
	return fmt.Sprintf("Function{Parameters: %s, Statements: %s}", n.Parameters, n.Statements)
}

func (n *Function) IsTrue() bool {
	return true
}

type Integer struct {
	Value int64
}

func (n *Integer) String() string {
	return fmt.Sprintf("Integer{Value: %d}", n.Value)
}

func (n *Integer) IsTrue() bool {
	return n.Value != 0
}

type String struct {
	Value string
}

func (n *String) String() string {
	return fmt.Sprintf("String{Value: %s}", n.Value)
}

type Boolean struct {
	Value bool
}

func (n *Boolean) String() string {
	return fmt.Sprintf("Boolean{Value: %t}", n.Value)
}

type Nil struct{}

func (n *Nil) String() string {
	return "Nil"
}
