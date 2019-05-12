package semantics

type definitionSet map[string]struct{}

func (s definitionSet) set(name string) {
	s[name] = struct{}{}
}

func (s definitionSet) has(name string) bool {
	_, ok := s[name]
	return ok
}

type DefinitionScope struct {
	parent      *DefinitionScope
	definitions definitionSet
}

func NewScope() *DefinitionScope {
	return &DefinitionScope{definitions: make(definitionSet)}
}

func (scope *DefinitionScope) newScope() *DefinitionScope {
	return &DefinitionScope{
		parent:      scope,
		definitions: make(definitionSet),
	}
}

func (scope *DefinitionScope) lookup(name string, index int) int {
	if scope.definitions.has(name) {
		return index
	}
	if scope.parent == nil {
		return -1
	}
	return scope.parent.lookup(name, index+1)
}

func (scope *DefinitionScope) declare(name string) {
	scope.definitions.set(name)
}

func (scope *DefinitionScope) Declare(name string) {
	scope.definitions.set(name)
}
