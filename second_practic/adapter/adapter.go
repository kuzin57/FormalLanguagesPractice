package adapter

import "github.com/kuzin57/FormalPractic/second_practic/builder"

func NewGrammarAdapter() GrammarAdapter {
	adapter := &grammarAdapter{
		currentConfigurations: make([]map[configuration]struct{}, 1),
	}
	adapter.initScan()
	adapter.initPredict()
	adapter.initComplete()
	return adapter
}

func (ga *grammarAdapter) BuildGrammar(input []string) {
	builder := builder.NewGrammarBuilder(input)
	ga.grammar = builder.Build()
}

func (ga *grammarAdapter) Flush() {
	ga.currentConfigurations = make([]map[configuration]struct{}, 1)
}

func newConfiguration(expression string, position int, startIndex int, terminal byte) configuration {
	return configuration{
		expression: expression,
		position:   position,
		startIndex: startIndex,
		terminal:   terminal,
	}
}

func (ga *grammarAdapter) initScan() {
	ga.scan = func(config configuration, indexToAdd int, word string) bool {
		var retValue bool
		if config.position < len(config.expression) {
			_, ok := ga.grammar.Rules[config.expression[config.position]]
			if !ok && word[indexToAdd] == config.expression[config.position] {
				config.position++
				_, retValue = ga.currentConfigurations[indexToAdd+1][config]
				retValue = !retValue
				ga.currentConfigurations[indexToAdd+1][config] = struct{}{}
			}
		}
		return retValue
	}
}

func (ga *grammarAdapter) initPredict() {
	ga.predict = func(config configuration, indexToAdd int, word string) bool {
		var retValue bool
		if config.position < len(config.expression) {
			rules, ok := ga.grammar.Rules[config.expression[config.position]]
			if ok {
				for _, rule := range rules {
					newConfig := newConfiguration(rule, 0, indexToAdd, config.expression[config.position])
					_, retValue = ga.currentConfigurations[indexToAdd][newConfig]
					retValue = !retValue
					ga.currentConfigurations[indexToAdd][newConfig] = struct{}{}
				}
			}
		}
		return retValue
	}
}

func (ga *grammarAdapter) initComplete() {
	ga.complete = func(config configuration, indexToAdd int, word string) bool {
		var retValue bool
		if config.position == len(config.expression) {
			terminal := config.terminal
			_, ok := ga.grammar.Rules[terminal]
			if ok {
				for conf := range ga.currentConfigurations[config.startIndex] {
					if conf.position < len(conf.expression) &&
						terminal == conf.expression[conf.position] {
						conf.position++
						_, retValue = ga.currentConfigurations[indexToAdd][conf]
						retValue = !retValue
						ga.currentConfigurations[indexToAdd][conf] = struct{}{}
					}
				}
			}
		}
		return retValue
	}
}

func (ga *grammarAdapter) Read(word string) bool {
	ga.word = word
	ga.grammar.AddNewStartSymbol()

	ga.currentConfigurations[0] = make(map[configuration]struct{})
	ga.currentConfigurations[0][newConfiguration("S", 0, 0, startSymbol)] = struct{}{}

	for {
		changedPredict := ga.updateCongigurations(0, ga.predict)
		changedComplete := ga.updateCongigurations(0, ga.complete)
		if !changedPredict && !changedComplete {
			break
		}
	}

	for i := range word {
		ga.currentConfigurations = append(
			ga.currentConfigurations,
			make(map[configuration]struct{}),
		)

		changedScan := ga.updateCongigurations(i, ga.scan)
		if !changedScan {
			return false
		}

		for {
			changedPredict := ga.updateCongigurations(i+1, ga.predict)
			changedComplete := ga.updateCongigurations(i+1, ga.complete)
			if !changedComplete && !changedPredict {
				break
			}
		}
	}
	_, inGrammar := ga.currentConfigurations[len(word)][newConfiguration("S", 1, 0, startSymbol)]
	return inGrammar
}

func (ga *grammarAdapter) updateCongigurations(j int, updateConfig func(configuration, int, string) bool) bool {
	var addedNew, retValue bool
	for config := range ga.currentConfigurations[j] {
		addedNew = updateConfig(config, j, ga.word)
		if addedNew {
			retValue = true
		}
	}
	return retValue
}