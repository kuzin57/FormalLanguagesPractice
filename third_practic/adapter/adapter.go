package adapter

import (
	"bufio"
	"os"

	"github.com/kuzin57/FormalPractic/third_practic/builder"
)

func newSituation(rightPart string, nonTerminal byte, position int, promise byte) situation {
	return situation{
		rightPart:   rightPart,
		nonTerminal: nonTerminal,
		position:    position,
		promise:     promise,
	}
}

func BuildAdapter(filename string) (GrammarAdapter, error) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	var rules []string
	for scanner.Scan() {
		rules = append(rules, scanner.Text())
	}

	grammarAdapter := &grammarAdapter{}
	grammarAdapter.BuildGrammar(rules)

	grammarAdapter.states = append(grammarAdapter.states,
		&state{transitions: make(map[byte]int), situations: make(map[situation]struct{})})

	grammarAdapter.addNewStartTerminal()
	grammarAdapter.states[0].situations[newSituation("S", startTerminal, 0, endSymbol)] = struct{}{}
	grammarAdapter.fillFirst()
	grammarAdapter.initState(0)

	err = grammarAdapter.buildTable()
	if err != nil {
		return nil, err
	}

	grammarAdapter.stack = append(grammarAdapter.stack, 0)
	return grammarAdapter, nil
}

func (ga *grammarAdapter) BuildGrammar(input []string) {
	builder := builder.NewGrammarBuilder(input)
	ga.grammar = builder.Build()

	ga.terminals = make(map[byte]struct{})
	for _, rules := range ga.grammar.Rules {
		for _, rule := range rules {
			for i := 0; i < len(rule); i++ {
				if ga.isNonTerminal(rule[i]) {
					continue
				}
				ga.terminals[rule[i]] = struct{}{}
			}
		}
	}
	ga.terminals[endSymbol] = struct{}{}
}

func (ga *grammarAdapter) Read(word string) bool {
	stateNumber := ga.stack[len(ga.stack)-1]
	firstSymbol := endSymbol
	if len(word) > 0 {
		firstSymbol = word[0]
	}

	action := ga.actions[stateNumber][firstSymbol]
	switch action.operationType {
	case reject:
		return false
	case shift:
		ga.stack = append(ga.stack, ga.states[stateNumber].transitions[firstSymbol])
		return ga.Read(word[1:])
	case reduce:
		ga.stack = ga.stack[:len(ga.stack)-len(ga.grammar.Rules[action.nonTerminal][action.rule])]
		ga.stack = append(ga.stack, ga.states[ga.stack[len(ga.stack)-1]].transitions[action.nonTerminal])
		return ga.Read(word)
	case accept:
		return true
	}

	return false
}

func (ga *grammarAdapter) addNewStartTerminal() {
	ga.grammar.Rules[startTerminal] = append(ga.grammar.Rules[startTerminal], "S")
}

func (ga *grammarAdapter) Flush() {
	ga.stack = []int{0}
}

func (ga *grammarAdapter) initState(index int) (int, error) {
	for curSituation := range ga.states[index].situations {
		if curSituation.position < len(curSituation.rightPart) {
			terminal := curSituation.rightPart[curSituation.position]
			if !ga.isNonTerminal(terminal) {
				continue
			}
			for _, right := range ga.grammar.Rules[terminal] {
				promise := ga.getFirstAfterExpression(curSituation.rightPart, curSituation.position)
				_, endSymbolExists := promise[endSymbol]
				if endSymbolExists && curSituation.promise != endSymbol {
					delete(promise, endSymbol)
					promise[curSituation.promise] = struct{}{}
				}
				for promiseSymbol := range promise {
					ga.states[index].situations[newSituation(
						right, terminal, 0, promiseSymbol)] = struct{}{}
				}
			}
		}
	}

	for i, state := range ga.states {
		if i == index {
			continue
		}
		if checkEqualStates(state, ga.states[index]) {
			return i, errStateExists
		}
	}

	situationsInNewStates := make(map[byte][]situation)
	for situation := range ga.states[index].situations {
		if situation.position == len(situation.rightPart) {
			continue
		}
		situation.position++
		situationsInNewStates[situation.rightPart[situation.position-1]] =
			append(situationsInNewStates[situation.rightPart[situation.position-1]], situation)
	}

	for symbol := range situationsInNewStates {
		newState := &state{
			situations:  make(map[situation]struct{}),
			transitions: make(map[byte]int),
		}
		for _, situation := range situationsInNewStates[symbol] {
			newState.situations[situation] = struct{}{}
		}
		ga.states = append(ga.states, newState)
		ind := len(ga.states) - 1
		toState, err := ga.initState(len(ga.states) - 1)
		if err != nil {
			ga.states[index].transitions[symbol] = toState
			ga.states = ga.states[:len(ga.states)-1]
			continue
		}
		ga.states[index].transitions[symbol] = ind
	}

	return index, nil
}

