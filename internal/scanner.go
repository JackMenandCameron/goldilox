package internal

import (
	// "strconv"

	"errors"
)

var keywords = map[string]TokenType{
	"and":    AND,
	"class":  CLASS,
	"else":   ELSE,
	"false":  FALSE,
	"for":    FOR,
	"fun":    FUN,
	"if":     IF,
	"nil":    NIL,
	"or":     OR,
	"print":  PRINT,
	"return": RETURN,
	"super":  SUPER,
	"this":   THIS,
	"true":   TRUE,
	"var":    VAR,
	"while":  WHILE,
}

type Scanner struct {
	source  string
	Tokens  []Token
	Reporter *Reporter
	start   int
	current int
	line    int
}

func NewScanner(source string, reporter *Reporter) (*Scanner, error) {
	if source == "" {
		return nil, errors.New(ErrorMessages[NEW_SCANNER_NO_SOURCE])
	}
	return &Scanner{source: source, start: 0, Reporter: reporter, current: 0, line: 1}, nil
}

func (s *Scanner) ScanTokens() error {
	if s.source == "" {
		return errors.New(ErrorMessages[SCAN_TOKEN_NO_SOURCE])
	}

	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
	}

	s.Tokens = append(s.Tokens, MakeToken(EOF, s.line))
	return nil
}

func (s *Scanner) scanToken() {
	c := s.advance()

	switch c {
	case '(':
		s.addToken(LEFT_PAREN)
	case ')':
		s.addToken(RIGHT_PAREN)
	case '{':
		s.addToken(LEFT_BRACE)
	case '}':
		s.addToken(RIGHT_BRACE)
	case ',':
		s.addToken(COMMA)
	case '.':
		s.addToken(DOT)
	case '-':
		s.addToken(MINUS)
	case '+':
		s.addToken(PLUS)
	case ';':
		s.addToken(SEMICOLON)
	case '*':
		s.addToken(STAR)
	case '!':
		if s.match('=') {
			s.addToken(BANG_EQUAL)
		} else {
			s.addToken(BANG)
		}
	case '=':
		if s.match('=') {
			s.addToken(EQUAL_EQUAL)
		} else {
			s.addToken(EQUAL)
		}
	case '<':
		if s.match('=') {
			s.addToken(LESS_EQUAL)
		} else {
			s.addToken(LESS)
		}
	case '>':
		if s.match('=') {
			s.addToken(GREATER_EQUAL)
		} else {
			s.addToken(GREATER)
		}
	case '/':
		if s.match('/') {
			// A comment goes until the end of the line.
			for s.peek() != '\n' && !s.isAtEnd() {
				s.advance()
			}
		} else {
			s.addToken(SLASH)
		}
	case ' ': // Ignore whitespace
	case '\r':
	case '\t':
	case '\n':
		s.line++
	case '"':
		s.readString()
	default:
		if isDigit(c) {
			s.readNumber()
		} else if isAlpha(c) {
			s.readIdentifier()
		} else {
			s.Reporter.compError(s.line, ErrorMessages[UNEXPECTED_CHARACTER])
		}
	}
}

// Adding Tokens
//

func (s *Scanner) addToken(tt TokenType) {
	s.addTokenV(tt, "")
}

func (s *Scanner) addTokenV(tt TokenType, value string) {
	t := Token{Ttype: tt, Lexeme: s.source[s.start:s.current], Value: value, Line: s.line}
	s.Tokens = append(s.Tokens, t)
}

// Reading Complex lexeme
//

func (s *Scanner) readString() {
	for s.peek() != '"' && !s.isAtEnd() {
		if s.peek() == '\n' {
			s.line++
		}
		s.advance()
	}

	if (s.isAtEnd()) {
	  s.Reporter.compError(s.line, ErrorMessages[UNTERMINATED_STRING])
	  return;
	}

	// The closing ".
	s.advance()

	// Trim the surrounding quotes.
	value := s.source[s.start+1 : s.current-1]
	s.addTokenV(STRING, value)
}

func (s *Scanner) readNumber() {
	for isDigit(s.peek()) {
		s.advance()
	}

	// Look for a fractional part.
	if s.peek() == '.' && isDigit(s.peekNext()) {
		// Consume the "."
		s.advance()

		for isDigit(s.peek()) {
			s.advance()
		}
	}

	value := s.source[s.start:s.current]
	s.addTokenV(NUMBER, value)
}

func (s *Scanner) readIdentifier() {
	for isAlphaNumeric(s.peek()) {
		s.advance()
	}
	text := s.source[s.start:s.current]
	iType, ok := keywords[text]
	if !ok {
		iType = IDENTIFIER
	}
	s.addToken(iType)
}

// Helper Functions
//

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) advance() byte {
	c := s.source[s.current]
	s.current = s.current + 1
	return c
}

func (s *Scanner) match(expected byte) bool {
	if s.isAtEnd() {
		return false
	}
	if s.source[s.current] != expected {
		return false
	}

	s.current++
	return true
}

func (s *Scanner) peek() byte {
	if s.isAtEnd() {
		return 0
	}
	return s.source[s.current]
}

func (s *Scanner) peekNext() byte {
	if s.current+1 >= len(s.source) {
		return 0
	}
	return s.source[s.current+1]
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func isAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		c == '_'
}

func isAlphaNumeric(c byte) bool {
	return isAlpha(c) || isDigit(c)
}
