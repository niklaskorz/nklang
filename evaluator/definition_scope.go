package evaluator

type definitionMap map[string]Object

type DefinitionScope struct {
	parent      *DefinitionScope
	definitions definitionMap
}

func NewScope() *DefinitionScope {
	return &DefinitionScope{
		definitions: make(definitionMap),
	}
}

func (scope *DefinitionScope) newScope() *DefinitionScope {
	return &DefinitionScope{
		parent:      scope,
		definitions: make(definitionMap),
	}
}

func (scope *DefinitionScope) lookup(name string, index int) Object {
	if index == 0 {
		return scope.definitions[name]
	}
	return scope.parent.lookup(name, index-1)
}

func (scope *DefinitionScope) declare(name string, value Object) {
	scope.definitions[name] = value
}

func (scope *DefinitionScope) Declare(name string, value Object) {
	scope.definitions[name] = value
}

func (scope *DefinitionScope) assign(name string, value Object, index int) {
	if index == 0 {
		scope.definitions[name] = value
		return
	}
	scope.parent.assign(name, value, index-1)
}
