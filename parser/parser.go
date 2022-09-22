package parser

func NewLexicalAnalyzer(input string) (parser *LexicalAnalyzer) {
	parser = &LexicalAnalyzer{
		inputStream: input,
		curChar:     0,
		curPos:      0,
		curToken:    "",
	}
	parser.nextChar()
	return parser
}

func (parser *LexicalAnalyzer) Token() Token {
	return parser.curToken
}

func (parser *LexicalAnalyzer) Pos() int {
	return parser.curPos
}

func (parser *LexicalAnalyzer) nextChar() {
	parser.curChar = parser.inputStream[parser.curPos]
	parser.curPos++
}

func (parser *LexicalAnalyzer) nextToken() {
	for parser.isWhitespace(parser.curChar) {
		parser.nextChar()
	}
	switch parser.curChar {
	case '(':
		parser.nextChar()
		parser.curToken = LParen
	case ')':
		parser.nextChar()
		parser.curToken = RParen
	case '$':
		parser.curToken = End
	default:
		parser.curToken += string(parser.curChar)
	}
}

func (parser *LexicalAnalyzer) isWhitespace(c byte) bool {
	return c == ' ' || c == '\r' || c == '\n' || c == '\t'
}
