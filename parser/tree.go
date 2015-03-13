package core

var Tree = &Node{
	Name: "Operators",
	Nodes: []*Node{
		&Node{
			Name: "or",
			Apply: func(re *Parser) {
			},
		},
		&Node{
			Name: "and",
			Apply: func(re *Parser) {
			},
		},
		&Node{
			Name: "eq",
			Apply: func(re *Parser) {
			},
		},
		&Node{
			Name: "in",
			Apply: func(re *Parser) {
			},
		},
		&Node{
			Name: "matches",
			Apply: func(re *Parser) {
			},
		},
		&Node{
			Name: "cont",
			Apply: func(re *Parser) {
			},
			IsOperator: true,
			Nodes: []*Node{
				&Node{
					Name: "any",
					Apply: func(re *Parser) {
					},
				},
			},
		},
		&Node{
			Name: "lt",
			Apply: func(re *Parser) {
			},
		},
		&Node{
			Name: "lteq",
			Apply: func(re *Parser) {
			},
		},
		&Node{
			Name: "gt",
			Apply: func(re *Parser) {
			},
		},
		&Node{
			Name: "gteq",
			Apply: func(re *Parser) {
			},
		},
		&Node{
			Name: "start",
			Apply: func(re *Parser) {
			},
		},
		&Node{
			Name: "end",
			Apply: func(re *Parser) {
			},
		},
		&Node{
			Name: "true",
			Apply: func(re *Parser) {
			},
		},
		&Node{
			Name: "false",
			Apply: func(re *Parser) {
			},
		},
		&Node{
			Name: "present",
			Apply: func(re *Parser) {
			},
		},
		&Node{
			Name: "blank",
			Apply: func(re *Parser) {
			},
		},
		&Node{
			Name: "null",
			Apply: func(re *Parser) {
			},
		},
		&Node{
			Name: "not",
			Nodes: []*Node{
				&Node{
					Name: "eq",
					Apply: func(re *Parser) {
					},
				},
				&Node{
					Name: "in",
					Apply: func(re *Parser) {
					},
				},
				&Node{
					Name: "cont",
					Apply: func(re *Parser) {
					},
					IsOperator: true,
					Nodes: []*Node{
						&Node{
							Name: "any",
							Apply: func(re *Parser) {
							},
						},
					},
				},
				&Node{
					Name: "start",
					Apply: func(re *Parser) {
					},
				},
				&Node{
					Name: "end",
					Apply: func(re *Parser) {
					},
				},
				&Node{
					Name: "true",
					Apply: func(re *Parser) {
					},
				},
				&Node{
					Name: "false",
					Apply: func(re *Parser) {
					},
				},
				&Node{
					Name: "null",
					Apply: func(re *Parser) {
					},
				},
			},
		},
		&Node{
			Name: "does",
			Nodes: []*Node{
				&Node{
					Name: "not",
					Nodes: []*Node{
						&Node{
							Name: "match",
							Apply: func(re *Parser) {
							},
						},
					},
				},
			},
		},
	},
}
