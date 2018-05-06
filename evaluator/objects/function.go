package objects

import "niklaskorz.de/nklang/ast"

type Function ast.Function

func (o *Function) IsTrue() bool {
	return true
}

func (o *Function) Equals(other Object) (Object, error) {
	return &Boolean{Value: o == other}, nil
}

func (o *Function) Lt(other Object) (Object, error) {
	return nil, operationNotSupported
}

func (o *Function) Lte(other Object) (Object, error) {
	return nil, operationNotSupported
}

func (o *Function) Gt(other Object) (Object, error) {
	return nil, operationNotSupported
}

func (o *Function) Gte(other Object) (Object, error) {
	return nil, operationNotSupported
}

func (o *Function) Add(other Object) (Object, error) {
	return nil, operationNotSupported
}

func (o *Function) Sub(other Object) (Object, error) {
	return nil, operationNotSupported
}

func (o *Function) Mul(other Object) (Object, error) {
	return nil, operationNotSupported
}

func (o *Function) Div(other Object) (Object, error) {
	return nil, operationNotSupported
}
