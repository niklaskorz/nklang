package evaluator

import "niklaskorz.de/nklang/ast"

func Evaluate(p *ast.Program) error {
	scope := &definitionScope{definitions: make(definitionMap)}
	return evaluateStatements(p.Statements, scope)
}
