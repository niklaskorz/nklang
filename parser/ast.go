package parser

type Node interface{}

type BlockNode struct {
	statements []Node
}

type DeclarationNode struct {
	Node
	Identifier string
	Value      Node
}

type AssignmentNode struct {
	Identifier string
	Value      Node
}

type StringNode struct {
	Value string
}

type IntegerNode struct {
	Value int64
}

type LookupNode struct {
	Identifier string
}
