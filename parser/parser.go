package parser

import (
	"strings"
)

const (
	SEPARATOR = "_"
)

func NewParser(root *Node) *Parser {
	return &Parser{
		separator: SEPARATOR,
		root:      root,
	}
}

type Parser struct {
	toEvaluate      []string
	evaluatedTokens []string
	separator       string
	pos             int
	param           *IncomingParam
	root            *Node
}

func (this *Parser) Parse(input string, param, result interface{}) {
	this.reset()

	this.tokenize(input)

	this.param = newIncomingParam(param)

	for this.pos = 0; this.pos < len(this.toEvaluate); this.pos++ {
		token := this.toEvaluate[this.pos]

		if node, isCandidate := this.isCandidateToOperator(token); isCandidate {
			if foundNode, found := this.find(node, this.pos); found {

				foundNode.Apply(this, result)

			} else {

				this.evaluated(token)

			}
		} else {

			this.evaluated(token)

		}
	}
}

func (this *Parser) reset() {
	this.toEvaluate = []string{}
	this.evaluatedTokens = []string{}
}

func (this *Parser) tokenize(input string) {
	this.toEvaluate = strings.Split(input, this.separator)
}

func (this *Parser) find(nodeParam *Node, pos int) (*Node, bool) {
	if pos >= len(this.toEvaluate) {
		return nil, false
	}

	next := this.toEvaluate[pos]

	if nodeParam.Name != next {
		return nil, false
	}

	if len(nodeParam.Nodes) > 0 {
		for _, node := range nodeParam.Nodes {
			if foundNode, found := this.find(node, pos+1); found {

				return foundNode, true

			}
		}

		//none of its children nodes matched, check if is itself an operator
		if nodeParam.IsOperator == true {
			this.pos = pos
			return nodeParam, true
		}

	} else {
		this.pos = pos
		return nodeParam, true
	}

	return nil, false
}

func (this *Parser) getLastField() string {
	field := strings.Join(this.evaluatedTokens, this.separator)
	this.evaluatedTokens = []string{}
	return field
}

func (this *Parser) evaluated(token string) {
	this.evaluatedTokens = append(this.evaluatedTokens, token)
}

func (this *Parser) isCandidateToOperator(item string) (*Node, bool) {
	for _, node := range this.root.Nodes {
		if node.Name == item {
			return node, true
		}
	}
	return &Node{}, false
}
