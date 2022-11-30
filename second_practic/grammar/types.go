package grammar

const (
	startSymbol = '$'
)

type Grammar struct {
	Rules map[byte][]string
}
