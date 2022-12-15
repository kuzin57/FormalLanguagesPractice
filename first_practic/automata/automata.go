package automata

import (
	"strings"

	"github.com/kuzin57/FormalPractic/first_practic/parser"
)

type Automata interface {
	ReadMaxPrefix(word string) int
	Flush()
	cycle()
	join(Automata)
	concat(Automata)
	getStartNode() *node
	getTerminals() []*node
	getNodes(*node, map[*node]bool)
}

type automata struct {
	startNode *node
	terminals []*node
}

func (a *automata) Flush() {
	used := make(map[*node]bool)
	a.getNodes(a.startNode, used)
	for node := range used {
		for _, edges := range node.next {
			for _, edge := range edges {
				edge.prefixRead = 0
				edge.visited = false
			}
		}
	}
}

func (a *automata) getNodes(curNode *node, used map[*node]bool) {
	used[curNode] = true

	for _, nodesTo := range curNode.next {
		for _, nodeTo := range nodesTo {
			_, ok := used[nodeTo.to]
			if !ok {
				a.getNodes(nodeTo.to, used)
			}
		}
	}
}

func CreateAutomata(parser *parser.Parser) Automata {
	if parser == nil {
		return nil
	}

	var (
		leftAutomata  Automata
		rightAutomata Automata
	)

	parser.Parse()

	if !strings.Contains(parser.Token, "+") &&
		!strings.Contains(parser.Token, ".") &&
		!strings.Contains(parser.Token, "*") {
		newAutomata := &automata{terminals: make([]*node, 0)}
		newStartNode := newNode()

		addFromNode, maxPos, _ := newStartNode.findMaxPrefix(parser.Token, 0, 0)
		newTerminal := addFromNode.createBranch(parser.Token[maxPos:], maxPos)

		newAutomata.startNode = newStartNode
		newAutomata.terminals = append(newAutomata.terminals, newTerminal)
		return newAutomata
	}

	leftAutomata = CreateAutomata(parser.ChildLeft)
	rightAutomata = CreateAutomata(parser.ChildRight)

	switch parser.Operation {
	case '+':
		leftAutomata.join(rightAutomata)
	case '.':
		leftAutomata.concat(rightAutomata)
	case '*':
		leftAutomata.cycle()
	}
	nodes := make([]*node, 0)
	used := make(map[*node]bool)
	leftAutomata.getNodes(leftAutomata.getStartNode(), used)
	for key := range used {
		nodes = append(nodes, key)
	}
	return leftAutomata
}

func (a *automata) ReadMaxPrefix(word string) int {
	_, _, maxPrefixRead := a.startNode.findMaxPrefix(word, 0, 0)
	return maxPrefixRead
}

func (a *automata) cycle() {
	newStartNode := newNode()
	newStartNode.next[empty] = []*edge{&edge{to: a.startNode}}
	newStartNode.isTerminal = true

	for _, terminal := range a.terminals {
		terminal.isTerminal = false
		terminal.next[empty] = []*edge{&edge{to: newStartNode}}
	}

	a.startNode = newStartNode
	a.terminals = []*node{a.startNode}
}

func (a *automata) join(other Automata) {
	a.startNode.next[empty] = append(a.startNode.next[empty], &edge{to: other.getStartNode()})
	a.terminals = append(a.terminals, other.getTerminals()...)
}

func (a *automata) getStartNode() *node {
	return a.startNode
}

func (a *automata) getTerminals() []*node {
	return a.terminals
}

func (a *automata) concat(other Automata) {
	otherStartNode := other.getStartNode()

	for _, terminal := range a.terminals {
		terminal.isTerminal = false
		terminal.next[empty] = append(terminal.next[empty], &edge{to: otherStartNode})
	}
	a.terminals = other.getTerminals()
}
