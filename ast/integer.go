package ast

import "strconv"

type Integer struct {
	Value int64
}

func (i Integer) String() string {
	return strconv.FormatInt(i.Value, 10)
}

func (i Integer) Evaluate() Node {
	return i
}

func (i Integer) IsTrue() bool {
	return i.Value != 0
}
