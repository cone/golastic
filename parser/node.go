package core

type OperatorFunction func(re *Parser)

type Node struct {
	Name       string
	Nodes      []*Node
	Apply      OperatorFunction
	IsOperator bool
}
