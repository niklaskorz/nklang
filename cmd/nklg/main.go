package main

import (
	"fmt"
	"os"
	"strconv"

	"niklaskorz.de/nklang/evaluator"

	"niklaskorz.de/nklang/lexer"
	"niklaskorz.de/nklang/parser"
	"niklaskorz.de/nklang/semantics"
)

func pfPrintln(params []evaluator.Object) (evaluator.Object, error) {
	s := ""
	for i, p := range params {
		if i != 0 {
			s += " "
		}
		switch p := p.(type) {
		case *evaluator.String:
			s += p.Value
		case *evaluator.Integer:
			s += strconv.FormatInt(p.Value, 10)
		case *evaluator.Boolean:
			if p.Value {
				s += "true"
			} else {
				s += "false"
			}
		case *evaluator.Nil:
			s += "nil"
		case *evaluator.Function:
			s += "[Function]"
		case *evaluator.PredefinedFunction:
			s += "[PredefinedFunction]"
		default:
			s += "[Object]"
		}
	}

	fmt.Println(s)
	return evaluator.NilObject, nil
}

func main() {
	f, err := os.Open("example.nk")
	if err != nil {
		fmt.Println(err)
		return
	}

	s := lexer.NewScanner(f)
	p, err := parser.Parse(s)
	if err != nil {
		fmt.Println(err)
		return
	}

	ds := semantics.NewScope()
	ds.Declare("println")
	if err := semantics.AnalyzeLookupsWithScope(p, ds); err != nil {
		fmt.Println(err)
		return
	}

	pfPrintln := evaluator.PredefinedFunction(pfPrintln)
	scope := evaluator.NewScope()
	scope.Declare("println", pfPrintln)

	if err := evaluator.EvaluateWithScope(p, scope); err != nil {
		fmt.Println(err)
		return
	}
}
