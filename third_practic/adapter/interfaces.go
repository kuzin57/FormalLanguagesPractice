package adapter

type GrammarAdapter interface {
	Read(string) bool
	BuildGrammar([]string)
	Flush()
}
