package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"

	"github.com/niklaskorz/nklang/evaluator"
	"github.com/niklaskorz/nklang/lexer"
	"github.com/niklaskorz/nklang/parser"
	"github.com/niklaskorz/nklang/semantics"
)

func paramsToString(params []evaluator.Object) string {
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
		case *evaluator.Float:
			s += strconv.FormatFloat(p.Value, 'f', -1, 64)
		case *evaluator.Boolean:
			if p.Value {
				s += "true"
			} else {
				s += "false"
			}
		case *evaluator.Nil:
			s += "nil"
		case *evaluator.Function:
			s += "func("
			for j, param := range p.Function.Parameters {
				if j != 0 {
					s += ", "
				}
				s += param
			}
			s += ")"
		case *evaluator.PredefinedFunction:
			s += "[PredefinedFunction]"
		default:
			s += "[Object]"
		}
	}
	return s
}

func pfPrintln(params []evaluator.Object) (evaluator.Object, error) {
	s := paramsToString(params)
	fmt.Println(s)
	return evaluator.NilObject, nil
}

func pfPrint(params []evaluator.Object) (evaluator.Object, error) {
	s := paramsToString(params)
	fmt.Print(s)
	return evaluator.NilObject, nil
}

func pfInput(params []evaluator.Object) (evaluator.Object, error) {
	pfPrint(params)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	return &evaluator.String{Value: text[:len(text)-1]}, nil
}

func pfEval(params []evaluator.Object) (evaluator.Object, error) {
	src := paramsToString(params)

	s := lexer.NewScanner(strings.NewReader(src))
	if err := s.ReadNext(); err != nil {
		return evaluator.NilObject, errors.Wrap(err, "Scanning eval string failed")
	}

	expr, err := parser.ParseExpression(s)
	if err != nil {
		return evaluator.NilObject, errors.Wrap(err, "Parsing eval string failed")
	}

	if err := semantics.AnalyzeExpression(semantics.NewScope(), expr); err != nil {
		return evaluator.NilObject, errors.Wrap(err, "Analyzing eval string failed")
	}

	result, err := evaluator.EvaluateExpression(expr, evaluator.NewScope())
	if err != nil {
		return evaluator.NilObject, errors.Wrap(err, "Evaluating eval string failed")
	}

	return result, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:", os.Args[0], "<source file>")
		return
	}

	f, err := os.Open(os.Args[1])
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
	ds.Declare("print")
	ds.Declare("input")
	ds.Declare("eval")
	if err := semantics.AnalyzeLookupsWithScope(p, ds); err != nil {
		fmt.Println(err)
		return
	}

	pfPrintln := evaluator.WrapFunction(pfPrintln)
	pfPrint := evaluator.WrapFunction(pfPrint)
	pfInput := evaluator.WrapFunction(pfInput)
	pfEval := evaluator.WrapFunction(pfEval)
	scope := evaluator.NewScope()
	scope.Declare("println", pfPrintln)
	scope.Declare("print", pfPrint)
	scope.Declare("input", pfInput)
	scope.Declare("eval", pfEval)

	if err := evaluator.EvaluateWithScope(p, scope); err != nil {
		fmt.Println(err)
		return
	}
}
