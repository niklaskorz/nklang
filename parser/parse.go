package parser

import (
	"strconv"

	"github.com/niklaskorz/nklang/ast"
	"github.com/niklaskorz/nklang/lexer"
)

func Parse(s *lexer.Scanner) (*ast.Program, error) {
	if err := s.ReadNext(); err != nil {
		return nil, err
	}

	statements := []ast.Statement{}
	for s.Token.Type != lexer.EOF {
		n, err := parseStatement(s)
		if err != nil {
			return nil, err
		}

		statements = append(statements, n)
	}

	p := ast.Program{Statements: statements}
	return &p, nil
}

func parseStatement(s *lexer.Scanner) (ast.Statement, error) {
	if s.Token.Type == lexer.IfKeyword {
		n, err := parseIfStatement(s)
		if err != nil {
			return nil, err
		}
		return n, nil
	}

	if s.Token.Type == lexer.WhileKeyword {
		n, err := parseWhileStatement(s)
		if err != nil {
			return nil, err
		}
		return n, nil
	}

	var n ast.Statement

	if s.Token.Type == lexer.ContinueKeyword {
		if err := s.ReadNext(); err != nil {
			return nil, err
		}
		n = &ast.ContinueStatement{}
	} else if s.Token.Type == lexer.BreakKeyword {
		if err := s.ReadNext(); err != nil {
			return nil, err
		}
		n = &ast.BreakStatement{}
	} else if s.Token.Type == lexer.ReturnKeyword {
		if err := s.ReadNext(); err != nil {
			return nil, err
		}
		e, err := parseExpression(s)
		if err != nil {
			return nil, err
		}
		n = &ast.ReturnStatement{Expression: e}
	} else if s.Token.Type == lexer.ID {
		identifier := s.Token.Value
		if err := s.ReadNext(); err != nil {
			return nil, err
		}
		if s.Token.Type == lexer.DeclarationOperator {
			if err := s.ReadNext(); err != nil {
				return nil, err
			}
			v, err := parseExpression(s)
			if err != nil {
				return nil, err
			}
			n = &ast.DeclarationStatement{Identifier: identifier, Value: v}
		} else if s.Token.Type == lexer.AssignmentOperator {
			if err := s.ReadNext(); err != nil {
				return nil, err
			}
			v, err := parseExpression(s)
			if err != nil {
				return nil, err
			}
			n = &ast.AssignmentStatement{Identifier: identifier, Value: v}
		} else {
			if err := s.Unread(); err != nil {
				return nil, err
			}
			e, err := parseExpression(s)
			if err != nil {
				return nil, err
			}
			n = &ast.ExpressionStatement{Expression: e}
		}
	} else {
		e, err := parseExpression(s)
		if err != nil {
			return nil, err
		}
		n = &ast.ExpressionStatement{Expression: e}
	}

	if s.Token.Type != lexer.Semicolon {
		return nil, unexpectedToken(s.Token, ";")
	}
	if err := s.ReadNext(); err != nil {
		return nil, err
	}
	return n, nil
}

func parseIfStatement(s *lexer.Scanner) (*ast.IfStatement, error) {
	if s.Token.Type != lexer.IfKeyword {
		return nil, unexpectedToken(s.Token, "if")
	}
	if err := s.ReadNext(); err != nil {
		return nil, err
	}

	e, err := parseExpression(s)
	if err != nil {
		return nil, err
	}

	statements, err := parseStatementBlock(s)
	if err != nil {
		return nil, err
	}

	n := &ast.IfStatement{Condition: e, Statements: statements}

	if s.Token.Type == lexer.ElseKeyword {
		if err := s.ReadNext(); err != nil {
			return nil, err
		}

		if s.Token.Type == lexer.IfKeyword {
			elseBranch, err := parseIfStatement(s)
			if err != nil {
				return nil, err
			}
			n.ElseBranch = elseBranch
		} else if s.Token.Type == lexer.LeftBrace {
			statements, err := parseStatementBlock(s)
			if err != nil {
				return nil, err
			}
			n.ElseBranch = &ast.IfStatement{Statements: statements}
		}
	}

	return n, nil
}

