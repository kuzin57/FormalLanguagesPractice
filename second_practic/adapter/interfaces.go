package adapter

type GrammarAdapter interface {
	Read(string) bool
	BuildGrammar([]string)
	Flush()
}

type configurationsGetter interface {
	GetConfigurations() []map[configuration]struct{}
	GetConfigurationsInfos() []map[configuration]configurationInfo
}
