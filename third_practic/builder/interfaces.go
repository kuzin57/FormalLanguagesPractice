package builder

import "github.com/kuzin57/FormalPractic/third_practic/grammar"

type GrammarBuilder interface {
	Build() *grammar.Grammar
}
