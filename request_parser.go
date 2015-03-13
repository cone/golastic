package golastic

import (
	"net/http"

	"github.com/cone/golastic/parser"
)

func NewRequestParser() *RequestParser {
	return &RequestParser{
		parser: parser.NewParser(esTree),
	}
}

type RequestParser struct {
	parser *parser.Parser
}

func (this *RequestParser) GetQuery(r *http.Request) string {
	return ""
}
