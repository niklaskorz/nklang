package ast

type Statement interface {
	Evaluate()
}

type IfStatement struct {
	Condition  Expression
	Statements []Statement
	ElseBranch *IfStatement
}

func (n IfStatement) evaluateStatements() {
	for _, s := range n.Statements {
		s.Evaluate()
	}
}

func (n IfStatement) Evaluate() {
	if n.Condition != nil {
		if n.Condition.IsTrue() {
			n.evaluateStatements()
		} else if n.ElseBranch != nil {
			n.ElseBranch.Evaluate()
		}
	} else {
		n.evaluateStatements()
	}
}

type WhileStatement struct {
	Condition  Expression
	Statements []Statement
}

func (n WhileStatement) Evaluate() {
	for n.Condition.IsTrue() {
		for _, s := range n.Statements {
			s.Evaluate()
		}
	}
}

type ExpressionStatement struct {
	Expression Expression
}

func (n ExpressionStatement) Evaluate() {
	n.Expression.Evaluate()
}

type Declaration struct {
	Identifier string
	Value      Expression
}

func (n Declaration) Evaluate() {
	// TODO: Implement declaration
	n.Value.Evaluate()
}

type Assignment struct {
	Identifier string
	Value      Expression
}

func (n Assignment) Evaluate() {
	// TODO: Implement assignment
	n.Value.Evaluate()
}

type ReturnStatement struct {
	Expression Expression
}

func (n ReturnStatement) Evaluate() {
	// TODO: Implement
	n.Expression.Evaluate()
}

type ContinueStatement struct{}

func (n ContinueStatement) Evaluate() {
	// TODO: Implement
}

type BreakStatement struct{}

func (n BreakStatement) Evaluate() {
	// TODO: Implement
}
