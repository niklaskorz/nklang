package objects

type PredefinedFunction func(params []Object) (Object, error)

func (o PredefinedFunction) IsTrue() bool {
	return true
}

func (o PredefinedFunction) Equals(other Object) (Object, error) {
	return &Boolean{Value: false}, nil
}

func (o PredefinedFunction) Lt(other Object) (Object, error) {
	return nil, operationNotSupported
}

func (o PredefinedFunction) Lte(other Object) (Object, error) {
	return nil, operationNotSupported
}

func (o PredefinedFunction) Gt(other Object) (Object, error) {
	return nil, operationNotSupported
}

func (o PredefinedFunction) Gte(other Object) (Object, error) {
	return nil, operationNotSupported
}

func (o PredefinedFunction) Add(other Object) (Object, error) {
	return nil, operationNotSupported
}

func (o PredefinedFunction) Sub(other Object) (Object, error) {
	return nil, operationNotSupported
}

func (o PredefinedFunction) Mul(other Object) (Object, error) {
	return nil, operationNotSupported
}

func (o PredefinedFunction) Div(other Object) (Object, error) {
	return nil, operationNotSupported
}
