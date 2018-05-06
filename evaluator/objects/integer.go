package objects

import "niklaskorz.de/nklang/ast"

type Integer ast.Integer

func (o *Integer) IsTrue() bool {
	return o.Value != 0
}

func (o *Integer) Equals(other Object) (Object, error) {
	switch other := other.(type) {
	case *Integer:
		return &Boolean{Value: o.Value == other.Value}, nil
	default:
		return nil, operationNotSupported
	}
}

func (o *Integer) Lt(other Object) (Object, error) {
	switch other := other.(type) {
	case *Integer:
		return &Boolean{Value: o.Value < other.Value}, nil
	default:
		return nil, operationNotSupported
	}
}

func (o *Integer) Lte(other Object) (Object, error) {
	switch other := other.(type) {
	case *Integer:
		return &Boolean{Value: o.Value <= other.Value}, nil
	default:
		return nil, operationNotSupported
	}
}

func (o *Integer) Gt(other Object) (Object, error) {
	switch other := other.(type) {
	case *Integer:
		return &Boolean{Value: o.Value > other.Value}, nil
	default:
		return nil, operationNotSupported
	}
}

func (o *Integer) Gte(other Object) (Object, error) {
	switch other := other.(type) {
	case *Integer:
		return &Boolean{Value: o.Value >= other.Value}, nil
	default:
		return nil, operationNotSupported
	}
}

func (o *Integer) Add(other Object) (Object, error) {
	switch other := other.(type) {
	case *Integer:
		return &Integer{Value: o.Value + other.Value}, nil
	default:
		return nil, operationNotSupported
	}
}

func (o *Integer) Sub(other Object) (Object, error) {
	switch other := other.(type) {
	case *Integer:
		return &Integer{Value: o.Value - other.Value}, nil
	default:
		return nil, operationNotSupported
	}
}

func (o *Integer) Mul(other Object) (Object, error) {
	switch other := other.(type) {
	case *Integer:
		return &Integer{Value: o.Value * other.Value}, nil
	default:
		return nil, operationNotSupported
	}
}

func (o *Integer) Div(other Object) (Object, error) {
	switch other := other.(type) {
	case *Integer:
		return &Integer{Value: o.Value / other.Value}, nil
	default:
		return nil, operationNotSupported
	}
}
