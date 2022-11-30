package adapter

import (
	"github.com/kuzin57/FormalPractic/second_practic/grammar"
)

const (
	startSymbol = '$'
)

type configuration struct {
	expression string
	position   int
	startIndex int
	terminal   byte
}

type grammarAdapter struct {
	grammar                 *grammar.Grammar
	currentConfigurations   []map[configuration]struct{}
	predict, complete, scan func(configuration, int, string) bool
	word                    string
}
