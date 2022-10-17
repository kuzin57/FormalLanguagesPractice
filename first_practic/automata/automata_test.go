package automata_test

import (
	"testing"

	"github.com/kuzin57/FormalPractic/first_practic/automata"
	"github.com/kuzin57/FormalPractic/first_practic/parser"

	"github.com/stretchr/testify/assert"
)

var (
	regularExpressions = []string{
		"(a.a+b.b+(a.b+b.a).(b.b+a.a)*.(b.a+a.b))*",
		"a.c",
		"a*",
		"a.b.c",
		"(a.a.b+a+a.b.a+b.b.a)*",
		"(a*.b)*",
		"(a.b+b.a+a.a.b)*",
		"(a+b)*",
	}
)

func SetupAutomata(regularExpression string) automata.Automata {
	parser := parser.NewParser(regularExpression, nil)

	return automata.CreateAutomata(parser)
}

func TestNewAutomata(t *testing.T) {
	newAutomata := SetupAutomata(regularExpressions[0])

	assert.NotNil(t, newAutomata)
}

func TestConcat(t *testing.T) {
	testCases := []string{"ac", "abc", "abcd", "aaa"}
	expected := []int{0, 3, 3, 0}

	newAutomata := SetupAutomata(regularExpressions[3])
	assert.NotNil(t, newAutomata)

	for i, testCase := range testCases {
		assert.Equal(t, expected[i], newAutomata.ReadMaxPrefix(testCase))
		newAutomata.Flush()
	}
}

func TestJoin(t *testing.T) {
	testCases := []string{"ababababab", "abbaaababa", "babaaabaabbb", "acc"}
	expected := []int{10, 9, 10, 0}

	newAutomata := SetupAutomata(regularExpressions[6])
	assert.NotNil(t, newAutomata)

	for i, testCase := range testCases {
		ans := newAutomata.ReadMaxPrefix(testCase)
		assert.Equal(t, expected[i], ans)
		newAutomata.Flush()
	}
}

func TestCycle(t *testing.T) {
	testCases := []string{"abababab", "aaababab", "ababababababaaaaa", "aaaaaaaaac", "bbbbbbbbbbc"}
	expected := []int{8, 8, 17, 9, 10}

	newAutomata := SetupAutomata(regularExpressions[7])
	assert.NotNil(t, newAutomata)

	for i, testCase := range testCases {
		ans := newAutomata.ReadMaxPrefix(testCase)
		assert.Equal(t, expected[i], ans)
		newAutomata.Flush()
	}
}
