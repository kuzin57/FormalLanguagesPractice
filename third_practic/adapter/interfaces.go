package adapter

type GrammarAdapter interface {
	Read(string) bool
	BuildGrammar([]string)
	Flush()
}

type AdapterInfoGetter interface {
	GetStates() []*state
}
