package objects

type Object interface {
	IsTrue() bool
	Equals(other Object) (Object, error)
	Lt(other Object) (Object, error)
	Lte(other Object) (Object, error)
	Gt(other Object) (Object, error)
	Gte(other Object) (Object, error)
	Add(other Object) (Object, error)
	Sub(other Object) (Object, error)
	Mul(other Object) (Object, error)
	Div(other Object) (Object, error)
}

type OperationNotSupportedError struct{}

func (e OperationNotSupportedError) Error() string {
	return "Operation not supported"
}

var operationNotSupported = OperationNotSupportedError{}