func parseWhileStatement(s *lexer.Scanner) (*ast.WhileStatement, error) {
	if s.Token.Type != lexer.WhileKeyword {
		return nil, unexpectedToken(s.Token, "while")
	}
	if err := s.ReadNext(); err != nil {
		return nil, err
	}

	e, err := parseExpression(s)
	if err != nil {
		return nil, err
	}

	statements, err := parseStatementBlock(s)
	if err != nil {
		return nil, err
	}

	return &ast.WhileStatement{Condition: e, Statements: statements}, nil
}

func parseStatementBlock(s *lexer.Scanner) ([]ast.Statement, error) {
	if s.Token.Type != lexer.LeftBrace {
		return nil, unexpectedToken(s.Token, "{")
	}
	if err := s.ReadNext(); err != nil {
		return nil, err
	}

	statements := []ast.Statement{}
	for s.Token.Type != lexer.RightBrace {
		n, err := parseStatement(s)
		if err != nil {
			return nil, err
		}
		statements = append(statements, n)
	}
	if err := s.ReadNext(); err != nil {
		return nil, err
	}

	return statements, nil
}

func parseExpression(s *lexer.Scanner) (ast.Expression, error) {
	switch s.Token.Type {
	case lexer.IfKeyword:
		return parseIfExpression(s)
	case lexer.FunctionKeyword:
		return parseFunction(s)
	default:
		return parseLogicalOr(s)
	}
}

func parseIfExpression(s *lexer.Scanner) (*ast.IfExpression, error) {
	if s.Token.Type != lexer.IfKeyword {
		return nil, unexpectedToken(s.Token, "if")
	}
	if err := s.ReadNext(); err != nil {
		return nil, err
	}

	cond, err := parseExpression(s)
	if err != nil {
		return nil, err
	}

	if s.Token.Type != lexer.LeftBrace {
		return nil, unexpectedToken(s.Token, "{")
	}
	if err := s.ReadNext(); err != nil {
		return nil, err
	}

	v, err := parseExpression(s)
	if err != nil {
		return nil, err
	}

	if s.Token.Type != lexer.RightBrace {
		return nil, unexpectedToken(s.Token, "}")
	}
	if err := s.ReadNext(); err != nil {
		return nil, err
	}

	n := &ast.IfExpression{Condition: cond, Value: v}

	if s.Token.Type != lexer.ElseKeyword {
		return nil, unexpectedToken(s.Token, "else")
	}
	if err := s.ReadNext(); err != nil {
		return nil, err
	}

	if s.Token.Type == lexer.IfKeyword {
		elseBranch, err := parseIfExpression(s)
		if err != nil {
			return nil, err
		}
		n.ElseBranch = elseBranch
	} else if s.Token.Type == lexer.LeftBrace {
		if err := s.ReadNext(); err != nil {
			return nil, err
		}

		v, err := parseExpression(s)
		if err != nil {
			return nil, err
		}

		if s.Token.Type != lexer.RightBrace {
			return nil, unexpectedToken(s.Token, "}")
		}
		if err := s.ReadNext(); err != nil {
			return nil, err
		}

		n.ElseBranch = &ast.IfExpression{Value: v}
	} else {
		return nil, unexpectedToken(s.Token, "one of: if, {")
	}

	return n, nil
}

