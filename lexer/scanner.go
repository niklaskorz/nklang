package lexer

import (
	"bufio"
	"fmt"
	"io"
	"unicode"
)

const bufferSize = 32
const eofRune = -1

type Scanner struct {
	rd            *bufio.Reader
	line, column  int
	eof           bool
	Token         *Token
	previousToken *Token
	nextToken     *Token
}

func NewScanner(rd io.Reader) *Scanner {
	return &Scanner{rd: bufio.NewReader(rd), line: 1, column: 0, eof: false}
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
		s.Token = &Token{Line: s.line, Column: s.column, Type: EOF, Value: "EOF"}
		return nil
	}
	return err
}

func (s *Scanner) readRune() (rune, error) {
	if s.eof {
		return eofRune, nil
	}

	r, _, err := s.rd.ReadRune()
	if err != nil {
		if err != io.EOF {
			return 0, err
		}
		r = eofRune
		s.eof = true
	}

	s.column++
	return r, nil
}

func (s *Scanner) unreadRune() error {
	if s.eof {
		return nil
	}
	if s.column == 0 {
		return fmt.Errorf("Cannot unread rune after line break")
	}
	s.column--
	return s.rd.UnreadRune()
}

func (s *Scanner) readNext() error {
	var r rune
	var err error

	for {
		// Skip whitespace
		r, err = s.readRune()
		if err != nil {
			return err
		}

		if r == eofRune {
			return io.EOF
		}

		if !unicode.IsSpace(r) {
			break
		}

		if r == '\n' || r == '\r' {
			s.line++
			s.column = 0
		}

		if r == '\r' {
			// Skip LF after CR
			r, err = s.readRune()
			if err != nil {
				return err
			}
			if r != '\n' {
				// No LF, rewind
				if err := s.unreadRune(); err != nil {
					return err
				}
			}
		}
	}

	line, column := s.line, s.column

	if r == '"' {
		// String
		v, err := s.scanString()
		if err != nil {
			return err
		}
		s.Token = &Token{Line: line, Column: column, Type: String, Value: v}
		return nil
	}

	if unicode.IsLetter(r) || r == '_' {
		// Identifier
		v, err := s.scanIdentifier()
		if err != nil {
			return err
		}
		id := string(r) + v
		s.Token = &Token{Line: line, Column: column, Type: ID, Value: id}

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
		case "nil":
			s.Token.Type = NilKeyword
		}

		return nil
	}

	if unicode.IsDigit(r) {
		// Number
		v, err := s.scanIdentifier()
		if err != nil {
			return err
		}
		num := string(r) + v
		s.Token = &Token{Line: line, Column: column, Type: Integer, Value: num}
		return nil
	}

	if r == ':' {
		r, err = s.readRune()
		if err != nil {
			return err
		}
		if r != '=' {
			return s.unexpectedSymbol(r)
		}
		s.Token = &Token{Line: line, Column: column, Type: DeclarationOperator, Value: ":="}
		return nil
	}

	if r == '=' {
		r, err := s.readRune()
		if err != nil {
			return err
		}

		if r == '=' {
			s.Token = &Token{Line: line, Column: column, Type: EqOperator, Value: "=="}
		} else {
			if err := s.unreadRune(); err != nil {
				return err
			}
			s.Token = &Token{Line: line, Column: column, Type: AssignmentOperator, Value: "="}
		}
		return nil
	}

	if r == '!' {
		r, err := s.readRune()
		if err != nil {
			return err
		}

		if r == '=' {
			s.Token = &Token{Line: line, Column: column, Type: NeOperator, Value: "!="}
		} else {
			if err := s.unreadRune(); err != nil {
				return err
			}
			s.Token = &Token{Line: line, Column: column, Type: LogicalNot, Value: "!"}
		}
		return nil
	}

	if r == '<' {
		r, err := s.readRune()
		if err != nil {
			return err
		}

		if r == '=' {
			s.Token = &Token{Line: line, Column: column, Type: LeOperator, Value: "<="}
		} else {
			if err := s.unreadRune(); err != nil {
				return err
			}
			s.Token = &Token{Line: line, Column: column, Type: LtOperator, Value: "<"}
		}
		return nil
	}

	if r == '>' {
		r, err := s.readRune()
		if err != nil {
			return err
		}

		if r == '=' {
			s.Token = &Token{Line: line, Column: column, Type: GeOperator, Value: ">="}
		} else {
			if err := s.unreadRune(); err != nil {
				return err
			}
			s.Token = &Token{Line: line, Column: column, Type: GtOperator, Value: ">"}
		}
		return nil
	}

	if r == '&' {
		r, err = s.readRune()
		if err != nil {
			return err
		}
		if r != '&' {
			return s.unexpectedSymbol(r)
		}
		s.Token = &Token{Line: line, Column: column, Type: LogicalAnd, Value: "&&"}
		return nil
	}

	if r == '|' {
		r, err = s.readRune()
		if err != nil {
			return err
		}
		if r != '|' {
			return s.unexpectedSymbol(r)
		}
		s.Token = &Token{Line: line, Column: column, Type: LogicalOr, Value: "||"}
		return nil
	}

	if r == '*' {
		s.Token = &Token{Line: line, Column: column, Type: MulOperator, Value: "*"}
		return nil
	}

	if r == '/' {
		s.Token = &Token{Line: line, Column: column, Type: DivOperator, Value: "/"}
		return nil
	}

	if r == '+' {
		s.Token = &Token{Line: line, Column: column, Type: Plus, Value: "+"}
		return nil
	}

	if r == '-' {
		s.Token = &Token{Line: line, Column: column, Type: Minus, Value: "-"}
		return nil
	}

	if r == ';' {
		s.Token = &Token{Line: line, Column: column, Type: Semicolon, Value: ";"}
		return nil
	}

	if r == ',' {
		s.Token = &Token{Line: line, Column: column, Type: Comma, Value: ","}
		return nil
	}

	if r == '(' {
		s.Token = &Token{Line: line, Column: column, Type: LeftParen, Value: "("}
		return nil
	}

	if r == ')' {
		s.Token = &Token{Line: line, Column: column, Type: RightParen, Value: ")"}
		return nil
	}

	if r == '{' {
		s.Token = &Token{Line: line, Column: column, Type: LeftBrace, Value: "{"}
		return nil
	}

	if r == '}' {
		s.Token = &Token{Line: line, Column: column, Type: RightBrace, Value: "}"}
		return nil
	}

	if r == '[' {
		s.Token = &Token{Line: line, Column: column, Type: LeftBracket, Value: "["}
		return nil
	}

	if r == ']' {
		s.Token = &Token{Line: line, Column: column, Type: RightBracket, Value: "]"}
		return nil
	}

	return s.unexpectedSymbol(r)
}

func (s *Scanner) scanString() (string, error) {
	str := ""
	for {
		r, _, err := s.rd.ReadRune()
		if err != nil {
			return "", err
		}
		s.column++

		if r == '\n' {
			s.line++
			s.column = 0
		}

		if r == '"' {
			return str, nil
		}

		str += string(r)
	}
}

func (s *Scanner) scanIdentifier() (string, error) {
	str := ""
	for {
		r, err := s.readRune()
		if err != nil {
			return "", err
		}

		if !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '_' {
			if err := s.unreadRune(); err != nil {
				return "", err
			}
			return str, nil
		}

		str += string(r)
	}
}

func (s *Scanner) scanInteger() (string, error) {
	str := ""
	for {
		r, err := s.readRune()
		if err != nil {
			return "", err
		}

		if !unicode.IsDigit(r) {
			if err := s.unreadRune(); err != nil {
				return "", err
			}
			return str, nil
		}

		str += string(r)
	}
}
