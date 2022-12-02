package adapter_test

import (
	"fmt"
	"testing"

	"github.com/kuzin57/FormalPractic/second_practic/adapter"
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
		grammarAdapter := adapter.BuildAdapter(grammar)
		fmt.Println("adapter", grammarAdapter)
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
		grammarAdapter := adapter.BuildAdapter(grammar)
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
