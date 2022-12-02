package adapter

import (
	"os"

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
	completeConfigurations  []configuration
	configurationInfos      []map[configuration]configurationInfo
	predict, complete, scan func(configuration, int, string) bool
	word                    string
	logger                  *logger
	currentConfNumber       int
}

type logger struct {
	file *os.File
}

type configurationInfo struct {
	method string
	number int
}
