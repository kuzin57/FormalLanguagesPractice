package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/kuzin57/FormalPractic/third_practic/adapter"
)

func main() {
	grammarAdapter, err := adapter.BuildAdapter(os.Args[1])
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		word := scanner.Text()
		word += string('$')
		if grammarAdapter.Read(word) {
			fmt.Println("YES")
			grammarAdapter.Flush()
			continue
		}
		fmt.Println("NO")
		grammarAdapter.Flush()
	}
}
