package main

import (
	"fmt"
	"os"
	"strconv"

	"niklaskorz.de/nklang/evaluator"
	"niklaskorz.de/nklang/evaluator/objects"

	"niklaskorz.de/nklang/lexer"
	"niklaskorz.de/nklang/parser"
	"niklaskorz.de/nklang/semantics"
)

func pfPrintln(params []objects.Object) (objects.Object, error) {
	s := ""
	for i, p := range params {
		if i != 0 {
			s += " "
		}
		switch p := p.(type) {
		case *objects.String:
			s += p.Value
		case *objects.Integer:
			s += strconv.FormatInt(p.Value, 10)
		case *objects.Boolean:
			if p.Value {
				s += "true"
			} else {
				s += "false"
			}
		case *objects.Nil:
			s += "nil"
		case *objects.Function:
			s += "[Function]"
		case *objects.PredefinedFunction:
			s += "[PredefinedFunction]"
		default:
			s += "[Object]"
		}
	}

	fmt.Println(s)
	return objects.NilObject, nil
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

	pfPrintln := objects.PredefinedFunction(pfPrintln)
	scope := evaluator.NewScope()
	scope.Declare("println", pfPrintln)

	if err := evaluator.EvaluateWithScope(p, scope); err != nil {
		fmt.Println(err)
		return
	}
}
