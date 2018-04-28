package evaluator

func (n *LogicalOrExpression) Evaluate() (Object, error) {
	a, err := n.A.Evaluate()
	if err != nil {
		return nil, err
	}
	if a.IsTrue() {
		return a, nil
	}
	return n.B.Evaluate()
}

func (n *LogicalAndExpression) Evaluate() (Object, error) {
	a, err := n.A.Evaluate()
	if err != nil {
		return nil, err
	}
	if a.IsTrue() {
		return n.B.Evaluate()
	}
	return a, nil
}
