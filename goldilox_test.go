package main

import (
	"os"
	"testing"
	// "fmt"

	. "github.com/JackMenandCameron/goldilox/internal"
)

func TestScannerBasic(t *testing.T) {
	report := NewReporter()
	s, err := NewScanner("and", report)
	if err != nil {
		t.Errorf(err.Error())
	}
	s.ScanTokens()
	// Remember the EOF as well
	if len(s.Tokens) != 2 || s.Tokens[0].Ttype != AND {
		t.Errorf("Expected AND, EOF got %s", s.Tokens)
	}

}

func fileToScanner(filename string, t *testing.T) *Scanner {
	data, err := os.ReadFile(filename)
	if err != nil {
		t.Errorf(err.Error())
	}
	report := NewReporter()
	s, err := NewScanner(string(data), report)
	if err != nil {
		t.Errorf(err.Error())
	}
	return s
}

func TestScannerReal(t *testing.T) {

	t.Run("Testing line", func(t *testing.T) {
		report := NewReporter()
		s, err := NewScanner("print \"hello lox\"", report)
		if err != nil {
			t.Errorf(err.Error())
		}
		s.ScanTokens()
		if len(s.Tokens) != 3 || s.Tokens[0].Ttype != PRINT || s.Tokens[1].Ttype != STRING {
			t.Errorf("Expected PRINT, STRING, and EOF got %s", s.Tokens)
		}
		if s.Tokens[1].Value != "hello lox" {
			t.Errorf("Expected \"hello lox\" got %s", s.Tokens[1].Value)
		}
	})

	t.Run("Testing line from file", func(t *testing.T) {
		s := fileToScanner("test_lox/hello.lox", t)
		s.ScanTokens()
		if len(s.Tokens) != 3 || s.Tokens[0].Ttype != PRINT || s.Tokens[1].Ttype != STRING {
			t.Errorf("Expected PRINT, STRING, and EOF got %s", s.Tokens)
		}
		if s.Tokens[1].Value != "hello lox" {
			t.Errorf("Expected \"hello lox\" got %s", s.Tokens[1].Value)
		}

	})

	t.Run("Testing conditionals", func(t *testing.T) {
		s := fileToScanner("test_lox/conditional.lox", t)
		s.ScanTokens()
		expected := []TokenType{VAR, IDENTIFIER, EQUAL, FALSE, IF, LEFT_PAREN,
			IDENTIFIER, RIGHT_PAREN, LEFT_BRACE, PRINT, STRING, SEMICOLON,
			RIGHT_BRACE, ELSE, LEFT_BRACE, PRINT, STRING, SEMICOLON,
			RIGHT_BRACE, EOF}
		if len(expected) != len(s.Tokens) {
			t.Errorf("Expected %d tokens got %d with %s", len(expected), len(s.Tokens), s.Tokens)
		}
		for i, token := range s.Tokens {
			if expected[i] != token.Ttype {
				t.Errorf("Expected %s, got %s", TTToString[expected[i]], TTToString[token.Ttype])
			}
		}

	})

	t.Run("Testing classes", func(t *testing.T) {
		s := fileToScanner("test_lox/class.lox", t)
		s.ScanTokens()
		expected := []TokenType{CLASS, IDENTIFIER, LEFT_BRACE,
			IDENTIFIER, LEFT_PAREN, IDENTIFIER, COMMA, IDENTIFIER, RIGHT_PAREN, LEFT_BRACE,
			THIS, DOT, IDENTIFIER, EQUAL, IDENTIFIER, SEMICOLON,
			THIS, DOT, IDENTIFIER, EQUAL, IDENTIFIER, SEMICOLON,
			RIGHT_BRACE,
			RIGHT_BRACE,
			VAR, IDENTIFIER, EQUAL, IDENTIFIER, LEFT_PAREN, STRING, COMMA, STRING, RIGHT_PAREN, SEMICOLON,
			IDENTIFIER, DOT, IDENTIFIER, LEFT_PAREN, STRING, RIGHT_PAREN, SEMICOLON, EOF,
		}
		if len(expected) != len(s.Tokens) {
			t.Errorf("Expected %d tokens got %d with %s", len(expected), len(s.Tokens), s.Tokens)
		}
		for i, token := range s.Tokens {
			if expected[i] != token.Ttype {
				t.Errorf("Expected %s, got %s", TTToString[expected[i]], TTToString[token.Ttype])
			}
		}

	})

	t.Run("Testing Reporter", func(t *testing.T) {
		s := fileToScanner("test_lox/unterminated.lox", t)
		s.ScanTokens()
		if len(s.Reporter.CompilationErrors) == 0 {
			t.Errorf("Compilation error not reported")
		}

	})
}
