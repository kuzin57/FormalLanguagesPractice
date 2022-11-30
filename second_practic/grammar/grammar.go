package grammar

func NewGrammar() Grammar {
	return &grammar{
		rules:                 make(map[byte][]string),
		currentConfigurations: make(map[configuration]struct{}),
	}
}

func newConfiguration(expression string, position int, startIndex int, terminal byte) configuration {
	return configuration{
		expression: expression,
		position:   position,
		startIndex: startIndex,
		terminal:   terminal,
	}
}

func (g *grammar) addNewStartSymbol() {
	g.rules[startSymbol] = []string{"S"}
}

func (g *grammar) CheckWord(word string) bool {
	g.addNewStartSymbol()

	for {
		changedPredict := g.predict(0)
		changedComplete := g.predict(0)
		if !changedPredict && !changedComplete {
			break
		}
	}

	for i := range word {
		changedScan := g.scan(i, word)
		if !changedScan {
			return false
		}

		for {
			changedPredict := g.predict(i + 1)
			changedComplete := g.complete(i + 1)
			if !changedComplete && !changedPredict {
				break
			}
		}
	}
	return true
}

func (g *grammar) AddRule(nt byte, rightPart string) {
	_, ok := g.rules[nt]
	if !ok {
		g.rules[nt] = []string{rightPart}
	} else {
		g.rules[nt] = append(g.rules[nt], rightPart)
	}
}

func (g *grammar) scan(index int, word string) bool {
	newConfigurations := make(map[configuration]struct{})
	var addedNew bool
	for config := range g.currentConfigurations {
		if config.position < len(config.expression) {
			_, ok := g.rules[config.expression[config.position]]
			if !ok && word[index] == config.expression[config.position] {
				addedNew = true
				config.position++
				newConfigurations[config] = struct{}{}
			}
		}
	}
	g.currentConfigurations = newConfigurations
	return addedNew
}

func (g *grammar) predict(j int) bool {
	var addedNew bool
	for config := range g.currentConfigurations {
		if config.position < len(config.expression) {
			rules, ok := g.rules[config.expression[config.position]]
			if ok {
				addedNew = true
				for _, rule := range rules {
					newConfig := newConfiguration(rule, 0, j, config.expression[config.position])
					g.currentConfigurations[newConfig] = struct{}{}
				}
			}
		}
	}
	return addedNew
}

func (g *grammar) complete(j int) bool {
	var addedNew bool
	for config := range g.currentConfigurations {
		if config.position == len(config.expression) {
			terminal := config.expression[config.position]
			_, ok := g.rules[terminal]
			if ok {
				addedNew = true
				for _, rule := range g.rules[terminal] {
					g.currentConfigurations[newConfiguration(rule, 0, j, terminal)] = struct{}{}
				}
			}
		}
	}
	return addedNew
}
