package main

import (
	"bufio"
	"fmt"
	. "github.com/JackMenandCameron/goldilox/internal"
	"os"
	"strings"
)

func main() {
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
		if err.Error() != ErrorMessages[COMPILATION_ERROR] {
			fmt.Println(err.Error())
		}
		os.Exit(65)
	}
}

func RunFile(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	err = run(string(data))
	return err
}

func runPrompt() error {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">")
		line, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		line = strings.Replace(line, "\n", "", -1)

		// End on a blank line
		if strings.Compare("", line) == 0 {
			return nil
		}

		if err = run(line); err != nil {
			fmt.Println(err.Error())
		}
	}
}

func run(source string) error {
	reporter := NewReporter()
	s, err := NewScanner(source, reporter)
	if err != nil {
		return err
	}
	s.ScanTokens()
	reporter.Report()
	return reporter.HasError()
}
