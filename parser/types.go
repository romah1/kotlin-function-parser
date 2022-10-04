package parser

type Token = string

const (
	Undefined       = ""
	LParen          = "("
	RParen          = ")"
	Comma           = ","
	Fun             = "fun"
	Colon           = ":"
	End             = "$"
	QuestionMark    = "?"
	ExclamationMark = "!"
)

type LexicalAnalyzer struct {
	inputStream string
	curChar     rune
	curPos      int
	curToken    Token
}

type Tree struct {
	node     string
	children []*Tree
}

type (
	Start                       = Tree
	Declaration                 = Tree
	FunctionName                = Tree
	Arguments                   = Tree
	VariableAndType             = Tree
	VariableAndTypeContinuation = Tree
	Variable                    = Tree
	Type                        = Tree
	TypeName                    = Tree
	TypeMark                    = Tree
	Ending                      = Tree
)

type Parser struct {
	lexicalAnalyzer *LexicalAnalyzer
}

type ParsingError struct {
	unexpected string
	expected   string
}
