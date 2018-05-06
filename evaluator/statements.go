package evaluator

import (
	"niklaskorz.de/nklang/ast"
	"niklaskorz.de/nklang/evaluator/objects"
)

func evaluateStatements(statements []ast.Statement, scope *definitionScope) error {
	for _, s := range statements {
		if err := evaluateStatement(s, scope); err != nil {
			return err
		}
	}
	return nil
}

func evaluateStatement(n ast.Statement, scope *definitionScope) error {
	switch s := n.(type) {
	case *ast.IfStatement:
		return evaluateIfStatement(s, scope)
	case *ast.WhileStatement:
		return evaluateWhileStatement(s, scope)
	case *ast.ExpressionStatement:
		return evaluateExpressionStatement(s, scope)
	case *ast.DeclarationStatement:
		return evaluateDeclarationStatement(s, scope)
	case *ast.AssignmentStatement:
		return evaluateAssignmentStatement(s, scope)
	case *ast.ReturnStatement:
		return evaluateReturnStatement(s, scope)
	case *ast.ContinueStatement:
		return evaluateContinueStatement(s, scope)
	case *ast.BreakStatement:
		return evaluateBreakStatement(s, scope)
	}
	return nil
}

func evaluateIfStatement(n *ast.IfStatement, scope *definitionScope) error {
	if n.Condition == nil {
		return evaluateStatements(n.Statements, scope.newScope())
	}

	c, err := evaluateExpression(n.Condition, scope)
	if err != nil {
		return err
	}
	if c.IsTrue() {
		return evaluateStatements(n.Statements, scope.newScope())
	}
	if n.ElseBranch != nil {
		return evaluateIfStatement(n.ElseBranch, scope)
	}
	return nil
}

func evaluateWhileStatement(n *ast.WhileStatement, scope *definitionScope) error {
	for {
		c, err := evaluateExpression(n.Condition, scope)
		if err != nil {
			return err
		}
		if !c.IsTrue() {
			return nil
		}

		if err := evaluateStatements(n.Statements, scope.newScope()); err != nil {
			return err
		}
	}
}

func evaluateExpressionStatement(n *ast.ExpressionStatement, scope *definitionScope) error {
	_, err := evaluateExpression(n.Expression, scope)
	return err
}

func evaluateDeclarationStatement(n *ast.DeclarationStatement, scope *definitionScope) error {
	value, err := evaluateExpression(n.Value, scope)
	if err != nil {
		return err
	}
	scope.declare(n.Identifier, value)
	return nil
}

func evaluateAssignmentStatement(n *ast.AssignmentStatement, scope *definitionScope) error {
	value, err := evaluateExpression(n.Value, scope)
	if err != nil {
		return err
	}
	scope.assign(n.Identifier, value, n.ScopeIndex)
	return nil
}

type returnError struct {
	value objects.Object
}

func (r *returnError) Error() string {
	return "Unexpected return statement"
}

func evaluateReturnStatement(n *ast.ReturnStatement, scope *definitionScope) error {
	value, err := evaluateExpression(n.Expression, scope)
	if err != nil {
		return err
	}
	return &returnError{value: value}
}

func evaluateContinueStatement(n *ast.ContinueStatement, scope *definitionScope) error {
	return nil
}

func evaluateBreakStatement(n *ast.BreakStatement, scope *definitionScope) error {
	return nil
}
