package ast

type Program struct {
	Statements []Node
}

func (p Program) String() string {
	str := ""
	for _, s := range p.Statements {
		str += s.String() + ";\n"
	}
	return str
}

func (p Program) Evaluate() Node {
	var ret Node = Void{}
	for _, s := range p.Statements {
		ret = s.Evaluate()
	}
	return ret
}

func (p Program) IsTrue() bool {
	return true
}
