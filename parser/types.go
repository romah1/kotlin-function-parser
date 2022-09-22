package parser

type Token = string

const (
	LParen = "("
	RParen = ")"
	Comma  = ","
	End    = "$"
)

type LexicalAnalyzer struct {
	inputStream string
	curChar     byte
	curPos      int
	curToken    Token
}