func (ga *grammarAdapter) hasEpsilonTransition(nonTerminal byte) bool {
	for _, rule := range ga.grammar.Rules[nonTerminal] {
		if len(rule) == 0 {
			return true
		}
	}
	return false
}

func (ga *grammarAdapter) getFirstAfterExpression(rightPart string, position int) map[byte]struct{} {
	ans := make(map[byte]struct{})
	for i := position + 1; i < len(rightPart); i++ {
		if !ga.isNonTerminal(rightPart[i]) {
			ans[rightPart[i]] = struct{}{}
			return ans
		}
		for first := range ga.first[rightPart[i]] {
			ans[first] = struct{}{}
		}
		if ga.isNonTerminal(rightPart[i]) && !ga.hasEpsilonTransition(rightPart[i]) {
			return ans
		}
	}
	ans[endSymbol] = struct{}{}
	return ans
}

func (ga *grammarAdapter) fillFirst() {
	ga.first = make(map[byte]map[byte]struct{})
	for nonTerminal, rules := range ga.grammar.Rules {
		ga.first[nonTerminal] = make(map[byte]struct{})
		for _, rule := range rules {
			if len(rule) > 0 && !ga.isNonTerminal(rule[0]) {
				ga.first[nonTerminal][rule[0]] = struct{}{}
			}
		}
	}

	somethingChanged := true
	for somethingChanged {
		somethingChanged = false
		for nonTerminal, rules := range ga.grammar.Rules {
			for _, rule := range rules {
				if len(rule) > 0 && rule[0] == nonTerminal {
					continue
				}
				if len(rule) > 0 && ga.isNonTerminal(rule[0]) {
					for terminal := range ga.first[rule[0]] {
						_, ok := ga.first[nonTerminal][terminal]
						if !ok {
							somethingChanged = true
							ga.first[nonTerminal][terminal] = struct{}{}
						}
					}
				}
			}
		}
	}
}

func (ga *grammarAdapter) isNonTerminal(symbol byte) bool {
	_, isTerm := ga.grammar.Rules[symbol]
	return isTerm
}

func checkEqualStates(first, second *state) bool {
	for situation := range first.situations {
		_, ok := second.situations[situation]
		if !ok {
			return false
		}
	}
	return true
}

func (ga *grammarAdapter) buildTable() error {
	for _, state := range ga.states {
		newTableLine := make(map[byte]action)
		for terminal := range ga.terminals {
			newTableLine[terminal] = action{
				rule:          -1,
				operationType: reject,
				nonTerminal:   startTerminal,
			}
		}

		for situation := range state.situations {
			ruleNumber, err := ga.findRuleBySituation(situation)
			if err != nil {
				return err
			}

			switch {
			case situation.position == len(situation.rightPart) &&
				situation.nonTerminal == startTerminal:
				newTableLine[situation.promise] = action{
					operationType: accept,
					rule:          ruleNumber,
					nonTerminal:   situation.nonTerminal,
				}
			case situation.position == len(situation.rightPart):
				if !((newTableLine[situation.promise].operationType == reduce &&
					newTableLine[situation.promise].rule == ruleNumber &&
					newTableLine[situation.promise].nonTerminal == situation.nonTerminal) ||
					(newTableLine[situation.promise].operationType == reject &&
						newTableLine[situation.promise].rule == -1)) {
					return errNotLR1
				}

				newTableLine[situation.promise] = action{
					operationType: reduce,
					rule:          ruleNumber,
					nonTerminal:   situation.nonTerminal,
				}
			case !ga.isNonTerminal(situation.rightPart[situation.position]):
				shiftSymbol := situation.rightPart[situation.position]
				if !((newTableLine[shiftSymbol].operationType == shift) ||
					(newTableLine[shiftSymbol].operationType == reject &&
						newTableLine[shiftSymbol].rule == -1)) {
					return errNotLR1
				}

				newTableLine[situation.rightPart[situation.position]] = action{
					operationType: shift,
					rule:          1,
					nonTerminal:   situation.nonTerminal,
				}
			}
		}
		ga.actions = append(ga.actions, newTableLine)
	}

	return nil
}

func (ga *grammarAdapter) findRuleBySituation(situation situation) (int, error) {
	for i, rule := range ga.grammar.Rules[situation.nonTerminal] {
		if rule == situation.rightPart {
			return i, nil
		}
	}
	return -1, errNoSuchRule
}
