package adapter

import (
	"fmt"

	"github.com/kuzin57/FormalPractic/third_practic/grammar"
)

const (
	endSymbol     = byte('$')
	startTerminal = byte('%')
)

var (
	errStateExists = fmt.Errorf("state exists")
	errNotLR1      = fmt.Errorf("grammar can't be processed by LR(1) parser")
	errNoSuchRule  = fmt.Errorf("no such rule")
)

type operation uint8

const (
	reject operation = iota
	reduce
	shift
	accept
)

type action struct {
	operationType operation
	rule          int
	nonTerminal   byte
}

type grammarAdapter struct {
	grammar   *grammar.Grammar
	states    []*state
	actions   []map[byte]action
	terminals map[byte]struct{}
	stack     []int
	first     map[byte]map[byte]struct{}
}

type state struct {
	transitions map[byte]int
	situations  map[situation]struct{}
}

type situation struct {
	rightPart   string
	position    int
	nonTerminal byte
	promise     byte
}
