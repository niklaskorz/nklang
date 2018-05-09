package evaluator

import (
	"fmt"

	"github.com/niklaskorz/nklang/ast"
)

func evaluateExpression(n ast.Expression, scope *definitionScope) (Object, error) {
	switch e := n.(type) {
	case *ast.Function:
		return &Function{Function: e, parentScope: scope}, nil
	case *ast.Integer:
		return (*Integer)(e), nil
	case *ast.String:
		return (*String)(e), nil
	case *ast.Boolean:
		return (*Boolean)(e), nil
	case *ast.IfExpression:
		return evaluateIfExpression(e, scope)
	case *ast.BinaryOperationExpression:
		return evaluateBinaryExpression(e, scope)
	case *ast.LookupExpression:
		return evaluateLookupExpression(e, scope)
	case *ast.CallExpression:
		return evaluateCallExpression(e, scope)
	}

	return nil, nil
}

func evaluateIfExpression(n *ast.IfExpression, scope *definitionScope) (Object, error) {
	if n.Condition == nil {
		return evaluateExpression(n.Value, scope)
	}

	c, err := evaluateExpression(n.Condition, scope)
	if err != nil {
		return nil, err
	}
	if c.IsTrue() {
		return evaluateExpression(n.Value, scope)
	}
	// Else branch must be set if condition is set
	return evaluateIfExpression(n.ElseBranch, scope)
}

func evaluateBinaryExpression(n *ast.BinaryOperationExpression, scope *definitionScope) (Object, error) {
	aValue, err := evaluateExpression(n.A, scope)
	if err != nil {
		return nil, err
	}

	bValue, err := evaluateExpression(n.B, scope)
	if err != nil {
		return nil, err
	}

	switch n.Operator {
	case ast.BinaryOperatorEq:
		return aValue.Equals(bValue)
	case ast.BinaryOperatorLt:
		return aValue.Lt(bValue)
	case ast.BinaryOperatorLe:
		return aValue.Lte(bValue)
	case ast.BinaryOperatorGt:
		return aValue.Gt(bValue)
	case ast.BinaryOperatorGe:
		return aValue.Gte(bValue)
	case ast.BinaryOperatorAdd:
		return aValue.Add(bValue)
	case ast.BinaryOperatorSub:
		return aValue.Sub(bValue)
	case ast.BinaryOperatorMul:
		return aValue.Mul(bValue)
	case ast.BinaryOperatorDiv:
		return aValue.Div(bValue)
	case ast.BinaryOperatorLand:
		if !aValue.IsTrue() {
			return aValue, nil
		}
		return bValue, nil
	case ast.BinaryOperatorLor:
		if aValue.IsTrue() {
			return aValue, nil
		}
		return bValue, nil
	}

	return nil, fmt.Errorf("Unknown binary expression")
}

func evaluateLookupExpression(n *ast.LookupExpression, scope *definitionScope) (Object, error) {
	return scope.lookup(n.Identifier, n.ScopeIndex), nil
}

func evaluateCallExpression(n *ast.CallExpression, scope *definitionScope) (Object, error) {
	callee, err := evaluateExpression(n.Callee, scope)
	if err != nil {
		return nil, err
	}

	switch callee := callee.(type) {
	case *Function:
		return evaluateFunctionCall(callee, n.Parameters, scope)
	case *PredefinedFunction:
		return evaluatePredefinedFunctionCall(callee, n.Parameters, scope)
	}

	return nil, OperationNotSupportedError{}
}

func evaluateFunctionCall(o *Function, params []ast.Expression, scope *definitionScope) (Object, error) {
	parameterScope := o.parentScope.newScope()
	for i, p := range params {
		v, err := evaluateExpression(p, scope)
		if err != nil {
			return nil, err
		}
		name := o.Parameters[i]
		parameterScope.declare(name, v)
	}

	if err := evaluateStatements(o.Statements, parameterScope.newScope()); err != nil {
		switch err := err.(type) {
		case *returnError:
			return err.value, nil
		default:
			return nil, err
		}
	}

	return NilObject, nil
}

func evaluatePredefinedFunctionCall(o *PredefinedFunction, params []ast.Expression, scope *definitionScope) (Object, error) {
	parameters := []Object{}
	for _, p := range params {
		v, err := evaluateExpression(p, scope)
		if err != nil {
			return nil, err
		}
		parameters = append(parameters, v)
	}

	return o.fn(parameters)
}
