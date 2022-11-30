package grammar

func NewGrammar() *Grammar {
	return &Grammar{
		Rules: make(map[byte][]string),
	}
}

// func newReader(grammar *grammar, word string) *reader {
// 	return &reader{
// 		grammar:               grammar,
// 		currentConfigurations: make([]map[configuration]struct{}, 1),
// 		word:                  word,
// 	}
// }

func (g *Grammar) AddNewStartSymbol() {
	g.Rules[startSymbol] = []string{"S"}
}

// func (g *Grammar) CheckWord(word string) bool {
// 	reader := newReader(g, word)
// 	reader.initScan()
// 	reader.initPredict()
// 	reader.initComplete()
// 	return reader.read(word)
// }

func (g *Grammar) AddRule(nt byte, rightPart string) {
	_, ok := g.Rules[nt]
	if !ok {
		g.Rules[nt] = []string{rightPart}
	} else {
		g.Rules[nt] = append(g.Rules[nt], rightPart)
	}
}
