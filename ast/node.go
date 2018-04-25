package ast

type Node interface {
	String() string
	Evaluate() Node
	IsTrue() bool
}
