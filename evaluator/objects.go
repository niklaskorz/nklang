package evaluator

import "github.com/niklaskorz/nklang/ast"

type Object interface {
	IsTrue() bool
	Equals(other Object) (*Boolean, error)
	Lt(other Object) (*Boolean, error)
	Lte(other Object) (*Boolean, error)
	Gt(other Object) (*Boolean, error)
	Gte(other Object) (*Boolean, error)
	Add(other Object) (Object, error)
	Sub(other Object) (Object, error)
	Mul(other Object) (Object, error)
	Div(other Object) (Object, error)
}

type ObjectWithPos interface {
	Pos() (Object, error)
}

type ObjectWithNeg interface {
	Neg() (Object, error)
}

type OperationNotSupportedError struct{}

func (e OperationNotSupportedError) Error() string {
	return "Operation not supported"
}

var operationNotSupported = OperationNotSupportedError{}

type String ast.String

func (o *String) IsTrue() bool {
	return o.Value != ""
}

func (o *String) Equals(other Object) (*Boolean, error) {
	switch other := other.(type) {
	case *String:
		return &Boolean{Value: o.Value == other.Value}, nil
	default:
		return nil, operationNotSupported
	}
}

func (o *String) Lt(other Object) (*Boolean, error) {
	return nil, operationNotSupported
}

func (o *String) Lte(other Object) (*Boolean, error) {
	return nil, operationNotSupported
}

func (o *String) Gt(other Object) (*Boolean, error) {
	return nil, operationNotSupported
}

func (o *String) Gte(other Object) (*Boolean, error) {
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

type Integer ast.Integer

func (o *Integer) IsTrue() bool {
	return o.Value != 0
}

func (o *Integer) Equals(other Object) (*Boolean, error) {
	switch other := other.(type) {
	case *Integer:
		return &Boolean{Value: o.Value == other.Value}, nil
	default:
		return nil, operationNotSupported
	}
}

func (o *Integer) Lt(other Object) (*Boolean, error) {
	switch other := other.(type) {
	case *Integer:
		return &Boolean{Value: o.Value < other.Value}, nil
	default:
		return nil, operationNotSupported
	}
}

func (o *Integer) Lte(other Object) (*Boolean, error) {
	switch other := other.(type) {
	case *Integer:
		return &Boolean{Value: o.Value <= other.Value}, nil
	default:
		return nil, operationNotSupported
	}
}

func (o *Integer) Gt(other Object) (*Boolean, error) {
	switch other := other.(type) {
	case *Integer:
		return &Boolean{Value: o.Value > other.Value}, nil
	default:
		return nil, operationNotSupported
	}
}

func (o *Integer) Gte(other Object) (*Boolean, error) {
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

func (o *Integer) Pos() (Object, error) {
	return o, nil
}

func (o *Integer) Neg() (Object, error) {
	return &Integer{Value: -o.Value}, nil
}

type Boolean ast.Boolean

func (o *Boolean) IsTrue() bool {
	return o.Value
}

func (o *Boolean) Equals(other Object) (*Boolean, error) {
	switch other := other.(type) {
	case *Boolean:
		return &Boolean{Value: o.Value == other.Value}, nil
	default:
		return nil, operationNotSupported
	}
}

func (o *Boolean) Lt(other Object) (*Boolean, error) {
	return nil, operationNotSupported
}

func (o *Boolean) Lte(other Object) (*Boolean, error) {
	return nil, operationNotSupported
}

func (o *Boolean) Gt(other Object) (*Boolean, error) {
	return nil, operationNotSupported
}

func (o *Boolean) Gte(other Object) (*Boolean, error) {
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

type Nil ast.Nil

var NilObject = &Nil{}

func (o *Nil) IsTrue() bool {
	return false
}

func (o *Nil) Equals(other Object) (*Boolean, error) {
	switch other.(type) {
	case *Nil:
		return &Boolean{Value: true}, nil
	default:
		return &Boolean{Value: false}, nil
	}
}

func (o *Nil) Lt(other Object) (*Boolean, error) {
	return nil, operationNotSupported
}

func (o *Nil) Lte(other Object) (*Boolean, error) {
	return nil, operationNotSupported
}

func (o *Nil) Gt(other Object) (*Boolean, error) {
	return nil, operationNotSupported
}

func (o *Nil) Gte(other Object) (*Boolean, error) {
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

type Function struct {
	*ast.Function
	parentScope *definitionScope
}

func (o *Function) IsTrue() bool {
	return true
}

func (o *Function) Equals(other Object) (*Boolean, error) {
	return &Boolean{Value: o == other}, nil
}

func (o *Function) Lt(other Object) (*Boolean, error) {
	return nil, operationNotSupported
}

func (o *Function) Lte(other Object) (*Boolean, error) {
	return nil, operationNotSupported
}

func (o *Function) Gt(other Object) (*Boolean, error) {
	return nil, operationNotSupported
}

func (o *Function) Gte(other Object) (*Boolean, error) {
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

type PredefinedFunction struct {
	fn func(params []Object) (Object, error)
}

func WrapFunction(fn func(params []Object) (Object, error)) *PredefinedFunction {
	return &PredefinedFunction{fn: fn}
}

func (o *PredefinedFunction) IsTrue() bool {
	return true
}

func (o *PredefinedFunction) Equals(other Object) (*Boolean, error) {
	return &Boolean{Value: o == other}, nil
}

func (o *PredefinedFunction) Lt(other Object) (*Boolean, error) {
	return nil, operationNotSupported
}

func (o *PredefinedFunction) Lte(other Object) (*Boolean, error) {
	return nil, operationNotSupported
}

func (o *PredefinedFunction) Gt(other Object) (*Boolean, error) {
	return nil, operationNotSupported
}

func (o *PredefinedFunction) Gte(other Object) (*Boolean, error) {
	return nil, operationNotSupported
}

func (o *PredefinedFunction) Add(other Object) (Object, error) {
	return nil, operationNotSupported
}

func (o *PredefinedFunction) Sub(other Object) (Object, error) {
	return nil, operationNotSupported
}

func (o *PredefinedFunction) Mul(other Object) (Object, error) {
	return nil, operationNotSupported
}

func (o *PredefinedFunction) Div(other Object) (Object, error) {
	return nil, operationNotSupported
}
