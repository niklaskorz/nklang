package objects

import "niklaskorz.de/nklang/ast"

type Nil ast.Nil

func (o *Nil) IsTrue() bool {
	return false
}

func (o *Nil) Equals(other Object) (Object, error) {
	switch other.(type) {
	case *Nil:
		return &Boolean{Value: true}, nil
	default:
		return &Boolean{Value: false}, nil
	}
}

func (o *Nil) Lt(other Object) (Object, error) {
	return nil, operationNotSupported
}

func (o *Nil) Lte(other Object) (Object, error) {
	return nil, operationNotSupported
}

func (o *Nil) Gt(other Object) (Object, error) {
	return nil, operationNotSupported
}

func (o *Nil) Gte(other Object) (Object, error) {
	return nil, operationNotSupported
}

func (o *Nil) Add(other Object) (Object, error) {
	return nil, operationNotSupported
}

func (o *Nil) Sub(other Object) (Object, error) {
	return nil, operationNotSupported
}

func (o *Nil) Mul(other Object) (Object, error) {
	return nil, operationNotSupported
}

func (o *Nil) Div(other Object) (Object, error) {
	return nil, operationNotSupported
}
