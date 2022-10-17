package main

import (
	"bufio"
	"fmt"
	"formal/first_practic/automata"
	"formal/first_practic/parser"
	"os"
)

func main() {
	var (
		expression string
		scanner    = bufio.NewScanner(os.Stdin)
	)

	scanner.Scan()
	expression = scanner.Text()
	parser := parser.NewParser(expression, nil)
	automata := automata.CreateAutomata(parser)

	for scanner.Scan() {
		nextWord := scanner.Text()
		result := automata.ReadMaxPrefix(nextWord)
		automata.Flush()
		fmt.Println(result)
	}
}
