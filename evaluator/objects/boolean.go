package objects

import "niklaskorz.de/nklang/ast"

type Boolean ast.Boolean

func (o *Boolean) IsTrue() bool {
	return o.Value
}

func (o *Boolean) Equals(other Object) (Object, error) {
	switch other := other.(type) {
	case *Boolean:
		return &Boolean{Value: o.Value == other.Value}, nil
	default:
		return nil, operationNotSupported
	}
}

func (o *Boolean) Lt(other Object) (Object, error) {
	return nil, operationNotSupported
}

func (o *Boolean) Lte(other Object) (Object, error) {
	return nil, operationNotSupported
}

func (o *Boolean) Gt(other Object) (Object, error) {
	return nil, operationNotSupported
}

func (o *Boolean) Gte(other Object) (Object, error) {
	return nil, operationNotSupported
}

func (o *Boolean) Add(other Object) (Object, error) {
	return nil, operationNotSupported
}

func (o *Boolean) Sub(other Object) (Object, error) {
	return nil, operationNotSupported
}

func (o *Boolean) Mul(other Object) (Object, error) {
	return nil, operationNotSupported
}

func (o *Boolean) Div(other Object) (Object, error) {
	return nil, operationNotSupported
}
