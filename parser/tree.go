package parser

import (
	"fmt"
	"net/url"
)

func (tree *Tree) Equals(to *Tree) bool {
	if tree.node != to.node || len(tree.children) != len(to.children) {
		return false
	}
	for i, leftChild := range tree.children {
		rightChild := to.children[i]
		if !leftChild.Equals(rightChild) {
			return false
		}
	}
	return true
}

func NewTree(node string, children ...*Tree) *Tree {
	return &Tree{
		node:     node,
		children: children,
	}
}

func NewStart(declaration *Declaration) *Start {
	return NewTree("Start", FunTree, declaration)
}

func NewDeclaration(functionName *FunctionName, arguments *Arguments, ending *Ending) *Declaration {
	return NewTree("Declaration", functionName, LParenTree, arguments, RParenTree, ending)
}

func NewArguments(varAndType *VariableAndType, continuation *VariableAndTypeContinuation) *Arguments {
	return NewTree("Arguments", varAndType, continuation)
}

var EmptyArguments = NewTree("Arguments")

func NewEnding(typeTree *Type) *Tree {
	return NewTree("Ending", ColonTree, typeTree)
}

var EmptyEnding = NewTree("Ending")

func NewVariableAndType(variable *Variable, variableType *Type) *Tree {
	return NewTree("VariableAndType", variable, ColonTree, variableType)
}

func NewVariableAndTypeContinuation(variableAndType *VariableAndType, continuation *VariableAndTypeContinuation) *Tree {
	return NewTree("VariableAndTypeContinuation", CommaTree, variableAndType, continuation)
}

var EmptyVariableAndTypeContinuation = NewTree("VariableAndTypeContinuation")

func NewVariable(name string) *Tree {
	return NewTree("Variable", NewTree(name))
}

func NewType(name *TypeName, mark *TypeMark) *Tree {
	return NewTree("Type", name, mark)
}

func NewTypeName(name string) *Tree {
	return NewTree("TypeName", NewTree(name))
}

var QuestionTypeMark = NewTree("TypeMark", NewTree("Question"))
var ExclamationTypeMark = NewTree("TypeMark", NewTree("Exclamation"))
var EmptyTypeMark = NewTree("TypeMark")

func NewFunctionName(name string) *Tree {
	return NewTree("FunctionName", NewTree(name))
}

var (
	FunTree    = NewTree("Fun")
	LParenTree = NewTree("LParen")
	RParenTree = NewTree("RParen")
	CommaTree  = NewTree("Comma")
	ColonTree  = NewTree("Colon")
)

func (tree *Tree) GraphUrl() (*url.URL, error) {
	return url.Parse("https://dreampuf.github.io/GraphvizOnline/#" + tree.ToGraphVizGraph())
}

func (tree *Tree) ToGraphVizGraph() string {
	return fmt.Sprintf("digraph G {\n%s\n}", tree.GraphVizRules("0"))
}

func (tree *Tree) GraphVizRules(nodeId string) string {
	res := formatNode(nodeId, tree.node)
	for i, child := range tree.children {
		childNodeId := nodeId + fmt.Sprint(i)
		res += formatEdge(nodeId, childNodeId)
		res += child.GraphVizRules(childNodeId)
	}
	return res
}

func formatEdge(from, to string) string {
	return fmt.Sprintf("\t%s -> %s;\n", from, to)
}

func formatNode(id, label string) string {
	return fmt.Sprintf("\t%s [label=%s]\n", id, label)
}
