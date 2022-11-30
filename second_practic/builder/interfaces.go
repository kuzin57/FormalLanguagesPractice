package builder

import "github.com/kuzin57/FormalPractic/second_practic/grammar"

type GrammarBuilder interface {
	Build() grammar.Grammar
}
