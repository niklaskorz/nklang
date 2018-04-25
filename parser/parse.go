package parser

import (
	"fmt"
	"strconv"

	"niklaskorz.de/nklang/lexer"
)

func Parse(s *lexer.Scanner) (BlockNode, error) {
	if err := s.ReadNext(); err != nil {
		return nil, err
	}

	statements := []*Node{}
	for s.Token.Type != lexer.EOF {
		n, err := parseStatement(s)
		if err != nil {
			return nil, err
		}

		statements = append(statements, n)
	}

	n := BlockNode{statements: statements}
	return n, nil
}

func parseStatement(s *lexer.Scanner) (Node, error) {
	n, err := parseExpression(s)
	if err != nil {
		return nil, err
	}
	if s.Token.Type != lexer.Semicolon {
		return nil, fmt.Errorf("Unexpected token %s", s.Token.Value)
	}
	if err := s.ReadNext(); err != nil {
		return nil, err
	}
	return n, nil
}

func parseExpression(s *lexer.Scanner) (Node, error) {
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
			return nil, fmt.Errorf("Unexpected token %s", s.Token.Value)
		}
		if err := s.ReadNext(); err != nil {
			return nil, err
		}
		return n, nil
	case lexer.Identifier:
		identifier := s.Token.Value
		if err := s.ReadNext(); err != nil {
			return nil, err
		}
		switch s.Token.Type {
		case lexer.DeclarationOperator:
			if err := s.ReadNext(); err != nil {
				return nil, err
			}
			n := DeclarationNode{Identifier: identifier}
			v, err := parseExpression(s)
			if err != nil {
				return nil, err
			}
			n.Value = v
			return n, nil
		case lexer.AssignmentOperator:
			if err := s.ReadNext(); err != nil {
				return nil, err
			}
			n := AssignmentNode{Identifier: identifier}
			v, err := parseExpression(s)
			if err != nil {
				return nil, err
			}
			n.Value = v
			return n, nil
		default:
			n := LookupNode{Identifier: identifier}
			return n, nil
		}
	case lexer.Integer:
		num, err := strconv.ParseInt(s.Token.Value, 10, 64)
		if err != nil {
			return nil, err
		}
		n := IntegerNode{Value: num}
		if err := s.ReadNext(); err != nil {
			return nil, err
		}
		return n, nil
	case lexer.String:
		n := StringNode{Value: s.Token.Value}
		if err := s.ReadNext(); err != nil {
			return nil, err
		}
		return n, nil
	}

	return nil, fmt.Errorf("Unexpected token %s", s.Token.Value)
}
