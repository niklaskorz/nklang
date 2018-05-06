package semantics

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

func (scope *definitionScope) declare(name string) {
	scope.definitions.set(name)
}
