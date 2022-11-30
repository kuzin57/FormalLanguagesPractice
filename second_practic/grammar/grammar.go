package grammar

func NewGrammar() Grammar {
	return &grammar{
		rules: make(map[byte][]string),
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

func newReader(grammar *grammar, word string) *reader {
	return &reader{
		grammar:               grammar,
		currentConfigurations: make([]map[configuration]struct{}, 1),
		word:                  word,
	}
}

func (g *grammar) addNewStartSymbol() {
	g.rules[startSymbol] = []string{"S"}
}

func (g *grammar) CheckWord(word string) bool {
	reader := newReader(g, word)
	reader.initScan()
	reader.initPredict()
	reader.initComplete()
	return reader.read(word)
}

func (r *reader) initScan() {
	r.scan = func(config configuration, indexToAdd int, word string) bool {
		var retValue bool
		if config.position < len(config.expression) {
			_, ok := r.grammar.rules[config.expression[config.position]]
			if !ok && word[indexToAdd] == config.expression[config.position] {
				config.position++
				_, retValue = r.currentConfigurations[indexToAdd+1][config]
				retValue = !retValue
				r.currentConfigurations[indexToAdd+1][config] = struct{}{}
			}
		}
		return retValue
	}
}

func (r *reader) initPredict() {
	r.predict = func(config configuration, indexToAdd int, word string) bool {
		var retValue bool
		if config.position < len(config.expression) {
			rules, ok := r.grammar.rules[config.expression[config.position]]
			if ok {
				for _, rule := range rules {
					newConfig := newConfiguration(rule, 0, indexToAdd, config.expression[config.position])
					_, retValue = r.currentConfigurations[indexToAdd][newConfig]
					retValue = !retValue
					r.currentConfigurations[indexToAdd][newConfig] = struct{}{}
				}
			}
		}
		return retValue
	}
}

func (r *reader) initComplete() {
	r.complete = func(config configuration, indexToAdd int, word string) bool {
		var retValue bool
		if config.position == len(config.expression) {
			terminal := config.terminal
			_, ok := r.grammar.rules[terminal]
			if ok {
				for conf := range r.currentConfigurations[config.startIndex] {
					if conf.position < len(conf.expression) &&
						terminal == conf.expression[conf.position] {
						conf.position++
						_, retValue = r.currentConfigurations[indexToAdd][conf]
						retValue = !retValue
						r.currentConfigurations[indexToAdd][conf] = struct{}{}
					}
				}
			}
		}
		return retValue
	}
}

func (r *reader) read(word string) bool {
	r.grammar.addNewStartSymbol()

	r.currentConfigurations[0] = make(map[configuration]struct{})
	r.currentConfigurations[0][newConfiguration("S", 0, 0, startSymbol)] = struct{}{}

	for {
		changedPredict := r.updateCongigurations(0, r.predict)
		changedComplete := r.updateCongigurations(0, r.complete)
		if !changedPredict && !changedComplete {
			break
		}
	}

	for i := range word {
		r.currentConfigurations = append(
			r.currentConfigurations,
			make(map[configuration]struct{}),
		)

		changedScan := r.updateCongigurations(i, r.scan)
		if !changedScan {
			return false
		}

		for {
			changedPredict := r.updateCongigurations(i+1, r.predict)
			changedComplete := r.updateCongigurations(i+1, r.complete)
			if !changedComplete && !changedPredict {
				break
			}
		}
	}
	_, inGrammar := r.currentConfigurations[len(word)][newConfiguration("S", 1, 0, startSymbol)]
	return inGrammar
}

func (g *grammar) AddRule(nt byte, rightPart string) {
	_, ok := g.rules[nt]
	if !ok {
		g.rules[nt] = []string{rightPart}
	} else {
		g.rules[nt] = append(g.rules[nt], rightPart)
	}
}

func (r *reader) updateCongigurations(j int, updateConfig func(configuration, int, string) bool) bool {
	var addedNew, retValue bool
	for config := range r.currentConfigurations[j] {
		addedNew = updateConfig(config, j, r.word)
		if addedNew {
			retValue = true
		}
	}
	return retValue
}
