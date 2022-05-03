package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	. "github.com/JackMenandCameron/goldilox/internal"
)

var hadError bool = false

func main() {

	t := Token{Ttype: EOF, Lexeme: "hey"}
	fmt.Println(t)
	var err error = nil
	if len(os.Args) > 2 {
		fmt.Println("Usage: goldilox [script]")
		os.Exit(64)
	} else if len(os.Args) == 2 {
		err = RunFile(os.Args[1])
	} else {
		err = runPrompt()
	}
	if err != nil {
		fmt.Println(err.Error())
	}
}

func RunFile(filename string) error {
	fmt.Println("Running", filename)
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	err = run(string(data))
	// TODO look into this because two different errors meaning
	// different things hurts my head
	if hadError {
		os.Exit(65)
	}
	return err
}

func runPrompt() error {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">")
		line, _ := reader.ReadString('\n')
		line = strings.Replace(line, "\n", "", -1)
		if strings.Compare("", line) == 0 {
			return nil
		}
		run(line)
		hadError = false

	}
}

func run(source string) error {
	// Scanner scanner = new Scanner(source);
	// List<Token> tokens = scanner.scanTokens();

	// // For now, just print the tokens.
	// for (Token token : tokens) {
	// System.out.println(token);
	// }
	fmt.Println(source)
	return nil
}

func compError(line int, message string) {
	report(line, "", message)
}

func report(line int, where string, message string) {
	fmt.Printf("[line %d] Error %s: %s", line, where, message)
	hadError = true
}
