package internal

import (
	"errors"
	"fmt"
)

type Reporter struct {
	CompilationErrors []string
}

func NewReporter() *Reporter {
	return &Reporter{}
}

func (r *Reporter) compError(line int, message string) {
	r.addError(line, "", message)
}

func (r *Reporter) addError(line int, where string, message string) {
	r.CompilationErrors = append(r.CompilationErrors,
		fmt.Sprintf("[line %d] Error %s: %s", line, where, message))
}

func (r *Reporter) Report() {
	for _, err := range r.CompilationErrors {
		fmt.Println(err)
	}
}

func (r *Reporter) HasError() error {
	if len(r.CompilationErrors) > 0 {
		return errors.New(ErrorMessages[COMPILATION_ERROR])
	}
	return nil
}

type ERRMSG int

const (
	// Scanner
	NEW_SCANNER_NO_SOURCE ERRMSG = iota
	SCAN_TOKEN_NO_SOURCE

	// Compilation
	UNEXPECTED_CHARACTER
	UNTERMINATED_STRING

	// Main
	COMPILATION_ERROR
)

var ErrorMessages = map[ERRMSG]string{
	// Scanner
	NEW_SCANNER_NO_SOURCE: "Scanner: NewScanner called without source",
	SCAN_TOKEN_NO_SOURCE:  "Scanner: ScanTokens called without source, did you not use NewScanner?",

	// Compilation
	UNEXPECTED_CHARACTER: "Unexpected Character",
	UNTERMINATED_STRING:  "Unterminated string",

	// Main
	COMPILATION_ERROR: "Compilation error detected",
}
