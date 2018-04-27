package parser

import (
	"strconv"

	"niklaskorz.de/nklang/ast"
	"niklaskorz.de/nklang/lexer"
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
		n = ast.ContinueStatement{}
	} else if s.Token.Type == lexer.BreakKeyword {
		if err := s.ReadNext(); err != nil {
			return nil, err
		}
		n = ast.BreakStatement{}
	} else if s.Token.Type == lexer.ReturnKeyword {
		if err := s.ReadNext(); err != nil {
			return nil, err
		}
		e, err := parseExpression(s)
		if err != nil {
			return nil, err
		}
		n = ast.ReturnStatement{Expression: e}
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
			n = ast.DeclarationStatement{Identifier: identifier, Value: v}
		} else if s.Token.Type == lexer.AssignmentOperator {
			if err := s.ReadNext(); err != nil {
				return nil, err
			}
			v, err := parseExpression(s)
			if err != nil {
				return nil, err
			}
			n = ast.AssignmentStatement{Identifier: identifier, Value: v}
		} else {
			if err := s.Unread(); err != nil {
				return nil, err
			}
			e, err := parseExpression(s)
			if err != nil {
				return nil, err
			}
			n = ast.ExpressionStatement{Expression: e}
		}
	} else {
		e, err := parseExpression(s)
		if err != nil {
			return nil, err
		}
		n = ast.ExpressionStatement{Expression: e}
	}

	if s.Token.Type != lexer.Semicolon {
		return nil, unexpectedToken(s.Token)
	}
	if err := s.ReadNext(); err != nil {
		return nil, err
	}
	return n, nil
}

func parseIfStatement(s *lexer.Scanner) (*ast.IfStatement, error) {
	if s.Token.Type != lexer.IfKeyword {
		return nil, unexpectedToken(s.Token)
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
		return nil, unexpectedToken(s.Token)
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
		return nil, unexpectedToken(s.Token)
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
		return nil, unexpectedToken(s.Token)
	}
	if err := s.ReadNext(); err != nil {
		return nil, err
	}

	cond, err := parseExpression(s)
	if err != nil {
		return nil, err
	}

	if s.Token.Type != lexer.LeftBrace {
		return nil, unexpectedToken(s.Token)
	}
	if err := s.ReadNext(); err != nil {
		return nil, err
	}

	v, err := parseExpression(s)
	if err != nil {
		return nil, err
	}

	if s.Token.Type != lexer.RightBrace {
		return nil, unexpectedToken(s.Token)
	}
	if err := s.ReadNext(); err != nil {
		return nil, err
	}

	n := &ast.IfExpression{Condition: cond, Value: v}

	if s.Token.Type != lexer.ElseKeyword {
		return nil, unexpectedToken(s.Token)
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
			return nil, unexpectedToken(s.Token)
		}
		if err := s.ReadNext(); err != nil {
			return nil, err
		}

		n.ElseBranch = &ast.IfExpression{Value: v}
	} else {
		return nil, unexpectedToken(s.Token)
	}

	return n, nil
}

func parseFunction(s *lexer.Scanner) (*ast.Function, error) {
	if s.Token.Type != lexer.FunctionKeyword {
		return nil, unexpectedToken(s.Token)
	}
	if err := s.ReadNext(); err != nil {
		return nil, err
	}

	if s.Token.Type != lexer.LeftParen {
		return nil, unexpectedToken(s.Token)
	}
	if err := s.ReadNext(); err != nil {
		return nil, err
	}

	parameters := []string{}

	for s.Token.Type == lexer.ID {
		if err := s.ReadNext(); err != nil {
			return nil, err
		}
		parameters = append(parameters, s.Token.Value)

		if s.Token.Type != lexer.Comma {
			break
		}
		if err := s.ReadNext(); err != nil {
			return nil, err
		}
	}

	if s.Token.Type != lexer.RightParen {
		return nil, unexpectedToken(s.Token)
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

		expr = ast.LogicalOrExpression{A: expr, B: e}
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

		expr = ast.LogicalAndExpression{A: expr, B: e}
	}

	return expr, nil
}

func parseComparison(s *lexer.Scanner) (ast.Expression, error) {
	expr, err := parseTerm(s)
	if err != nil {
		return nil, err
	}

	var op ast.ComparisonOperator = -1
	switch s.Token.Type {
	case lexer.EqOperator:
		op = ast.ComparisonOperatorEq
	case lexer.LtOperator:
		op = ast.ComparisonOperatorLt
	case lexer.LeOperator:
		op = ast.ComparisonOperatorLe
	case lexer.GtOperator:
		op = ast.ComparisonOperatorGt
	case lexer.GeOperator:
		op = ast.ComparisonOperatorGe
	}

	if op != -1 {
		if err := s.ReadNext(); err != nil {
			return nil, err
		}

		e, err := parseTerm(s)
		if err != nil {
			return nil, err
		}

		return ast.ComparisonExpression{
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

	for s.Token.Type == lexer.AddOperator || s.Token.Type == lexer.SubOperator {
		isAddition := s.Token.Type == lexer.AddOperator
		if err := s.ReadNext(); err != nil {
			return nil, err
		}

		e, err := parseAddend(s)
		if err != nil {
			return nil, err
		}

		if isAddition {
			expr = ast.AdditionExpression{A: expr, B: e}
		} else {
			expr = ast.SubstractionExpression{A: expr, B: e}
		}
	}

	return expr, nil
}

func parseAddend(s *lexer.Scanner) (ast.Expression, error) {
	expr, err := parseFactor(s)
	if err != nil {
		return nil, err
	}

	for s.Token.Type == lexer.MulOperator || s.Token.Type == lexer.DivOperator {
		isMultiplication := s.Token.Type == lexer.MulOperator
		if err := s.ReadNext(); err != nil {
			return nil, err
		}

		e, err := parseFactor(s)
		if err != nil {
			return nil, err
		}

		if isMultiplication {
			expr = ast.MultiplicationExpression{A: expr, B: e}
		} else {
			expr = ast.DivisionExpression{A: expr, B: e}
		}
	}

	return expr, nil
}

func parseFactor(s *lexer.Scanner) (ast.Expression, error) {
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
			return nil, unexpectedToken(s.Token)
		}
		if err := s.ReadNext(); err != nil {
			return nil, err
		}
		if s.Token.Type == lexer.LeftParen {
			return parseCall(n, s)
		}
		return n, nil
	case lexer.ID:
		n := ast.LookupExpression{Identifier: s.Token.Value}
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
		n := ast.Integer{Value: num}
		if err := s.ReadNext(); err != nil {
			return nil, err
		}
		return n, nil
	case lexer.String:
		n := ast.String{Value: s.Token.Value}
		if err := s.ReadNext(); err != nil {
			return nil, err
		}
		return n, nil
	case lexer.TrueKeyword:
		n := ast.Boolean{Value: true}
		if err := s.ReadNext(); err != nil {
			return nil, err
		}
		return n, nil
	case lexer.FalseKeyword:
		n := ast.Boolean{Value: false}
		if err := s.ReadNext(); err != nil {
			return nil, err
		}
		return n, nil
	}

	return nil, unexpectedToken(s.Token)
}

func parseCall(callee ast.Expression, s *lexer.Scanner) (ast.Expression, error) {
	if s.Token.Type != lexer.LeftParen {
		return nil, unexpectedToken(s.Token)
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
		return nil, unexpectedToken(s.Token)
	}
	if err := s.ReadNext(); err != nil {
		return nil, err
	}

	return ast.CallExpression{Callee: callee, Parameters: parameters}, nil
}
