package parser

type OperatorFunction func(re *Parser, o interface{})

type Node struct {
	Name       string
	Nodes      []*Node
	Apply      OperatorFunction
	IsOperator bool
}
