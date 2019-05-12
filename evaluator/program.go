package evaluator

import "github.com/niklaskorz/nklang/ast"

func Evaluate(p *ast.Program) error {
	return EvaluateWithScope(p, NewScope())
}

func EvaluateWithScope(p *ast.Program, scope *DefinitionScope) error {
	return evaluateStatements(p.Statements, scope)
}
