package adapter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleFirst(t *testing.T) {
	grammarAdapter, err := BuildAdapter("./test/simple/test_grammar1.txt")
	if err != nil {
		panic(err)
	}

	testCasesSuccess := []string{
		"aabb",
		"aaabbb",
		"ab",
		"abababab",
		"abaaabbb",
		"abaabbaaabbbabab",
		"aaabbbabaaabbbabaaabbb",
		"aaaaabbaabbbabbb",
		"aaabbabbaaaaabbbabbb",
	}

	testCasesFail := []string{
		"b",
		"abb",
		"abbbababababaabbbbbbbabab",
		"ababbabbbbbbbbbbabbabbbabababbababa",
		"aaaaaaaabbbbbbbbbbaaaaabb",
		"aaaaaaaaaaaaaaaaaaaab",
		"aabababababaababababaaaaaabb",
		"aabababababababaaaaaaaabbbb",
	}

	for _, testCase := range testCasesSuccess {
		assert.True(t, grammarAdapter.Read(testCase))
		grammarAdapter.Flush()
	}

	for _, testCase := range testCasesFail {
		assert.False(t, grammarAdapter.Read(testCase))
		grammarAdapter.Flush()
	}
}

func TestSimpleSecond(t *testing.T) {
	grammarAdapter, err := BuildAdapter("./test/simple/test_grammar2.txt")
	if err != nil {
		panic(err)
	}

	testCasesSuccess := []string{
		"",
		"ababaabb",
		"aaabbbab",
		"abababababababababab",
		"aaaaabbbaabbbb",
		"abaabbabab",
		"aaaaaaaaaabbbbbaaabbbbbbbb",
		"ababaaabbbababaabbaaaabbbbababaaabbbabaaabbb",
		"abababababaaabbbaabbabab",
	}

	testCasesFail := []string{
		"abb",
		"abbbabbababababa",
		"aaabbbabbbbbaaabababababababa",
		"ababaababababbbabbabababababababbbbbba",
		"ababababababbbbbbbbbaaaa",
		"b",
		"abababababbbbbbbbbbbabbb",
	}

	for _, testCase := range testCasesSuccess {
		assert.True(t, grammarAdapter.Read(testCase))
		grammarAdapter.Flush()
	}

	for _, testCase := range testCasesFail {
		assert.False(t, grammarAdapter.Read(testCase))
		grammarAdapter.Flush()
	}
}

func TestSimpleThird(t *testing.T) {
	grammarAdapter, err := BuildAdapter("./test/simple/test_grammar3.txt")
	if err != nil {
		panic(err)
	}

	testCasesSuccess := []string{
		"aaaaba",
		"aaaaaaca",
		"aaba",
		"aaaaaaaaba",
		"aaaaaaaaaaca",
		"aaaaaaaaaaaaaaaaaaaaba",
		"aaaaaaaaaaaaaaaaca",
		"aaaaca",
		"aaaaaaaaaaaaba",
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaba",
	}

	testCasesFail := []string{
		"aaaba",
		"aaaaaaaaaaaca",
		"aba",
		"aaaaaaaaaaaaaaaaa",
		"bba",
		"abababababababa",
		"aaaaaaaca",
		"aaaaaaaaaca",
	}

	for _, testCase := range testCasesSuccess {
		assert.True(t, grammarAdapter.Read(testCase))
		grammarAdapter.Flush()
	}

	for _, testCase := range testCasesFail {
		assert.False(t, grammarAdapter.Read(testCase))
		grammarAdapter.Flush()
	}
}

