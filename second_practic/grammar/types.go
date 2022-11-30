package grammar

const (
	startSymbol = '$'
)

type grammar struct {
	rules                 map[byte][]string
	currentConfigurations map[configuration]struct{}
}

type configuration struct {
	expression string
	position   int
	startIndex int
	terminal   byte
}
