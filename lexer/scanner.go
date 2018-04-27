package lexer

import (
	"bufio"
	"fmt"
	"io"
	"unicode"
)

const bufferSize = 32

type Scanner struct {
	rd            *bufio.Reader
	Token         *Token
	previousToken *Token
	nextToken     *Token
}

func NewScanner(rd io.Reader) *Scanner {
	return &Scanner{rd: bufio.NewReader(rd)}
}

func (s *Scanner) Unread() error {
	if s.previousToken == nil {
		return fmt.Errorf("Can't unread more than one token")
	}
	s.nextToken = s.Token
	s.Token = s.previousToken
	s.previousToken = nil
	return nil
}

func (s *Scanner) ReadNext() error {
	s.previousToken = s.Token

	if s.nextToken != nil {
		s.Token = s.nextToken
		s.nextToken = nil
		return nil
	}

	err := s.readNext()
	if err == io.EOF {
		s.Token = &Token{Type: EOF, Value: "EOF"}
		return nil
	}
	return err
}

func (s *Scanner) readNext() error {
	var r rune
	var err error

	for {
		// Skip whitespace

		r, _, err = s.rd.ReadRune()
		if err != nil {
			return err
		}

		if !unicode.IsSpace(r) {
			break
		}
	}

	if r == '"' {
		// String
		v, err := scanString(s.rd)
		if err != nil {
			return err
		}
		s.Token = &Token{Type: String, Value: v}
		return nil
	}

	if unicode.IsLetter(r) || r == '_' {
		// Identifier
		v, err := scanIdentifier(s.rd)
		if err != nil {
			return err
		}
		id := string(r) + v
		s.Token = &Token{Type: ID, Value: id}

		switch id {
		case "return":
			s.Token.Type = ReturnKeyword
		case "continue":
			s.Token.Type = ContinueKeyword
		case "break":
			s.Token.Type = BreakKeyword
		case "func":
			s.Token.Type = FunctionKeyword
		case "if":
			s.Token.Type = IfKeyword
		case "else":
			s.Token.Type = ElseKeyword
		case "while":
			s.Token.Type = WhileKeyword
		case "true":
			s.Token.Type = TrueKeyword
		case "false":
			s.Token.Type = FalseKeyword
		}

		return nil
	}

	if unicode.IsDigit(r) {
		// Number
		v, err := scanIdentifier(s.rd)
		if err != nil {
			return err
		}
		num := string(r) + v
		s.Token = &Token{Type: Integer, Value: num}
		return nil
	}

	if r == ':' {
		r, _, err = s.rd.ReadRune()
		if err != nil {
			return err
		}
		if r != '=' {
			return fmt.Errorf("Unexpected symbol %c", r)
		}
		s.Token = &Token{Type: DeclarationOperator, Value: ":="}
		return nil
	}

	if r == '=' {
		r, _, err := s.rd.ReadRune()
		if err != nil {
			return err
		}

		if r == '=' {
			s.Token = &Token{Type: EqOperator, Value: "=="}
		} else {
			if err := s.rd.UnreadRune(); err != nil {
				return err
			}
			s.Token = &Token{Type: AssignmentOperator, Value: "="}
		}
		return nil
	}

	if r == '<' {
		r, _, err := s.rd.ReadRune()
		if err != nil {
			return err
		}

		if r == '=' {
			s.Token = &Token{Type: LeOperator, Value: "<="}
		} else {
			if err := s.rd.UnreadRune(); err != nil {
				return err
			}
			s.Token = &Token{Type: LtOperator, Value: "<"}
		}
		return nil
	}

	if r == '>' {
		r, _, err := s.rd.ReadRune()
		if err != nil {
			return err
		}

		if r == '=' {
			s.Token = &Token{Type: GeOperator, Value: ">="}
		} else {
			if err := s.rd.UnreadRune(); err != nil {
				return err
			}
			s.Token = &Token{Type: GtOperator, Value: ">"}
		}
		return nil
	}

	if r == '&' {
		r, _, err = s.rd.ReadRune()
		if err != nil {
			return err
		}
		if r != '&' {
			return fmt.Errorf("Unexpected symbol %c", r)
		}
		s.Token = &Token{Type: LogicalAnd, Value: "&&"}
		return nil
	}

	if r == '|' {
		r, _, err = s.rd.ReadRune()
		if err != nil {
			return err
		}
		if r != '|' {
			return fmt.Errorf("Unexpected symbol %c", r)
		}
		s.Token = &Token{Type: LogicalOr, Value: "||"}
		return nil
	}

	if r == '*' {
		s.Token = &Token{Type: MulOperator, Value: "*"}
		return nil
	}

	if r == '/' {
		s.Token = &Token{Type: DivOperator, Value: "/"}
		return nil
	}

	if r == '+' {
		s.Token = &Token{Type: AddOperator, Value: "+"}
		return nil
	}

	if r == '-' {
		s.Token = &Token{Type: SubOperator, Value: "-"}
		return nil
	}

	if r == ';' {
		s.Token = &Token{Type: Semicolon, Value: ";"}
		return nil
	}

	if r == ',' {
		s.Token = &Token{Type: Comma, Value: ","}
		return nil
	}

	if r == '(' {
		s.Token = &Token{Type: LeftParen, Value: "("}
		return nil
	}

	if r == ')' {
		s.Token = &Token{Type: RightParen, Value: ")"}
		return nil
	}

	if r == '{' {
		s.Token = &Token{Type: LeftBrace, Value: "{"}
		return nil
	}

	if r == '}' {
		s.Token = &Token{Type: RightBrace, Value: "}"}
		return nil
	}

	if r == '[' {
		s.Token = &Token{Type: LeftBracket, Value: "["}
		return nil
	}

	if r == ']' {
		s.Token = &Token{Type: RightBracket, Value: "]"}
		return nil
	}

	return fmt.Errorf("Unexpected symbol %c", r)
}

func scanString(rd *bufio.Reader) (string, error) {
	s := ""
	for {
		r, _, err := rd.ReadRune()
		if err != nil {
			return "", err
		}

		if r == '"' {
			return s, nil
		}

		s += string(r)
	}
}

func scanIdentifier(rd *bufio.Reader) (string, error) {
	s := ""
	for {
		r, _, err := rd.ReadRune()
		if err != nil {
			return "", err
		}

		if !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '_' {
			if err := rd.UnreadRune(); err != nil {
				return "", err
			}
			return s, nil
		}

		s += string(r)
	}
}

func scanInteger(rd *bufio.Reader) (string, error) {
	s := ""
	for {
		r, _, err := rd.ReadRune()
		if err != nil {
			return "", err
		}

		if !unicode.IsDigit(r) {
			if err := rd.UnreadRune(); err != nil {
				return "", err
			}
			return s, nil
		}

		s += string(r)
	}
}
