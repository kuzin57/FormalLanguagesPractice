package builder_test

import (
	"testing"

	"github.com/kuzin57/FormalPractic/third_practic/builder"
	"github.com/stretchr/testify/assert"
)

func TestBuild(t *testing.T) {
	testCases := [][]string{
		{
			"S->SS|AA|a|",
			"A->AA|a|aAa|aaA|",
		},
		{
			"S->SS|B|A|AB",
			"A->a|b|AbAba|",
			"B->b|BBb|BaBb",
		},
		{
			"S->CC|B|a|",
			"C->Cc|c|",
			"B->bB|bb|BC|",
		},
	}

	builder0 := builder.NewGrammarBuilder(testCases[0])
	grammar0 := builder0.Build()
	assert.Equal(t, 4, len(grammar0.Rules['S']))
	assert.Equal(t, 5, len(grammar0.Rules['A']))

	builder1 := builder.NewGrammarBuilder(testCases[1])
	grammar1 := builder1.Build()
	assert.Equal(t, 4, len(grammar1.Rules['S']))
	assert.Equal(t, 4, len(grammar1.Rules['A']))
	assert.Equal(t, 3, len(grammar1.Rules['B']))

	builder2 := builder.NewGrammarBuilder(testCases[2])
	grammar2 := builder2.Build()
	assert.Equal(t, 4, len(grammar2.Rules['S']))
	assert.Equal(t, 3, len(grammar2.Rules['C']))
	assert.Equal(t, 4, len(grammar2.Rules['B']))
}