func TestHardFirst(t *testing.T) {
	grammarAdapter, err := BuildAdapter("./test/hard/test_grammar1.txt")
	if err != nil {
		panic(err)
	}

	testCasesSuccess := []string{
		"cabc",
		"abcc",
		"aaaaaaaaaaaaaaaaaaaaaaaaab",
		"abcabc",
		"abcabababababababababc",
		"abababababababcabababababababc",
		"aaaaaaaaaaaaaaaaaaaaaabababababcababababc",
		"aaaaaaaaaab",
		"aaaaaaaaaaacc",
		"cc",
		"aaaaaaaaaaaabcc",
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaabcabc",
	}

	testCasesFail := []string{
		"abc",
		"acbc",
		"abababababababaababbab",
		"ababababababab",
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		"cababb",
		"cacbcbcbca",
		"cccccccccccccccccccc",
		"cbacbacbacbacbd",
		"dddddddddd",
		"sbcjdshcbsjhcbsj",
		"hjcdsbchsbcjhsbcks",
		"baaaaaaaabc",
		"aaaaaac",
	}

	for _, testCase := range testCasesSuccess {
		assert.True(t, grammarAdapter.Read(testCase))
		grammarAdapter.Flush()
	}

	for _, testCase := range testCasesFail {
		assert.False(t, grammarAdapter.Read(testCase))
		grammarAdapter.Flush()
	}
}

func TestHardSecond(t *testing.T) {
	grammarAdapter, err := BuildAdapter("./test/hard/test_grammar2.txt")
	if err != nil {
		panic(err)
	}

	testCasesSuccess := []string{
		"ccccccccccccccccccccccccdcccccccccccccccccccccccccccccccd",
		"acccccccccccccccccccccccccccd",
		"",
		"ccccccccdccccccccd",
		"acccd",
		"ad",
		"dd",
		"cccccccccccccccccdccccccccccccccccccccccccccccd",
		"ccccccccccccccccccccccccccccccccccdcccccccccccccccccccccccccccccccccccccccccccd",
		"accccccccccccccccccccccccccccccccccccccd",
		"cccccccccccccccccccccccccccdccccccccccccccccccccccccccccccd",
		"cccccccccccccccccccccccccccdccccccccccccccccccccccccccccccccccccccccccd",
	}

	testCasesFail := []string{
		"add",
		"aaaaaaaaaaacccccccccccccc",
		"acccccccccccccdcccccccccccccccd",
		"a",
		"acccccccccccccccccdcd",
		"acdcd",
		"sdkchsbksjvkd",
		"ccccccccccccdcdcdcddc",
		"cccccccccccccccccccccccccccccc",
		"cccccccccccccccccccccccccccccccccccccccccccccc",
		"skhdsbckscbskch",
		"ababababaabbcbbcbcbd",
		"azsbjascbsjhcsbcs",
		"ddddddddddddddddddc",
	}

	for _, testCase := range testCasesSuccess {
		assert.True(t, grammarAdapter.Read(testCase))
		grammarAdapter.Flush()
	}

	for _, testCase := range testCasesFail {
		assert.False(t, grammarAdapter.Read(testCase))
		grammarAdapter.Flush()
	}
}

func TestHardThird(t *testing.T) {
	grammarAdapter, err := BuildAdapter("./test/hard/test_grammar3.txt")
	if err != nil {
		panic(err)
	}

	testCasesSuccess := []string{
		"aaaaaaaaaacaaaaaaaaaa",
		"aca",
		"aaaaabbbcbbbaaaaa",
		"aacaa",
		"abababcbababa",
		"abababbabababababcbabababababbababa",
		"bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbcbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb",
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaacaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		"aaabaabaabababababababcbababababababaabaabaaa",
		"ababababababababababababbacabbabababababababababababa",
		"aaaaaaaaaaaaacaaaaaaaaaaaaa",
		"bcb",
		"aaaaaaaaaaaaaaaaaaaacbbbbbbbbbbbbbbbbbbbb",
	}

	testCasesFail := []string{
		"ababababababaabababcababababababababaababbb",
		"dhbbhcbshcbshcbd",
		"sdnskcjnskcjsnkc",
		"abababababababaabccccccababababaab",
		"bacbaaaa",
		"bbbbbbbaaaaaaaacbbbbbbba",
		"bababaababababcabababaaaaaaaaaaaaaaa",
		"babaababababaccccccccccd",
		"bbbbbbbbbbbbbbbbbbbbcbbbbb",
		"aaaaaaaaaaaaaaaaaaacaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		"ababababaababcabababaabbbbbbb",
	}

	for _, testCase := range testCasesSuccess {
		assert.True(t, grammarAdapter.Read(testCase))
		grammarAdapter.Flush()
	}

	for _, testCase := range testCasesFail {
		assert.False(t, grammarAdapter.Read(testCase))
		grammarAdapter.Flush()
	}
}

