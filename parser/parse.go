package parser

import (
	"fmt"
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
		return nil, fmt.Errorf("Unexpected token %s", s.Token)
	}
	if err := s.ReadNext(); err != nil {
		return nil, err
	}
	return n, nil
}

func parseIfStatement(s *lexer.Scanner) (*ast.IfStatement, error) {
	if s.Token.Type != lexer.IfKeyword {
		return nil, fmt.Errorf("Unexpected token %s", s.Token)
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
		return nil, fmt.Errorf("Unexpected token %s", s.Token)
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
		return nil, fmt.Errorf("Unexpected token %s", s.Token)
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
	case lexer.LeftParen:
		if err := s.ReadNext(); err != nil {
			return nil, err
		}
		n, err := parseExpression(s)
		if err != nil {
			return nil, err
		}
		if s.Token.Type != lexer.RightParen {
			return nil, fmt.Errorf("Unexpected token %s", s.Token)
		}
		if err := s.ReadNext(); err != nil {
			return nil, err
		}
		return n, nil
	case lexer.ID:
		n := ast.LookupExpression{Identifier: s.Token.Value}
		if err := s.ReadNext(); err != nil {
			return nil, err
		}
		return n, nil
	case lexer.Integer:
		num, err := strconv.ParseInt(s.Token.Value, 10, 64)
		if err != nil {
			return nil, err
		}
		n := ast.IntegerExpression{Value: num}
		if err := s.ReadNext(); err != nil {
			return nil, err
		}
		return n, nil
	case lexer.String:
		n := ast.StringExpression{Value: s.Token.Value}
		if err := s.ReadNext(); err != nil {
			return nil, err
		}
		return n, nil
	}

	return nil, fmt.Errorf("Unexpected token %s", s.Token)
}

func parseIfExpression(s *lexer.Scanner) (*ast.IfExpression, error) {
	if s.Token.Type != lexer.IfKeyword {
		return nil, fmt.Errorf("Unexpected token %s", s.Token)
	}
	if err := s.ReadNext(); err != nil {
		return nil, err
	}

	cond, err := parseExpression(s)
	if err != nil {
		return nil, err
	}

	if s.Token.Type != lexer.LeftBrace {
		return nil, fmt.Errorf("Unexpected token %s", s.Token)
	}
	if err := s.ReadNext(); err != nil {
		return nil, err
	}

	v, err := parseExpression(s)
	if err != nil {
		return nil, err
	}

	if s.Token.Type != lexer.RightBrace {
		return nil, fmt.Errorf("Unexpected token %s", s.Token)
	}
	if err := s.ReadNext(); err != nil {
		return nil, err
	}

	n := &ast.IfExpression{Condition: cond, Value: v}

	if s.Token.Type != lexer.ElseKeyword {
		return nil, fmt.Errorf("Unexpected token %s", s.Token)
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
			return nil, fmt.Errorf("Unexpected token %s", s.Token)
		}
		if err := s.ReadNext(); err != nil {
			return nil, err
		}

		n.ElseBranch = &ast.IfExpression{Value: v}
	} else {
		return nil, fmt.Errorf("Unexpected token %s", s.Token)
	}

	return n, nil
}
