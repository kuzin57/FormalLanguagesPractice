package adapter

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimple(t *testing.T) {
	testGrammars := []string{
		"./test/simple/test_grammar1.txt",
		"./test/simple/test_grammar2.txt",
		"./test/simple/test_grammar3.txt",
	}
	testWords := [][]string{
		{"abc", "ab", "aabcbc", "aabbc"},
		{"aa", "aab", "aaab"},
		{"aaaabbbb", "ab", "", "aaaabb"},
	}

	for i, grammar := range testGrammars {
		grammarAdapter := BuildAdapter(grammar, "log.txt")
		for _, word := range testWords[i] {
			assert.Equal(t, true, grammarAdapter.Read(word))
			grammarAdapter.Flush()
		}
	}
}

func TestHard(t *testing.T) {
	testGrammars := []string{
		"./test/hard/test_grammar1.txt",
		"./test/hard/test_grammar2.txt",
	}
	testWordsSuccess := [][]string{
		{
			"aaabbbbbbbaa",
			"aababbbbb",
			"aaabbbbbbbaaaaabbbbbbbaa",
			"abbbbbbbbbbbb",
			"aababb",
			"cccabaaaa",
			"cccabaaaacccabaaaa",
			"cccabaabaabacaa",
			"cccabaabaabacaacccabaaaacaa",
			"abababababaaaabbbbbbbbbbbbbbb",
			"ababababaaabbbbbbbaaaa",
		},
		{
			"abbbabb",
			"a",
			"abbbabba",
			"bbbbbabbbbb",
			"abbbbbabbbbb",
			"",
			"bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb",
			"bbbbbbbbbbbbbbbbbbbbabb",
			"bb",
			"bbbbabbbbbbbbbbbbbbbbbbbbbbbbbbbb",
		},
	}

	testWordsFail := [][]string{
		{
			"cca",
			"abcbabcbab",
			"cccccccd",
			"sjskdjfkj",
			"cbbbbb",
			"ccccccccc",
		},
		{
			"cccc",
			"bbbbbbaaaaaabbbbbbb",
			"bbbbbaabbbbaaa",
			"aaaaaaaaaaaaab",
			"aaaabbbbbbbbaaaaaaaabaaaaaaaaab",
			"abababababababababbbbbb",
		},
	}

	for i, grammar := range testGrammars {
		grammarAdapter := BuildAdapter(grammar, "log.txt")
		for _, word := range testWordsSuccess[i] {
			assert.True(t, grammarAdapter.Read(word))
			grammarAdapter.Flush()
		}
		for _, word := range testWordsFail[i] {
			assert.False(t, grammarAdapter.Read(word))
			grammarAdapter.Flush()
		}
	}
}

func TestCompletePredict(t *testing.T) {
	grammarAdapter := BuildAdapter("./test/test_predict_complete/test_grammar.txt", "log.txt")
	word := "aabbab"
	grammarAdapter.Read(word)

	configurationsGetter := newConfigurationGetter(grammarAdapter)
	configurations := configurationsGetter.GetConfigurations()

	assert.Equal(t, 7, len(configurations))
	assert.Equal(t, 5, len(configurations[0]))
	assert.Equal(t, 5, len(configurations[1]))
	assert.Equal(t, 5, len(configurations[2]))
	assert.Equal(t, 3, len(configurations[3]))
	assert.Equal(t, 3, len(configurations[4]))
	assert.Equal(t, 5, len(configurations[5]))

	configurationInfos := configurationsGetter.GetConfigurationsInfos()
	var operationsCounter map[string]int
	expectedNumberOperations := make([]map[string]int, len(configurationInfos))
	for i := 0; i < len(configurationInfos); i++ {
		expectedNumberOperations[i] = make(map[string]int)
	}

	expectedNumberOperations[0]["init"] = 1
	expectedNumberOperations[0]["predict"] = 2
	expectedNumberOperations[0]["complete"] = 2

	expectedNumberOperations[1]["scan"] = 1
	expectedNumberOperations[1]["predict"] = 2
	expectedNumberOperations[1]["complete"] = 2

	expectedNumberOperations[2]["scan"] = 1
	expectedNumberOperations[2]["predict"] = 2
	expectedNumberOperations[2]["complete"] = 2

	expectedNumberOperations[3]["scan"] = 1
	expectedNumberOperations[3]["complete"] = 2

	expectedNumberOperations[4]["scan"] = 1
	expectedNumberOperations[4]["complete"] = 2

	expectedNumberOperations[5]["scan"] = 1
	expectedNumberOperations[5]["predict"] = 2
	expectedNumberOperations[5]["complete"] = 2

	for i, configInfos := range configurationInfos {
		operationsCounter = make(map[string]int)
		for _, configInfo := range configInfos {
			pureMethod := strings.Split(configInfo.method, " ")[0]
			_, ok := operationsCounter[pureMethod]
			if !ok {
				operationsCounter[pureMethod] = 1
				continue
			}
			operationsCounter[pureMethod]++
		}

		for operation, number := range expectedNumberOperations[i] {
			assert.Equal(t, number, operationsCounter[operation])
		}
	}

}