func parseFunction(s *lexer.Scanner) (*ast.Function, error) {
	if s.Token.Type != lexer.FunctionKeyword {
		return nil, unexpectedToken(s.Token, "func")
	}
	if err := s.ReadNext(); err != nil {
		return nil, err
	}

	if s.Token.Type != lexer.LeftParen {
		return nil, unexpectedToken(s.Token, "(")
	}
	if err := s.ReadNext(); err != nil {
		return nil, err
	}

	parameters := []string{}

	for s.Token.Type == lexer.ID {
		parameters = append(parameters, s.Token.Value)
		if err := s.ReadNext(); err != nil {
			return nil, err
		}

		if s.Token.Type != lexer.Comma {
			break
		}
		if err := s.ReadNext(); err != nil {
			return nil, err
		}
	}

	if s.Token.Type != lexer.RightParen {
		return nil, unexpectedToken(s.Token, ")")
	}
	if err := s.ReadNext(); err != nil {
		return nil, err
	}

	statements, err := parseStatementBlock(s)
	if err != nil {
		return nil, err
	}

	return &ast.Function{Parameters: parameters, Statements: statements}, nil
}

func parseLogicalOr(s *lexer.Scanner) (ast.Expression, error) {
	expr, err := parseLogicalAnd(s)
	if err != nil {
		return nil, err
	}

	for s.Token.Type == lexer.LogicalOr {
		if err := s.ReadNext(); err != nil {
			return nil, err
		}

		e, err := parseLogicalAnd(s)
		if err != nil {
			return nil, err
		}

		expr = &ast.BinaryOperationExpression{Operator: ast.BinaryOperatorLor, A: expr, B: e}
	}

	return expr, nil
}

func parseLogicalAnd(s *lexer.Scanner) (ast.Expression, error) {
	expr, err := parseComparison(s)
	if err != nil {
		return nil, err
	}

	for s.Token.Type == lexer.LogicalAnd {
		if err := s.ReadNext(); err != nil {
			return nil, err
		}

		e, err := parseComparison(s)
		if err != nil {
			return nil, err
		}

		expr = &ast.BinaryOperationExpression{Operator: ast.BinaryOperatorLand, A: expr, B: e}
	}

	return expr, nil
}

func parseComparison(s *lexer.Scanner) (ast.Expression, error) {
	expr, err := parseTerm(s)
	if err != nil {
		return nil, err
	}

	var op ast.BinaryOperator = -1
	switch s.Token.Type {
	case lexer.EqOperator:
		op = ast.BinaryOperatorEq
	case lexer.NeOperator:
		op = ast.BinaryOperatorNe
	case lexer.LtOperator:
		op = ast.BinaryOperatorLt
	case lexer.LeOperator:
		op = ast.BinaryOperatorLe
	case lexer.GtOperator:
		op = ast.BinaryOperatorGt
	case lexer.GeOperator:
		op = ast.BinaryOperatorGe
	}

	if op != -1 {
		if err := s.ReadNext(); err != nil {
			return nil, err
		}

		e, err := parseTerm(s)
		if err != nil {
			return nil, err
		}

		return &ast.BinaryOperationExpression{
			Operator: op,
			A:        expr,
			B:        e,
		}, nil
	}

	return expr, nil
}

func parseTerm(s *lexer.Scanner) (ast.Expression, error) {
	expr, err := parseAddend(s)
	if err != nil {
		return nil, err
	}

L:
	for {
		var op ast.BinaryOperator
		switch s.Token.Type {
		case lexer.Plus:
			op = ast.BinaryOperatorAdd
		case lexer.Minus:
			op = ast.BinaryOperatorSub
		default:
			break L
		}

		if err := s.ReadNext(); err != nil {
			return nil, err
		}

		e, err := parseAddend(s)
		if err != nil {
			return nil, err
		}

		expr = &ast.BinaryOperationExpression{Operator: op, A: expr, B: e}
	}

	return expr, nil
}

func parseAddend(s *lexer.Scanner) (ast.Expression, error) {
	expr, err := parseFactor(s)
	if err != nil {
		return nil, err
	}

L:
	for {
		var op ast.BinaryOperator
		switch s.Token.Type {
		case lexer.MulOperator:
			op = ast.BinaryOperatorMul
		case lexer.DivOperator:
			op = ast.BinaryOperatorDiv
		default:
			break L
		}

		if err := s.ReadNext(); err != nil {
			return nil, err
		}

		e, err := parseFactor(s)
		if err != nil {
			return nil, err
		}

		expr = &ast.BinaryOperationExpression{Operator: op, A: expr, B: e}
	}

	return expr, nil
}

