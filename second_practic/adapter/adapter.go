package adapter

import (
	"bufio"
	"fmt"
	"os"

	"github.com/kuzin57/FormalPractic/second_practic/builder"
)

func newGrammarAdapter(logFile string) GrammarAdapter {
	newLogger, err := newLogger(logFile)
	if err != nil {
		panic(err)
	}

	adapter := &grammarAdapter{
		currentConfigurations: make([]map[configuration]struct{}, 1),
		logger:                newLogger,
		configurationInfos:    make([]map[configuration]configurationInfo, 1),
	}
	adapter.initScan()
	adapter.initPredict()
	adapter.initComplete()
	return adapter
}

func BuildAdapter(filename string, logFile string) GrammarAdapter {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	var rules []string
	for scanner.Scan() {
		rules = append(rules, scanner.Text())
	}

	grammarAdapter := newGrammarAdapter(logFile)
	grammarAdapter.BuildGrammar(rules)
	return grammarAdapter
}

func (ga *grammarAdapter) BuildGrammar(input []string) {
	builder := builder.NewGrammarBuilder(input)
	ga.grammar = builder.Build()
}

func (ga *grammarAdapter) Flush() {
	ga.currentConfigurations = make([]map[configuration]struct{}, 1)
	ga.configurationInfos = make([]map[configuration]configurationInfo, 1)
	ga.logger.printEmptyLine()
	ga.currentConfNumber = 0
}

func newConfiguration(expression string, position int, startIndex int, terminal byte) configuration {
	return configuration{
		expression: expression,
		position:   position,
		startIndex: startIndex,
		terminal:   terminal,
	}
}

func newConfigurationInfo(method string, number int) configurationInfo {
	return configurationInfo{
		method: method,
		number: number,
	}
}

func (ga *grammarAdapter) initScan() {
	ga.scan = func(config configuration, indexToAdd int, word string) bool {
		var retValue bool
		if config.position >= len(config.expression) {
			return retValue
		}
		_, ok := ga.grammar.Rules[config.expression[config.position]]
		if !ok && word[indexToAdd] == config.expression[config.position] {
			newConf := config
			newConf.position++

			_, retValue = ga.currentConfigurations[indexToAdd+1][newConf]
			retValue = !retValue
			if !retValue {
				return retValue
			}

			ga.completeConfigurations = append(ga.completeConfigurations, newConf)
			ga.configurationInfos[indexToAdd+1][newConf] = newConfigurationInfo(
				fmt.Sprintf("scan %d", ga.configurationInfos[indexToAdd][config].number),
				ga.currentConfNumber,
			)
			ga.logger.info(ga.configurationInfos[indexToAdd+1][newConf], newConf)
			ga.currentConfNumber++
			ga.currentConfigurations[indexToAdd+1][newConf] = struct{}{}
		}
		return retValue
	}
}

func (ga *grammarAdapter) initPredict() {
	ga.predict = func(config configuration, indexToAdd int, word string) bool {
		var retValue bool
		if config.position >= len(config.expression) {
			return retValue
		}
		rules, ok := ga.grammar.Rules[config.expression[config.position]]
		if !ok {
			return retValue
		}
		for _, rule := range rules {
			newConfig := newConfiguration(rule, 0, indexToAdd, config.expression[config.position])
			_, retValue = ga.currentConfigurations[indexToAdd][newConfig]
			retValue = !retValue
			if !retValue {
				continue
			}

			ga.completeConfigurations = append(ga.completeConfigurations, newConfig)
			ga.configurationInfos[indexToAdd][newConfig] = newConfigurationInfo(
				fmt.Sprintf("predict %d", ga.configurationInfos[indexToAdd][config].number),
				ga.currentConfNumber,
			)
			ga.logger.info(ga.configurationInfos[indexToAdd][newConfig], newConfig)
			ga.currentConfNumber++
			ga.currentConfigurations[indexToAdd][newConfig] = struct{}{}
		}
		return retValue
	}
}

func (ga *grammarAdapter) initComplete() {
	ga.complete = func(config configuration, indexToAdd int, word string) bool {
		var retValue bool
		if config.position != len(config.expression) {
			return retValue
		}
		terminal := config.terminal
		_, ok := ga.grammar.Rules[terminal]
		if !ok {
			return retValue
		}
		for conf := range ga.currentConfigurations[config.startIndex] {
			if conf.position >= len(conf.expression) || terminal != conf.expression[conf.position] {
				continue
			}
			newConf := conf
			newConf.position++

			_, retValue = ga.currentConfigurations[indexToAdd][newConf]
			retValue = !retValue
			if !retValue {
				continue
			}

			ga.completeConfigurations = append(ga.completeConfigurations, newConf)
			ga.configurationInfos[indexToAdd][newConf] = newConfigurationInfo(
				fmt.Sprintf(
					"complete %d, %d",
					ga.configurationInfos[indexToAdd][config].number,
					ga.configurationInfos[indexToAdd][conf].number,
				),
				ga.currentConfNumber,
			)

			ga.logger.info(ga.configurationInfos[indexToAdd][newConf], newConf)
			ga.currentConfNumber++
			ga.currentConfigurations[indexToAdd][newConf] = struct{}{}
		}
		return retValue
	}
}

func (ga *grammarAdapter) Read(word string) bool {
	ga.word = word
	ga.addNewStartSymbol()

	ga.currentConfigurations[0] = make(map[configuration]struct{})
	ga.configurationInfos[0] = make(map[configuration]configurationInfo)

	startConfiguration := newConfiguration("S", 0, 0, startSymbol)
	ga.currentConfigurations[0][startConfiguration] = struct{}{}
	ga.completeConfigurations = make([]configuration, 1)
	ga.completeConfigurations[0] = startConfiguration

	ga.configurationInfos[0][startConfiguration] = newConfigurationInfo("init", 0)
	ga.currentConfNumber++

	ga.logger.printD(0)
	ga.logger.info(ga.configurationInfos[0][startConfiguration], startConfiguration)
	for {
		changedPredict := ga.updateCongigurations(0, ga.predict)
		length := len(ga.completeConfigurations)
		changedComplete := ga.updateCongigurations(0, ga.complete)
		ga.completeConfigurations = ga.completeConfigurations[length:]
		if !changedPredict && !changedComplete {
			break
		}
	}

	for i := range word {
		ga.currentConfigurations = append(
			ga.currentConfigurations,
			make(map[configuration]struct{}),
		)
		ga.configurationInfos = append(
			ga.configurationInfos,
			make(map[configuration]configurationInfo),
		)

		ga.logger.printD(i + 1)
		ga.completeConfigurations = nil
		changedScan := ga.updateCongigurations(i, ga.scan)
		if !changedScan {
			return false
		}

		for {
			changedPredict := ga.updateCongigurations(i+1, ga.predict)
			length := len(ga.completeConfigurations)
			changedComplete := ga.updateCongigurations(i+1, ga.complete)
			ga.completeConfigurations = ga.completeConfigurations[length:]
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

func (ga *grammarAdapter) addNewStartSymbol() {
	ga.grammar.Rules[startSymbol] = []string{"S"}
}

func newConfigurationGetter(adapter GrammarAdapter) configurationsGetter {
	grammarAdapter := adapter.(*grammarAdapter)
	return &configGetter{
		adapter: grammarAdapter,
	}
}

func (cg *configGetter) GetConfigurations() []map[configuration]struct{} {
	return cg.adapter.currentConfigurations
}

func (cg *configGetter) GetConfigurationsInfos() []map[configuration]configurationInfo {
	return cg.adapter.configurationInfos
}
