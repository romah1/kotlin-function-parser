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
	if lexicalAnalyzer.curToken == End {
		return errors.New(fmt.Sprintf("Tried to get token after %s was seen", End))
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
	case '$':
		lexicalAnalyzer.curToken = End
	case 'f':
		if lexicalAnalyzer.checkMatches("fun") {
			lexicalAnalyzer.curToken = Fun
			return nil
		}
		fallthrough
	default:
		charSet := map[rune]bool{
			'(': true,
			')': true,
			',': true,
			':': true,
			'$': true,
		}
		newToken := ""
		for !charSet[lexicalAnalyzer.curChar] && !lexicalAnalyzer.isWhitespace(lexicalAnalyzer.curChar) {
			newToken += Token(lexicalAnalyzer.curChar)
			lexicalAnalyzer.nextChar()
		}
		lexicalAnalyzer.curToken = newToken
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
	if !lexicalAnalyzer.isWhitespace(lexicalAnalyzer.curChar) {

	}
	return nil
}

func (lexicalAnalyzer *LexicalAnalyzer) checkMatches(s string) bool {
	oldPos := lexicalAnalyzer.curPos
	oldChar := lexicalAnalyzer.curChar
	for _, c := range s {
		if lexicalAnalyzer.curChar == c {
			lexicalAnalyzer.nextChar()
		} else {
			lexicalAnalyzer.curPos = oldPos
			return false
		}
	}
	//lexicalAnalyzer.nextChar()
	if !lexicalAnalyzer.isWhitespace(lexicalAnalyzer.curChar) {
		lexicalAnalyzer.curPos = oldPos
		lexicalAnalyzer.curChar = oldChar
		return false
	}
	return true
}
