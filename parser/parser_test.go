package parser

import (
	"fmt"
	"testing"
)

func testParser(t *testing.T, parser *Parser, correctTree *Tree) {
	tree, err := parser.Parse()
	if err != nil {
		t.Fatal(err)
	}
	url1, err := tree.GraphUrl()
	if err != nil {
		t.Fatal(err)
	}
	url2, err := correctTree.GraphUrl()
	if err != nil {
		t.Fatal(err)
	}
	if !tree.Equals(correctTree) {
		t.Fatalf("Invalid tree: %s\nMust be equal: %s", url1.String(), url2.String())
	}
	t.Log(url1.String())
}

var funcName = "functionName"
var returnTypeName = "returnTypeName"
var argumentName = "argName"
var argumentType = "argType"

func TestNoArgsNoReturnType(t *testing.T) {
	parser := NewParser(fmt.Sprintf("fun %s()", funcName))
	correctTree := NewStart(
		NewDeclaration(NewFunctionName(funcName), EmptyArguments, EmptyEnding),
	)
	testParser(t, parser, correctTree)
}

func TestNoArgsWithReturnType(t *testing.T) {
	parser := NewParser(fmt.Sprintf("fun %s(): %s", funcName, returnTypeName))
	correctTree := NewStart(
		NewDeclaration(
			NewFunctionName(funcName),
			EmptyArguments,
			NewEnding(NewType(returnTypeName)),
		),
	)
	testParser(t, parser, correctTree)
}

func TestOneArgWithReturnType(t *testing.T) {
	parser := NewParser(fmt.Sprintf(
		"fun %s(%s:%s): %s",
		funcName, argumentName, argumentType, returnTypeName))
	correctTree := NewStart(
		NewDeclaration(
			NewFunctionName(funcName),
			NewArguments(
				NewVariableAndType(NewVariable(argumentName), NewType(argumentType)),
				EmptyVariableAndTypeContinuation,
			),
			NewEnding(NewType(returnTypeName)),
		),
	)
	testParser(t, parser, correctTree)
}

func TestTwoArgumentsWithReturnType(t *testing.T) {
	parser := NewParser(fmt.Sprintf(
		"fun %s(%s:%s, %s:%s): %s",
		funcName, argumentName, argumentType, argumentName, argumentType, returnTypeName))
	correctTree := NewStart(
		NewDeclaration(
			NewFunctionName(funcName),
			NewArguments(
				NewVariableAndType(NewVariable(argumentName), NewType(argumentType)),
				NewVariableAndTypeContinuation(
					NewVariableAndType(NewVariable(argumentName), NewType(argumentType)),
					EmptyVariableAndTypeContinuation,
				),
			),
			NewEnding(NewType(returnTypeName)),
		),
	)
	testParser(t, parser, correctTree)
}

func TestThreeArgumentsWithReturnType(t *testing.T) {
	parser := NewParser(fmt.Sprintf(
		"fun %s(%s:%s, %s:%s, %s:%s): %s",
		funcName,
		argumentName, argumentType,
		argumentName, argumentType,
		argumentName, argumentType,
		returnTypeName))
	correctTree := NewStart(
		NewDeclaration(
			NewFunctionName(funcName),
			NewArguments(
				NewVariableAndType(NewVariable(argumentName), NewType(argumentType)),
				NewVariableAndTypeContinuation(
					NewVariableAndType(NewVariable(argumentName), NewType(argumentType)),
					NewVariableAndTypeContinuation(
						NewVariableAndType(NewVariable(argumentName), NewType(argumentType)),
						EmptyVariableAndTypeContinuation,
					),
				),
			),
			NewEnding(NewType(returnTypeName)),
		),
	)
	testParser(t, parser, correctTree)
}

func testExpectFailure(t *testing.T, parser *Parser) {
	tree, err := parser.Parse()
	if err == nil {
		url, err := tree.GraphUrl()
		if err != nil {
			t.Fatal(err)
		}
		t.Fatal(fmt.Sprintf("Expected failure, got: %s", url.String()))
	} else {
		t.Log(err.Error())
	}
}

func TestEmpty(t *testing.T) {
	parser := NewParser("")
	testExpectFailure(t, parser)
}

func TestOnlyFun(t *testing.T) {
	parser := NewParser("fun")
	testExpectFailure(t, parser)
}

func TestFunAndBracket(t *testing.T) {
	parser := NewParser("fun (")
	testExpectFailure(t, parser)
}

func TestNoArgumentNameAndType(t *testing.T) {
	parser := NewParser("fun (:)")
	testExpectFailure(t, parser)
}

func TestNoArgumentName(t *testing.T) {
	parser := NewParser("fun (:type)")
	testExpectFailure(t, parser)
}

func TestNoArgumentType(t *testing.T) {
	parser := NewParser("fun (arg)")
	testExpectFailure(t, parser)
}

func TestNoArgumentTypeWithColon(t *testing.T) {
	parser := NewParser("fun (arg:)")
	testExpectFailure(t, parser)
}

func TestBadEnding(t *testing.T) {
	parser := NewParser("fun name():")
	testExpectFailure(t, parser)
}
