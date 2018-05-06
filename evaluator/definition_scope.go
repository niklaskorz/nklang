package evaluator

import "niklaskorz.de/nklang/evaluator/objects"

type definitionMap map[string]objects.Object

type definitionScope struct {
	parent      *definitionScope
	definitions definitionMap
}

func (scope *definitionScope) newScope() *definitionScope {
	return &definitionScope{
		parent:      scope,
		definitions: make(definitionMap),
	}
}

func (scope *definitionScope) lookup(name string, index int) objects.Object {
	if index == 0 {
		return scope.definitions[name]
	}
	return scope.parent.lookup(name, index-1)
}

func (scope *definitionScope) declare(name string, value objects.Object) {
	scope.definitions[name] = value
}

func (scope *definitionScope) assign(name string, value objects.Object, index int) {
	if index == 0 {
		scope.definitions[name] = value
	}
	scope.assign(name, value, index-1)
}
