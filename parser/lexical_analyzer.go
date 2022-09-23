package parser

import (
	"errors"
	"fmt"
)

func NewLexicalAnalyzer(input string) (parser *LexicalAnalyzer) {
	parser = &LexicalAnalyzer{
		inputStream: input,
		curChar:     0,
		curPos:      0,
		curToken:    Undefined,
	}
	parser.nextChar()
	return parser
}

func (lexicalAnalyzer *LexicalAnalyzer) Token() Token {
	return lexicalAnalyzer.curToken
}

func (lexicalAnalyzer *LexicalAnalyzer) Pos() int {
	return lexicalAnalyzer.curPos
}

func (lexicalAnalyzer *LexicalAnalyzer) nextChar() {
	lexicalAnalyzer.curChar = rune(lexicalAnalyzer.inputStream[lexicalAnalyzer.curPos])
	lexicalAnalyzer.curPos++
}

func (lexicalAnalyzer *LexicalAnalyzer) nextToken() error {
	for lexicalAnalyzer.isWhitespace(lexicalAnalyzer.curChar) {
		lexicalAnalyzer.nextChar()
	}
	switch lexicalAnalyzer.curChar {
	case '(':
		lexicalAnalyzer.nextChar()
		lexicalAnalyzer.curToken = LParen
	case ')':
		lexicalAnalyzer.nextChar()
		lexicalAnalyzer.curToken = RParen
	case ',':
		lexicalAnalyzer.nextChar()
		lexicalAnalyzer.curToken = Comma
	case 'f':
		err := lexicalAnalyzer.matchString("fun")
		if err != nil {
			return err
		}
		lexicalAnalyzer.curToken = Fun
	case '$':
		lexicalAnalyzer.curToken = End
	default:
		return errors.New(fmt.Sprintf("Illegal character %c at pos %d", lexicalAnalyzer.curChar, lexicalAnalyzer.curPos))
	}
	return nil
}

func (lexicalAnalyzer *LexicalAnalyzer) isWhitespace(c rune) bool {
	return c == ' ' || c == '\r' || c == '\n' || c == '\t'
}

func (lexicalAnalyzer *LexicalAnalyzer) matchString(s string) error {
	oldPos := lexicalAnalyzer.curPos
	for _, c := range s {
		if lexicalAnalyzer.curChar == c {
			lexicalAnalyzer.nextChar()
		} else {
			lexicalAnalyzer.curPos = oldPos
			return errors.New(fmt.Sprintf(
				"%c character expected, %c found at pos %d",
				c,
				lexicalAnalyzer.curChar,
				lexicalAnalyzer.curPos))
		}
	}
	lexicalAnalyzer.nextChar()
	return nil
}
