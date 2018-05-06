package objects

import "niklaskorz.de/nklang/ast"

type String ast.String

func (o *String) IsTrue() bool {
	return o.Value != ""
}

func (o *String) Equals(other Object) (Object, error) {
	switch other := other.(type) {
	case *String:
		return &Boolean{Value: o.Value == other.Value}, nil
	default:
		return nil, operationNotSupported
	}
}

func (o *String) Lt(other Object) (Object, error) {
	return nil, operationNotSupported
}

func (o *String) Lte(other Object) (Object, error) {
	return nil, operationNotSupported
}

func (o *String) Gt(other Object) (Object, error) {
	return nil, operationNotSupported
}

func (o *String) Gte(other Object) (Object, error) {
	return nil, operationNotSupported
}

func (o *String) Add(other Object) (Object, error) {
	switch other := other.(type) {
	case *String:
		return &String{Value: o.Value + other.Value}, nil
	default:
		return nil, operationNotSupported
	}
}

func (o *String) Sub(other Object) (Object, error) {
	return nil, operationNotSupported
}

func (o *String) Mul(other Object) (Object, error) {
	return nil, operationNotSupported
}

func (o *String) Div(other Object) (Object, error) {
	return nil, operationNotSupported
}
