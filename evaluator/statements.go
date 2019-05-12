package evaluator

import (
	"github.com/niklaskorz/nklang/ast"
)

func evaluateStatements(statements []ast.Statement, scope *DefinitionScope) error {
	for _, s := range statements {
		if err := evaluateStatement(s, scope); err != nil {
			return err
		}
	}
	return nil
}

func evaluateStatement(n ast.Statement, scope *DefinitionScope) error {
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

func evaluateIfStatement(n *ast.IfStatement, scope *DefinitionScope) error {
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

func evaluateWhileStatement(n *ast.WhileStatement, scope *DefinitionScope) error {
	for {
		c, err := evaluateExpression(n.Condition, scope)
		if err != nil {
			return err
		}
		if !c.IsTrue() {
			return nil
		}

		if err := evaluateStatements(n.Statements, scope.newScope()); err != nil {
			switch err := err.(type) {
			case *continueError:
				continue
			case *breakError:
				return nil
			default:
				return err
			}
		}
	}
}

func evaluateExpressionStatement(n *ast.ExpressionStatement, scope *DefinitionScope) error {
	_, err := evaluateExpression(n.Expression, scope)
	return err
}

func evaluateDeclarationStatement(n *ast.DeclarationStatement, scope *DefinitionScope) error {
	value, err := evaluateExpression(n.Value, scope)
	if err != nil {
		return err
	}
	scope.declare(n.Identifier, value)
	return nil
}

func evaluateAssignmentStatement(n *ast.AssignmentStatement, scope *DefinitionScope) error {
	value, err := evaluateExpression(n.Value, scope)
	if err != nil {
		return err
	}
	scope.assign(n.Identifier, value, n.ScopeIndex)
	return nil
}

func evaluateReturnStatement(n *ast.ReturnStatement, scope *DefinitionScope) error {
	value, err := evaluateExpression(n.Expression, scope)
	if err != nil {
		return err
	}
	return &returnError{value: value}
}

func evaluateContinueStatement(n *ast.ContinueStatement, scope *DefinitionScope) error {
	return &continueError{}
}

func evaluateBreakStatement(n *ast.BreakStatement, scope *DefinitionScope) error {
	return &breakError{}
}
