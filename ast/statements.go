package ast

type Statement interface {
	Evaluate() error
}

type IfStatement struct {
	Condition  Expression
	Statements []Statement
	ElseBranch *IfStatement
}

func (n IfStatement) evaluateStatements() error {
	for _, s := range n.Statements {
		if err := s.Evaluate(); err != nil {
			return err
		}
	}
	return nil
}

func (n IfStatement) Evaluate() error {
	if n.Condition == nil {
		return n.evaluateStatements()
	}

	c, err := n.Condition.Evaluate()
	if err != nil {
		return err
	}
	if c.IsTrue() {
		return n.evaluateStatements()
	}
	if n.ElseBranch != nil {
		return n.ElseBranch.Evaluate()
	}
	return nil
}

type WhileStatement struct {
	Condition  Expression
	Statements []Statement
}

func (n WhileStatement) Evaluate() error {
	for {
		c, err := n.Condition.Evaluate()
		if err != nil {
			return err
		}
		if !c.IsTrue() {
			return nil
		}

		for _, s := range n.Statements {
			if err := s.Evaluate(); err != nil {
				return err
			}
		}
	}
}

type ExpressionStatement struct {
	Expression Expression
}

func (n ExpressionStatement) Evaluate() error {
	_, err := n.Expression.Evaluate()
	return err
}

type DeclarationStatement struct {
	Identifier string
	Value      Expression
}

func (n DeclarationStatement) Evaluate() error {
	// TODO: Implement declaration
	n.Value.Evaluate()
	return nil
}

type AssignmentStatement struct {
	Identifier string
	Value      Expression
}

func (n AssignmentStatement) Evaluate() error {
	// TODO: Implement assignment
	n.Value.Evaluate()
	return nil
}

type ReturnStatement struct {
	Expression Expression
}

func (n ReturnStatement) Evaluate() error {
	// TODO: Implement
	n.Expression.Evaluate()
	return nil
}

type ContinueStatement struct{}

func (n ContinueStatement) Evaluate() error {
	// TODO: Implement
	return nil
}

type BreakStatement struct{}

func (n BreakStatement) Evaluate() error {
	// TODO: Implement
	return nil
}
