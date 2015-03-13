package parser

func newFakeResultGenerator() *Node {
	return &Node{
		Nodes: []*Node{
			&Node{
				Name: "test",
				Apply: func(p *Parser, o interface{}) {
					rg := o.(*string)

					field := p.getLastField()

					*rg = field + " is a test"
				},
			},
		},
	}
}
