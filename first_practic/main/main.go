package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/kuzin57/FormalPractic/first_practic/parser"

	"github.com/kuzin57/FormalPractic/first_practic/automata"
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
