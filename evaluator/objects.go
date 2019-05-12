package evaluator

import "github.com/niklaskorz/nklang/ast"

type Object interface {
	IsTrue() bool
	Equals(other Object) (*Boolean, error)
}

type ObjectWithPos interface {
	Pos() (Object, error)
}

type ObjectWithNeg interface {
	Neg() (Object, error)
}

type Comparable interface {
	Lt(other Object) (*Boolean, error)
	Lte(other Object) (*Boolean, error)
	Gt(other Object) (*Boolean, error)
	Gte(other Object) (*Boolean, error)
}

type Addable interface {
	Add(other Object) (Object, error)
}

type Subtractable interface {
	Sub(other Object) (Object, error)
}

type Multipliable interface {
	Mul(other Object) (Object, error)
}

type Dividable interface {
	Div(other Object) (Object, error)
}

type Callable interface {
	Call()
}

type String ast.String

func (o *String) IsTrue() bool {
	return o.Value != ""
}

func (o *String) Equals(other Object) (*Boolean, error) {
	switch other := other.(type) {
	case *String:
		return &Boolean{Value: o.Value == other.Value}, nil
	}
	return &Boolean{Value: false}, nil
}

func (o *String) Add(other Object) (Object, error) {
	switch other := other.(type) {
	case *String:
		return &String{Value: o.Value + other.Value}, nil
	}
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
	case *Float:
		return &Boolean{Value: float64(o.Value) == other.Value}, nil
	}
	return &Boolean{Value: false}, nil
}

func (o *Integer) Lt(other Object) (*Boolean, error) {
	switch other := other.(type) {
	case *Integer:
		return &Boolean{Value: o.Value < other.Value}, nil
	case *Float:
		return &Boolean{Value: float64(o.Value) < other.Value}, nil
	}
	return nil, operationNotSupported
}

func (o *Integer) Lte(other Object) (*Boolean, error) {
	switch other := other.(type) {
	case *Integer:
		return &Boolean{Value: o.Value <= other.Value}, nil
	case *Float:
		return &Boolean{Value: float64(o.Value) <= other.Value}, nil
	}
	return nil, operationNotSupported
}

func (o *Integer) Gt(other Object) (*Boolean, error) {
	switch other := other.(type) {
	case *Integer:
		return &Boolean{Value: o.Value > other.Value}, nil
	case *Float:
		return &Boolean{Value: float64(o.Value) > other.Value}, nil
	}
	return nil, operationNotSupported
}

func (o *Integer) Gte(other Object) (*Boolean, error) {
	switch other := other.(type) {
	case *Integer:
		return &Boolean{Value: o.Value >= other.Value}, nil
	case *Float:
		return &Boolean{Value: float64(o.Value) >= other.Value}, nil
	}
	return nil, operationNotSupported
}

func (o *Integer) Add(other Object) (Object, error) {
	switch other := other.(type) {
	case *Integer:
		return &Integer{Value: o.Value + other.Value}, nil
	case *Float:
		return &Float{Value: float64(o.Value) + other.Value}, nil
	}
	return nil, operationNotSupported
}

func (o *Integer) Sub(other Object) (Object, error) {
	switch other := other.(type) {
	case *Integer:
		return &Integer{Value: o.Value - other.Value}, nil
	case *Float:
		return &Float{Value: float64(o.Value) - other.Value}, nil
	}
	return nil, operationNotSupported
}

func (o *Integer) Mul(other Object) (Object, error) {
	switch other := other.(type) {
	case *Integer:
		return &Integer{Value: o.Value * other.Value}, nil
	case *Float:
		return &Float{Value: float64(o.Value) * other.Value}, nil
	}
	return nil, operationNotSupported
}

func (o *Integer) Div(other Object) (Object, error) {
	switch other := other.(type) {
	case *Integer:
		return &Integer{Value: o.Value / other.Value}, nil
	case *Float:
		return &Float{Value: float64(o.Value) / other.Value}, nil
	}
	return nil, operationNotSupported
}

