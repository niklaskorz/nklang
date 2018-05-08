package evaluator

type definitionMap map[string]Object

type definitionScope struct {
	parent      *definitionScope
	definitions definitionMap
}

func NewScope() *definitionScope {
	return &definitionScope{
		definitions: make(definitionMap),
	}
}

func (scope *definitionScope) newScope() *definitionScope {
	return &definitionScope{
		parent:      scope,
		definitions: make(definitionMap),
	}
}

func (scope *definitionScope) lookup(name string, index int) Object {
	if index == 0 {
		return scope.definitions[name]
	}
	return scope.parent.lookup(name, index-1)
}

func (scope *definitionScope) declare(name string, value Object) {
	scope.definitions[name] = value
}

func (scope *definitionScope) Declare(name string, value Object) {
	scope.definitions[name] = value
}

func (scope *definitionScope) assign(name string, value Object, index int) {
	if index == 0 {
		scope.definitions[name] = value
		return
	}
	scope.parent.assign(name, value, index-1)
}