func TestStatesInfoFirst(t *testing.T) {
	grammarAdapter, err := BuildAdapter("./test/simple/test_grammar1.txt")
	if err != nil {
		panic(err)
	}

	infoGetter := newInfoGetter(grammarAdapter)
	states := infoGetter.getStates()
	assert.Equal(t, 10, len(states))

	assert.Equal(t, 3, len(states[0].situations))
	_, ok := states[0].transitions['a']
	assert.True(t, ok)
	_, ok = states[0].transitions['S']
	assert.True(t, ok)

	statesNum := states[0].transitions['S']
	assert.Equal(t, 1, len(states[statesNum].situations))
	assert.Equal(t, 0, len(states[statesNum].transitions))

	statesNum = states[0].transitions['a']
	assert.Equal(t, 3, len(states[statesNum].situations))
	_, ok = states[statesNum].transitions['a']
	assert.True(t, ok)
	_, ok = states[statesNum].transitions['S']
	assert.True(t, ok)
}

func TestStatesInfoSecond(t *testing.T) {
	grammarAdapter, err := BuildAdapter("./test/simple/test_grammar2.txt")
	if err != nil {
		panic(err)
	}

	infoGetter := newInfoGetter(grammarAdapter)
	states := infoGetter.getStates()
	assert.Equal(t, 8, len(states))

	assert.Equal(t, 1, len(states[0].transitions))
	assert.Equal(t, 5, len(states[0].situations))
	_, ok := states[0].transitions['S']
	assert.True(t, ok)

	statesNum := states[0].transitions['S']
	assert.Equal(t, 3, len(states[statesNum].situations))
	assert.Equal(t, 1, len(states[statesNum].transitions))

	_, ok = states[statesNum].transitions['a']
	assert.True(t, ok)
	statesNum = states[statesNum].transitions['a']
	assert.Equal(t, 6, len(states[statesNum].situations))
	assert.Equal(t, 1, len(states[statesNum].transitions))

	_, ok = states[statesNum].transitions['S']
	assert.True(t, ok)
	statesNum = states[statesNum].transitions['S']
	assert.Equal(t, 4, len(states[statesNum].situations))
	assert.Equal(t, 2, len(states[statesNum].transitions))

	_, ok = states[statesNum].transitions['a']
	assert.True(t, ok)
	_, ok = states[statesNum].transitions['b']
	assert.True(t, ok)

	statesNum = states[statesNum].transitions['a']
	assert.Equal(t, 6, len(states[statesNum].situations))
	assert.Equal(t, 1, len(states[statesNum].transitions))

	_, ok = states[statesNum].transitions['S']
	assert.True(t, ok)

	statesNum = states[statesNum].transitions['S']
	assert.Equal(t, 4, len(states[statesNum].situations))
	assert.Equal(t, 2, len(states[statesNum].transitions))

	_, ok = states[statesNum].transitions['a']
	assert.True(t, ok)
	_, ok = states[statesNum].transitions['b']
	assert.True(t, ok)

	statesNum = states[statesNum].transitions['b']
	assert.Equal(t, 2, len(states[statesNum].situations))
	assert.Equal(t, 0, len(states[statesNum].transitions))
}
