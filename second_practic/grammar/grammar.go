package grammar

func NewGrammar() *Grammar {
	return &Grammar{
		Rules: make(map[byte][]string),
	}
}

func (g *Grammar) AddNewStartSymbol() {
	g.Rules[startSymbol] = []string{"S"}
}

func (g *Grammar) AddRule(nt byte, rightPart string) {
	_, ok := g.Rules[nt]
	if !ok {
		g.Rules[nt] = []string{rightPart}
	} else {
		g.Rules[nt] = append(g.Rules[nt], rightPart)
	}
}
