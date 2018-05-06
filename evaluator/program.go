package evaluator

import "niklaskorz.de/nklang/ast"

func Evaluate(p *ast.Program) error {
	scope := NewScope()
	return evaluateStatements(p.Statements, scope)
}

func Evaluate(p *ast.Program, scope *definitionScope) error {
	return evaluateStatements(p.Statements, scope)
}
