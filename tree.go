package golastic

import (
	"github.com/cone/golastic/parser"
)

var esTree = &parser.Node{
	Name: "Operators",
	Nodes: []*parser.Node{
		&parser.Node{
			Name: "or",
			Apply: func(p *parser.Parser, o interface{}) {
			},
		},
		&parser.Node{
			Name: "and",
			Apply: func(p *parser.Parser, o interface{}) {
			},
		},
		&parser.Node{
			Name: "eq",
			Apply: func(p *parser.Parser, o interface{}) {

			},
		},
		&parser.Node{
			Name: "in",
			Apply: func(p *parser.Parser, o interface{}) {
			},
		},
		&parser.Node{
			Name: "matches",
			Apply: func(p *parser.Parser, o interface{}) {
			},
		},
		&parser.Node{
			Name: "cont",
			Apply: func(p *parser.Parser, o interface{}) {
			},
			IsOperator: true,
			Nodes: []*parser.Node{
				&parser.Node{
					Name: "any",
					Apply: func(p *parser.Parser, o interface{}) {
					},
				},
			},
		},
		&parser.Node{
			Name: "lt",
			Apply: func(p *parser.Parser, o interface{}) {
			},
		},
		&parser.Node{
			Name: "lteq",
			Apply: func(p *parser.Parser, o interface{}) {
			},
		},
		&parser.Node{
			Name: "gt",
			Apply: func(p *parser.Parser, o interface{}) {
			},
		},
		&parser.Node{
			Name: "gteq",
			Apply: func(p *parser.Parser, o interface{}) {
			},
		},
		&parser.Node{
			Name: "start",
			Apply: func(p *parser.Parser, o interface{}) {
			},
		},
		&parser.Node{
			Name: "end",
			Apply: func(p *parser.Parser, o interface{}) {
			},
		},
		&parser.Node{
			Name: "true",
			Apply: func(p *parser.Parser, o interface{}) {
			},
		},
		&parser.Node{
			Name: "false",
			Apply: func(p *parser.Parser, o interface{}) {
			},
		},
		&parser.Node{
			Name: "present",
			Apply: func(p *parser.Parser, o interface{}) {
			},
		},
		&parser.Node{
			Name: "blank",
			Apply: func(p *parser.Parser, o interface{}) {
			},
		},
		&parser.Node{
			Name: "null",
			Apply: func(p *parser.Parser, o interface{}) {
			},
		},
		&parser.Node{
			Name: "not",
			Nodes: []*parser.Node{
				&parser.Node{
					Name: "eq",
					Apply: func(p *parser.Parser, o interface{}) {
					},
				},
				&parser.Node{
					Name: "in",
					Apply: func(p *parser.Parser, o interface{}) {
					},
				},
				&parser.Node{
					Name: "cont",
					Apply: func(p *parser.Parser, o interface{}) {
					},
					IsOperator: true,
					Nodes: []*parser.Node{
						&parser.Node{
							Name: "any",
							Apply: func(p *parser.Parser, o interface{}) {
							},
						},
					},
				},
				&parser.Node{
					Name: "start",
					Apply: func(p *parser.Parser, o interface{}) {
					},
				},
				&parser.Node{
					Name: "end",
					Apply: func(p *parser.Parser, o interface{}) {
					},
				},
				&parser.Node{
					Name: "true",
					Apply: func(p *parser.Parser, o interface{}) {
					},
				},
				&parser.Node{
					Name: "false",
					Apply: func(p *parser.Parser, o interface{}) {
					},
				},
				&parser.Node{
					Name: "null",
					Apply: func(p *parser.Parser, o interface{}) {
					},
				},
			},
		},
		&parser.Node{
			Name: "does",
			Nodes: []*parser.Node{
				&parser.Node{
					Name: "not",
					Nodes: []*parser.Node{
						&parser.Node{
							Name: "match",
							Apply: func(p *parser.Parser, o interface{}) {
							},
						},
					},
				},
			},
		},
	},
}
