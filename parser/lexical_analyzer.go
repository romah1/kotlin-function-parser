package parser

import (
	"errors"
	"fmt"
)

func NewLexicalAnalyzer(input string) *LexicalAnalyzer {
	lexicalAnalyzer := &LexicalAnalyzer{
		inputStream: input + End,
		curChar:     0,
		curPos:      0,
		curToken:    Undefined,
	}
	lexicalAnalyzer.nextChar()
	return lexicalAnalyzer
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

func (lexicalAnalyzer *LexicalAnalyzer) hasNext() bool {
	return lexicalAnalyzer.curPos < len(lexicalAnalyzer.inputStream)
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
	case ':':
		lexicalAnalyzer.nextChar()
		lexicalAnalyzer.curToken = Colon
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

//func (lexicalAnalyzer *LexicalAnalyzer) takeUntil(c rune) error {
//	newToken := Undefined
//	for lexicalAnalyzer.hasNext() && lexicalAnalyzer.curChar != c {
//		newToken += Token(lexicalAnalyzer.curChar)
//		lexicalAnalyzer.nextChar()
//	}
//	if lexicalAnalyzer.curChar == c {
//		lexicalAnalyzer.curToken = newToken
//		return nil
//	} else {
//		return errors.New(fmt.Sprintf("Failed to reach character %c", c))
//	}
//}

func (lexicalAnalyzer *LexicalAnalyzer) takeUntil(ignoreRunes map[rune]bool) error {
	newToken := Undefined
	for lexicalAnalyzer.hasNext() && !ignoreRunes[lexicalAnalyzer.curChar] {
		newToken += Token(lexicalAnalyzer.curChar)
		lexicalAnalyzer.nextChar()
	}
	if ignoreRunes[lexicalAnalyzer.curChar] {
		lexicalAnalyzer.curToken = newToken
		return nil
	} else {
		return errors.New(fmt.Sprintf("Failed to reach character %T", ignoreRunes))
	}
}
