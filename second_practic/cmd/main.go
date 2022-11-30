package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/kuzin57/FormalPractic/second_practic/builder"
)

const (
	inputFile = "../grammar.txt"
)

func main() {
	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	var rules []string
	for scanner.Scan() {
		rules = append(rules, scanner.Text())
	}

	gramBuilder := builder.NewGrammarBuilder(rules)
	grammar := gramBuilder.Build()

	scanner = bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		word := scanner.Text()
		if grammar.CheckWord(word) {
			fmt.Println("YES")
			continue
		}
		fmt.Println("NO")
	}
}
