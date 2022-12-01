package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/kuzin57/FormalPractic/second_practic/adapter"
)

func main() {
	file, err := os.Open("../grammar.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	var rules []string
	for scanner.Scan() {
		rules = append(rules, scanner.Text())
	}

	grammarAdapter := adapter.NewGrammarAdapter()
	grammarAdapter.BuildGrammar(rules)

	scanner = bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		word := scanner.Text()
		if grammarAdapter.Read(word) {
			fmt.Println("YES")
			grammarAdapter.Flush()
			continue
		}
		fmt.Println("NO")
		grammarAdapter.Flush()
	}
}
