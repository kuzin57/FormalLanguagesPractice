package grammar

const (
	startSymbol = '$'
)

type grammar struct {
	rules map[byte][]string
}

type configuration struct {
	expression string
	position   int
	startIndex int
	terminal   byte
}

type reader struct {
	grammar                 *grammar
	currentConfigurations   []map[configuration]struct{}
	predict, complete, scan func(configuration, int, string) bool
	word                    string
}
