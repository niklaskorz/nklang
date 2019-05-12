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

type Integer struct {
	Value int64
}

func (n *Integer) String() string {
	return fmt.Sprintf("Integer{Value: %d}", n.Value)
}

type Float struct {
	Value float64
}

func (n *Float) String() string {
	return fmt.Sprintf("Float{Value: %f}", n.Value)
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
