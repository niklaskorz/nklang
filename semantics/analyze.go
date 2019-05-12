package semantics

import (
	"fmt"

	"github.com/niklaskorz/nklang/ast"
)

func AnalyzeLookups(p *ast.Program) error {
	globalScope := &DefinitionScope{definitions: make(definitionSet)}
	return AnalyzeLookupsWithScope(p, globalScope)
}

func AnalyzeLookupsWithScope(p *ast.Program, scope *DefinitionScope) error {
	for _, n := range p.Statements {
		if err := analyzeStatement(scope, n); err != nil {
			return err
		}
	}

	return nil
}

func analyzeStatement(scope *DefinitionScope, n ast.Statement) error {
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
		if scope.definitions.has(s.Identifier) {
			return fmt.Errorf("Redeclaration of %s in same scope", s.Identifier)
		}
		scope.declare(s.Identifier)
		if err := analyzeExpression(scope, s.Value); err != nil {
			return err
		}
	case *ast.AssignmentStatement:
		scopeIndex := scope.lookup(s.Identifier, 0)
		if scopeIndex == -1 {
			return fmt.Errorf("%s must be declared before assignment", s.Identifier)
		}
		s.ScopeIndex = scopeIndex
		if err := analyzeExpression(scope, s.Value); err != nil {
			return err
		}
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

func analyzeExpression(scope *DefinitionScope, n ast.Expression) error {
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
	case *ast.BinaryOperationExpression:
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
			ds.declare(p)
		}
		scope := ds.newScope()
		for _, s := range e.Statements {
			if err := analyzeStatement(scope, s); err != nil {
				return err
			}
		}
	case *ast.ArrayExpression:
		for _, item := range e.Items {
			if err := analyzeExpression(scope, item); err != nil {
				return err
			}
		}
	}

	return nil
}

func AnalyzeExpression(scope *DefinitionScope, n ast.Expression) error {
	return analyzeExpression(scope, n)
}