func parseFactor(s *lexer.Scanner) (ast.Expression, error) {
	var e *ast.UnaryOperationExpression
L:
	for {
		var op ast.UnaryOperator
		switch s.Token.Type {
		case lexer.LogicalNot:
			op = ast.UnaryOperatorLnot
		case lexer.Plus:
			op = ast.UnaryOperatorPos
		case lexer.Minus:
			op = ast.UnaryOperatorNeg
		default:
			break L
		}

		if err := s.ReadNext(); err != nil {
			return nil, err
		}

		operation := &ast.UnaryOperationExpression{Operator: op}
		if e != nil {
			e.A = operation
		}
		e = operation
	}

	v, err := parseValue(s)
	if err != nil {
		return nil, err
	}

	if e != nil {
		e.A = v
		return e, nil
	}
	return v, nil
}

func parseValue(s *lexer.Scanner) (ast.Expression, error) {
	switch s.Token.Type {
	case lexer.LeftParen:
		if err := s.ReadNext(); err != nil {
			return nil, err
		}
		n, err := parseExpression(s)
		if err != nil {
			return nil, err
		}
		if s.Token.Type != lexer.RightParen {
			return nil, unexpectedToken(s.Token, ")")
		}
		if err := s.ReadNext(); err != nil {
			return nil, err
		}
		if s.Token.Type == lexer.LeftParen {
			return parseCall(n, s)
		}
		return n, nil
	case lexer.ID:
		n := &ast.LookupExpression{Identifier: s.Token.Value}
		if err := s.ReadNext(); err != nil {
			return nil, err
		}
		if s.Token.Type == lexer.LeftParen {
			return parseCall(n, s)
		}
		return n, nil
	case lexer.Integer:
		num, err := strconv.ParseInt(s.Token.Value, 10, 64)
		if err != nil {
			return nil, err
		}
		n := &ast.Integer{Value: num}
		if err := s.ReadNext(); err != nil {
			return nil, err
		}
		return n, nil
	case lexer.String:
		n := &ast.String{Value: s.Token.Value}
		if err := s.ReadNext(); err != nil {
			return nil, err
		}
		return n, nil
	case lexer.TrueKeyword:
		n := &ast.Boolean{Value: true}
		if err := s.ReadNext(); err != nil {
			return nil, err
		}
		return n, nil
	case lexer.FalseKeyword:
		n := &ast.Boolean{Value: false}
		if err := s.ReadNext(); err != nil {
			return nil, err
		}
		return n, nil
	case lexer.NilKeyword:
		n := &ast.Nil{}
		if err := s.ReadNext(); err != nil {
			return nil, err
		}
		return n, nil
	}

	return nil, unexpectedToken(s.Token, "one of: (, ID, Integer, String, true, false")
}

func parseCall(callee ast.Expression, s *lexer.Scanner) (ast.Expression, error) {
	if s.Token.Type != lexer.LeftParen {
		return nil, unexpectedToken(s.Token, "(")
	}
	if err := s.ReadNext(); err != nil {
		return nil, err
	}

	parameters := []ast.Expression{}

	for {
		e, err := parseExpression(s)
		if err != nil {
			return nil, err
		}

		parameters = append(parameters, e)

		if s.Token.Type != lexer.Comma {
			break
		}
		if err := s.ReadNext(); err != nil {
			return nil, err
		}
	}

	if s.Token.Type != lexer.RightParen {
		return nil, unexpectedToken(s.Token, ")")
	}
	if err := s.ReadNext(); err != nil {
		return nil, err
	}

	return &ast.CallExpression{Callee: callee, Parameters: parameters}, nil
}
