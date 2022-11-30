package builder

import (
	"strings"

	"github.com/kuzin57/FormalPractic/second_practic/grammar"
)

func NewGrammarBuilder(input []string) GrammarBuilder {
	newBuilder := &grammarBuilder{}
	newBuilder.input = input
	return newBuilder
}

func splitByArrow(line string) (string, string) {
	splitted := strings.SplitN(line, arrow, 2)
	return splitted[0], splitted[1]
}

func (gb *grammarBuilder) Build() grammar.Grammar {
	newGrammar := grammar.NewGrammar()
	for _, line := range gb.input {
		leftPart, rightPart := splitByArrow(line)
		for _, expression := range strings.Split(rightPart, pipe) {
			newGrammar.AddRule(leftPart[0], strings.TrimSpace(expression))
		}
	}
	return newGrammar
}
