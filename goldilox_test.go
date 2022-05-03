package main

import (
	"os"
	"testing"
	// "fmt"

	. "github.com/JackMenandCameron/goldilox/internal"
)

func TestScannerBasic(t *testing.T) {
	s := NewScanner("and")
	s.ScanTokens()
	// Remember the EOF as well
	if len(s.Tokens) != 2 || s.Tokens[0].Ttype != AND {
		t.Errorf("Expected AND, EOF got %s", s.Tokens)
	}

}

func TestScannerReal(t *testing.T) {
	t.Run("Testing line", func(t *testing.T) {
		s := NewScanner("print \"hello lox\"")
		s.ScanTokens()
		if len(s.Tokens) != 3 || s.Tokens[0].Ttype != PRINT || s.Tokens[1].Ttype != STRING {
			t.Errorf("Expected PRINT, STRING, and EOF got %s", s.Tokens)
		}
		if s.Tokens[1].Value != "hello lox" {
			t.Errorf("Expected \"hello lox\" got %s", s.Tokens[1].Value)
		}
	})

	t.Run("Testing line from file", func(t *testing.T) {
		data, err := os.ReadFile("test_lox/hello.lox")
		if err != nil {
			t.Errorf(err.Error())
		}
		s := NewScanner(string(data))
		s.ScanTokens()
		if len(s.Tokens) != 3 || s.Tokens[0].Ttype != PRINT || s.Tokens[1].Ttype != STRING {
			t.Errorf("Expected PRINT, STRING, and EOF got %s", s.Tokens)
		}
		if s.Tokens[1].Value != "hello lox" {
			t.Errorf("Expected \"hello lox\" got %s", s.Tokens[1].Value)
		}

	})

	t.Run("Testing conditionals", func(t *testing.T) {
		data, err := os.ReadFile("test_lox/conditional.lox")
		if err != nil {
			t.Errorf(err.Error())
		}
		s := NewScanner(string(data))
		s.ScanTokens()
		expected := []TokenType{VAR, IDENTIFIER, EQUAL, FALSE, IF, LEFT_PAREN,
			IDENTIFIER, RIGHT_PAREN, LEFT_BRACE, PRINT, STRING, SEMICOLON,
			RIGHT_BRACE, ELSE, LEFT_BRACE, PRINT, STRING, SEMICOLON,
			RIGHT_BRACE, EOF}
			if len(expected) != len(s.Tokens) {
				t.Errorf("Expected %d tokens got %d with %s", len(expected), len(s.Tokens), s.Tokens)
			}
		for i, token := range s.Tokens{
			if expected[i] != token.Ttype {
				t.Errorf("Expected %s, got %s", TTToString[expected[i]], TTToString[token.Ttype])
			}
		}

	})

	t.Run("Testing classes", func(t *testing.T) {
		data, err := os.ReadFile("test_lox/class.lox")
		if err != nil {
			t.Errorf(err.Error())
		}
		s := NewScanner(string(data))
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
		for i, token := range s.Tokens{
			if expected[i] != token.Ttype {
				t.Errorf("Expected %s, got %s", TTToString[expected[i]], TTToString[token.Ttype])
			}
		}

	})
}
