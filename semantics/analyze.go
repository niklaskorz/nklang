package semantics

import (
	"fmt"

	"niklaskorz.de/nklang/ast"
)

type definitionSet map[string]struct{}

func (s definitionSet) set(name string) {
	s[name] = struct{}{}
}

func (s definitionSet) has(name string) bool {
	_, ok := s[name]
	return ok
}

type definitionScope struct {
	parent      *definitionScope
	definitions definitionSet
}

func (scope *definitionScope) newScope() *definitionScope {
	return &definitionScope{
		parent:      scope,
		definitions: make(definitionSet),
	}
}

func (scope *definitionScope) lookup(name string, index int) int {
	if scope.definitions.has(name) {
		return index
	}
	if scope.parent == nil {
		return -1
	}
	return scope.parent.lookup(name, index+1)
}

func AnalyzeLookups(p *ast.Program) error {
	globalScope := &definitionScope{definitions: make(definitionSet)}

	for _, n := range p.Statements {
		if err := analyzeStatement(globalScope, n); err != nil {
			return err
		}
	}

	return nil
}

func analyzeStatement(scope *definitionScope, n ast.Statement) error {
	switch s := n.(type) {
	case *ast.IfStatement:
		if s.Condition != nil {
			if err := analyzeExpression(scope, s.Condition); err != nil {
				return err
			}
		}
		ds := scope.newScope()
		for _, n := range s.Statements {
			if err := analyzeStatement(ds, n); err != nil {
				return err
			}
		}
		if s.ElseBranch != nil {
			if err := analyzeStatement(scope, s.ElseBranch); err != nil {
				return err
			}
		}
	case *ast.WhileStatement:
		if err := analyzeExpression(scope, s.Condition); err != nil {
			return err
		}
		ds := scope.newScope()
		for _, n := range s.Statements {
			if err := analyzeStatement(ds, n); err != nil {
				return err
			}
		}
	case *ast.DeclarationStatement:
		if err := analyzeExpression(scope, s.Value); err != nil {
			return err
		}
		if scope.definitions.has(s.Identifier) {
			return fmt.Errorf("Redeclaration of %s in same scope", s.Identifier)
		}
		scope.definitions.set(s.Identifier)
	case *ast.AssignmentStatement:
		if err := analyzeExpression(scope, s.Value); err != nil {
			return err
		}
		scopeIndex := scope.lookup(s.Identifier, 0)
		if scopeIndex == -1 {
			return fmt.Errorf("%s must be declared before assignment", s.Identifier)
		}
		s.ScopeIndex = scopeIndex
	case *ast.ReturnStatement:
		if err := analyzeExpression(scope, s.Expression); err != nil {
			return err
		}
	case *ast.ExpressionStatement:
		if err := analyzeExpression(scope, s.Expression); err != nil {
			return err
		}
	}

	return nil
}

func analyzeExpression(scope *definitionScope, n ast.Expression) error {
	switch e := n.(type) {
	case *ast.IfExpression:
		if e.Condition != nil {
			if err := analyzeExpression(scope, e.Condition); err != nil {
				return err
			}
		}
		if err := analyzeExpression(scope, e.Value); err != nil {
			return err
		}
		if e.ElseBranch != nil {
			if err := analyzeExpression(scope, e.ElseBranch); err != nil {
				return err
			}
		}
	case *ast.LogicalOrExpression:
		if err := analyzeExpression(scope, e.A); err != nil {
			return err
		}
		if err := analyzeExpression(scope, e.B); err != nil {
			return err
		}
	case *ast.LogicalAndExpression:
		if err := analyzeExpression(scope, e.A); err != nil {
			return err
		}
		if err := analyzeExpression(scope, e.B); err != nil {
			return err
		}
	case *ast.ComparisonExpression:
		if err := analyzeExpression(scope, e.A); err != nil {
			return err
		}
		if err := analyzeExpression(scope, e.B); err != nil {
			return err
		}
	case *ast.AdditionExpression:
		if err := analyzeExpression(scope, e.A); err != nil {
			return err
		}
		if err := analyzeExpression(scope, e.B); err != nil {
			return err
		}
	case *ast.SubstractionExpression:
		if err := analyzeExpression(scope, e.A); err != nil {
			return err
		}
		if err := analyzeExpression(scope, e.B); err != nil {
			return err
		}
	case *ast.MultiplicationExpression:
		if err := analyzeExpression(scope, e.A); err != nil {
			return err
		}
		if err := analyzeExpression(scope, e.B); err != nil {
			return err
		}
	case *ast.DivisionExpression:
		if err := analyzeExpression(scope, e.A); err != nil {
			return err
		}
		if err := analyzeExpression(scope, e.B); err != nil {
			return err
		}
	case *ast.LookupExpression:
		scopeIndex := scope.lookup(e.Identifier, 0)
		if scopeIndex == -1 {
			return fmt.Errorf("%s must be declared before usage", e.Identifier)
		}
		e.ScopeIndex = scopeIndex
	case *ast.CallExpression:
		if err := analyzeExpression(scope, e.Callee); err != nil {
			return err
		}
		for _, p := range e.Parameters {
			if err := analyzeExpression(scope, p); err != nil {
				return err
			}
		}
	case *ast.Function:
		ds := scope.newScope()
		for _, p := range e.Parameters {
			ds.definitions.set(p)
		}
		for _, s := range e.Statements {
			if err := analyzeStatement(ds, s); err != nil {
				return err
			}
		}
	}

	return nil
}
