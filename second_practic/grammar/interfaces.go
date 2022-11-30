package grammar

type Grammar interface {
	CheckWord(string) bool
	AddRule(byte, string)
}
