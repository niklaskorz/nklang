package evaluator

import (
	"fmt"

	"niklaskorz.de/nklang/ast"
	"niklaskorz.de/nklang/evaluator/objects"
)

func evaluateExpression(n ast.Expression, scope *definitionScope) (objects.Object, error) {
	switch e := n.(type) {
	case *ast.Function:
		return (*objects.Function)(e), nil
	case *ast.Integer:
		return (*objects.Integer)(e), nil
	case *ast.String:
		return (*objects.String)(e), nil
	case *ast.Boolean:
		return (*objects.Boolean)(e), nil
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

func evaluateIfExpression(n *ast.IfExpression, scope *definitionScope) (objects.Object, error) {
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

func evaluateBinaryExpression(n *ast.BinaryOperationExpression, scope *definitionScope) (objects.Object, error) {
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

func evaluateLookupExpression(n *ast.LookupExpression, scope *definitionScope) (objects.Object, error) {
	return scope.lookup(n.Identifier, n.ScopeIndex), nil
}

func evaluateCallExpression(n *ast.CallExpression, scope *definitionScope) (objects.Object, error) {
	callee, err := evaluateExpression(n.Callee, scope)
	if err != nil {
		return nil, err
	}

	switch callee := callee.(type) {
	case *objects.Function:
		return evaluateFunctionCall(callee, n.Parameters, scope)
	case *objects.PredefinedFunction:
		return evaluatePredefinedFunction(callee, n.Parameters, scope)
	}

	return nil, objects.OperationNotSupportedError{}
}

func evaluateFunctionCall(o *objects.Function, params []ast.Expression, scope *definitionScope) (objects.Object, error) {
	parameterScope := scope.newScope()
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

	return objects.NilObject, nil
}

func evaluatePredefinedFunction(o *objects.PredefinedFunction, params []ast.Expression, scope *definitionScope) (objects.Object, error) {
	parameters := []objects.Object{}
	for _, p := range params {
		v, err := evaluateExpression(p, scope)
		if err != nil {
			return nil, err
		}
		parameters = append(parameters, v)
	}

	return (*o)(parameters)
}