func (o *Integer) Pos() (Object, error) {
	return o, nil
}

func (o *Integer) Neg() (Object, error) {
	return &Integer{Value: -o.Value}, nil
}

type Float ast.Float

func (o *Float) IsTrue() bool {
	return o.Value != 0
}

func (o *Float) Equals(other Object) (*Boolean, error) {
	switch other := other.(type) {
	case *Integer:
		return &Boolean{Value: o.Value == float64(other.Value)}, nil
	case *Float:
		return &Boolean{Value: o.Value == other.Value}, nil
	}
	return &Boolean{Value: false}, nil
}

func (o *Float) Lt(other Object) (*Boolean, error) {
	switch other := other.(type) {
	case *Integer:
		return &Boolean{Value: o.Value < float64(other.Value)}, nil
	case *Float:
		return &Boolean{Value: o.Value < other.Value}, nil
	}
	return nil, operationNotSupported
}

func (o *Float) Lte(other Object) (*Boolean, error) {
	switch other := other.(type) {
	case *Integer:
		return &Boolean{Value: o.Value <= float64(other.Value)}, nil
	case *Float:
		return &Boolean{Value: o.Value <= other.Value}, nil
	}
	return nil, operationNotSupported
}

func (o *Float) Gt(other Object) (*Boolean, error) {
	switch other := other.(type) {
	case *Integer:
		return &Boolean{Value: o.Value > float64(other.Value)}, nil
	case *Float:
		return &Boolean{Value: o.Value > other.Value}, nil
	}
	return nil, operationNotSupported
}

func (o *Float) Gte(other Object) (*Boolean, error) {
	switch other := other.(type) {
	case *Integer:
		return &Boolean{Value: o.Value >= float64(other.Value)}, nil
	case *Float:
		return &Boolean{Value: o.Value >= other.Value}, nil
	}
	return nil, operationNotSupported
}

func (o *Float) Add(other Object) (Object, error) {
	switch other := other.(type) {
	case *Integer:
		return &Float{Value: o.Value + float64(other.Value)}, nil
	case *Float:
		return &Float{Value: o.Value + other.Value}, nil
	}
	return nil, operationNotSupported
}

func (o *Float) Sub(other Object) (Object, error) {
	switch other := other.(type) {
	case *Integer:
		return &Float{Value: o.Value - float64(other.Value)}, nil
	case *Float:
		return &Float{Value: o.Value - other.Value}, nil
	}
	return nil, operationNotSupported
}

func (o *Float) Mul(other Object) (Object, error) {
	switch other := other.(type) {
	case *Integer:
		return &Float{Value: o.Value * float64(other.Value)}, nil
	case *Float:
		return &Float{Value: o.Value * other.Value}, nil
	}
	return nil, operationNotSupported
}

func (o *Float) Div(other Object) (Object, error) {
	switch other := other.(type) {
	case *Integer:
		return &Float{Value: o.Value / float64(other.Value)}, nil
	case *Float:
		return &Float{Value: o.Value / other.Value}, nil
	}
	return nil, operationNotSupported
}

func (o *Float) Pos() (Object, error) {
	return o, nil
}

func (o *Float) Neg() (Object, error) {
	return &Float{Value: -o.Value}, nil
}

type Boolean ast.Boolean

func (o *Boolean) IsTrue() bool {
	return o.Value
}

func (o *Boolean) Equals(other Object) (*Boolean, error) {
	switch other := other.(type) {
	case *Boolean:
		return &Boolean{Value: o.Value == other.Value}, nil
	}
	return &Boolean{Value: false}, nil
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
	}
	return &Boolean{Value: false}, nil
}

type Function struct {
	*ast.Function
	parentScope *DefinitionScope
}

func (o *Function) IsTrue() bool {
	return true
}

func (o *Function) Equals(other Object) (*Boolean, error) {
	return &Boolean{Value: o == other}, nil
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
