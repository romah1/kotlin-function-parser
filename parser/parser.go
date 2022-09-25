package parser

import (
	"errors"
	"fmt"
)

func NewParser(input string) *Parser {
	return &Parser{
		lexicalAnalyzer: NewLexicalAnalyzer(input),
	}
}

func (parser *Parser) Parse() (*Tree, error) {
	err := parser.lexicalAnalyzer.nextToken()
	if err != nil {
		return nil, err
	}
	return parser.start()
}

func (parser *Parser) start() (*Tree, error) {
	switch parser.lexicalAnalyzer.curToken {
	case Fun:
		declarationTree, err := parser.declaration()
		if err != nil {
			return nil, err
		}
		return NewStart(declarationTree), nil
	default:
		return nil, errors.New(fmt.Sprintf("Found unexpected token %s", parser.lexicalAnalyzer.curToken))
	}
}

func (parser *Parser) declaration() (*Tree, error) {
	functionNameTree, err := parser.functionName()
	if err != nil {
		return nil, err
	}
	if parser.lexicalAnalyzer.curToken == LParen {
		_ = parser.lexicalAnalyzer.nextToken()
		argumentsTree, err := parser.arguments()
		if err != nil {
			return nil, err
		}
		if parser.lexicalAnalyzer.curToken == RParen {
			err := parser.lexicalAnalyzer.nextToken()
			if err != nil {
				return nil, err
			}
			endingTree, err := parser.ending()
			if err != nil {
				return nil, err
			}
			return NewDeclaration(functionNameTree, argumentsTree, endingTree), nil
		}
	}
	return nil, errors.New("wtf")
}

func (parser *Parser) functionName() (*Tree, error) {
	err := parser.lexicalAnalyzer.takeUntil(map[rune]bool{rune(LParen[0]): true})
	if err != nil {
		return nil, err
	}
	name := parser.lexicalAnalyzer.curToken
	err = parser.lexicalAnalyzer.nextToken()
	if err != nil {
		return nil, err
	}
	return NewFunctionName(name), nil
}

func (parser *Parser) ending() (*Tree, error) {
	switch parser.lexicalAnalyzer.curToken {
	case Colon:
		_ = parser.lexicalAnalyzer.nextToken()
		typeTree, err := parser.variableType()
		if err != nil {
			return nil, err
		}
		return NewEnding(typeTree), nil
	case End:
		return EmptyEnding, nil
	default:
		return nil, errors.New("wtf")
	}
}

func (parser *Parser) arguments() (*Tree, error) {
	if parser.lexicalAnalyzer.curToken == RParen {
		return EmptyArguments, nil
	}
	variableAndTypeTree, err := parser.variableAndType()
	if err != nil {
		return nil, err
	}
	variableAndTypeTreeContinuation, err := parser.variableAndTypeContinuation()
	if err != nil {
		return nil, err
	}
	return NewArguments(variableAndTypeTree, variableAndTypeTreeContinuation), nil
}

func (parser *Parser) variableAndType() (*Tree, error) {
	variableTree, err := parser.variable()
	if err != nil {
		return nil, err
	}
	if parser.lexicalAnalyzer.curToken == Colon {
		_ = parser.lexicalAnalyzer.nextToken()
		typeTree, err := parser.variableType()
		if err != nil {
			return nil, err
		}
		return NewVariableAndType(variableTree, typeTree), nil
	} else {
		return nil, errors.New("wtf")
	}

}

func (parser *Parser) variableAndTypeContinuation() (*Tree, error) {
	if parser.lexicalAnalyzer.curToken == Comma {
		_ = parser.lexicalAnalyzer.nextToken()
		variableAndTypeTree, err := parser.variableAndType()
		if err != nil {
			return nil, err
		}
		return NewVariableAndTypeContinuation(variableAndTypeTree), nil
	} else {
		return EmptyVariableAndTypeContinuation, nil
	}
}

func (parser *Parser) variable() (*Tree, error) {
	err := parser.lexicalAnalyzer.takeUntil(map[rune]bool{':': true})
	if err != nil {
		return nil, err
	}
	variable := parser.lexicalAnalyzer.curToken
	err = parser.lexicalAnalyzer.nextToken()
	if err != nil {
		return nil, err
	}
	return NewVariable(variable), nil
}

func (parser *Parser) variableType() (*Tree, error) {
	err := parser.lexicalAnalyzer.takeUntil(map[rune]bool{',': true, ')': true, '$': true})
	if err != nil {
		return nil, err
	}
	variable := parser.lexicalAnalyzer.curToken
	err = parser.lexicalAnalyzer.nextToken()
	if err != nil {
		return nil, err
	}
	return NewType(variable), nil
}
