package evaluator

import (
	"github.com/niklaskorz/nklang/ast"
)

func evaluateExpression(n ast.Expression, scope *DefinitionScope) (Object, error) {
	switch e := n.(type) {
	case *ast.Function:
		return &Function{Function: e, parentScope: scope}, nil
	case *ast.Integer:
		return (*Integer)(e), nil
	case *ast.Float:
		return (*Float)(e), nil
	case *ast.String:
		return (*String)(e), nil
	case *ast.Boolean:
		return (*Boolean)(e), nil
	case *ast.ArrayExpression:
		return evaluateArrayExpression(e, scope)
	case *ast.IfExpression:
		return evaluateIfExpression(e, scope)
	case *ast.BinaryOperationExpression:
		return evaluateBinaryExpression(e, scope)
	case *ast.UnaryOperationExpression:
		return evaluateUnaryExpression(e, scope)
	case *ast.LookupExpression:
		return evaluateLookupExpression(e, scope)
	case *ast.CallExpression:
		return evaluateCallExpression(e, scope)
	case *ast.SubscriptExpression:
		return evaluateSubscriptExpression(e, scope)
	}

	return nil, nil
}

func evaluateArrayExpression(n *ast.ArrayExpression, scope *DefinitionScope) (Object, error) {
	items := []Object{}
	for _, e := range n.Items {
		v, err := evaluateExpression(e, scope)
		if err != nil {
			return nil, err
		}
		items = append(items, v)
	}
	return &Array{Items: items}, nil
}

func evaluateIfExpression(n *ast.IfExpression, scope *DefinitionScope) (Object, error) {
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

func evaluateBinaryExpression(n *ast.BinaryOperationExpression, scope *DefinitionScope) (Object, error) {
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
	case ast.BinaryOperatorNe:
		o, err := aValue.Equals(bValue)
		if err != nil {
			return nil, err
		}
		return &Boolean{Value: !o.Value}, nil
	case ast.BinaryOperatorLt:
		if v, ok := aValue.(Comparable); ok {
			return v.Lt(bValue)
		}
	case ast.BinaryOperatorLe:
		if v, ok := aValue.(Comparable); ok {
			return v.Lte(bValue)
		}
	case ast.BinaryOperatorGt:
		if v, ok := aValue.(Comparable); ok {
			return v.Gt(bValue)
		}
	case ast.BinaryOperatorGe:
		if v, ok := aValue.(Comparable); ok {
			return v.Gte(bValue)
		}
	case ast.BinaryOperatorAdd:
		if v, ok := aValue.(Addable); ok {
			return v.Add(bValue)
		}
	case ast.BinaryOperatorSub:
		if v, ok := aValue.(Subtractable); ok {
			return v.Sub(bValue)
		}
	case ast.BinaryOperatorMul:
		if v, ok := aValue.(Multipliable); ok {
			return v.Mul(bValue)
		}
	case ast.BinaryOperatorDiv:
		if v, ok := aValue.(Dividable); ok {
			return v.Div(bValue)
		}
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

	return nil, operationNotSupported
}

func evaluateUnaryExpression(n *ast.UnaryOperationExpression, scope *DefinitionScope) (Object, error) {
	value, err := evaluateExpression(n.A, scope)
	if err != nil {
		return nil, err
	}

	switch n.Operator {
	case ast.UnaryOperatorLnot:
		return &Boolean{Value: !value.IsTrue()}, nil
	case ast.UnaryOperatorPos:
		if value, ok := value.(ObjectWithPos); ok {
			return value.Pos()
		}
	case ast.UnaryOperatorNeg:
		if value, ok := value.(ObjectWithNeg); ok {
			return value.Neg()
		}
	}

	return nil, operationNotSupported
}

func evaluateLookupExpression(n *ast.LookupExpression, scope *DefinitionScope) (Object, error) {
	return scope.lookup(n.Identifier, n.ScopeIndex), nil
}

func evaluateCallExpression(n *ast.CallExpression, scope *DefinitionScope) (Object, error) {
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

func evaluateFunctionCall(o *Function, params []ast.Expression, scope *DefinitionScope) (Object, error) {
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
		case *continueError:
			return nil, err.syntaxError()
		case *breakError:
			return nil, err.syntaxError()
		default:
			return nil, err
		}
	}

	return NilObject, nil
}

func evaluatePredefinedFunctionCall(o *PredefinedFunction, params []ast.Expression, scope *DefinitionScope) (Object, error) {
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

func evaluateSubscriptExpression(n *ast.SubscriptExpression, scope *DefinitionScope) (Object, error) {
	target, err := evaluateExpression(n.Target, scope)
	if err != nil {
		return nil, err
	}

	o, ok := target.(Subscriptable)
	if !ok {
		return nil, operationNotSupported
	}

	index, err := evaluateExpression(n.Index, scope)
	if err != nil {
		return nil, err
	}

	return o.Subscript(index)
}

func EvaluateExpression(n ast.Expression, scope *DefinitionScope) (Object, error) {
	return evaluateExpression(n, scope)
}
