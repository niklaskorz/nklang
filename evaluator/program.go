package evaluator

import "niklaskorz.de/nklang/ast"

func Evaluate(p *ast.Program) error {
	return EvaluateWithScope(p, NewScope())
}

func EvaluateWithScope(p *ast.Program, scope *definitionScope) error {
	return evaluateStatements(p.Statements, scope)
}
