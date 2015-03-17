package golastic

import (
	"encoding/json"
)

const (
	MATCH_QUERY         = "match"
	MATCH_PHRASE        = "match_phrase"
	MATCH_PHRASE_PREFIX = "match_phrase_prefix"
)

func NewQuery(t string) *Query {
	return &Query{
		queryType: t,
		data:      map[string]interface{}{},
	}
}

type Query struct {
	queryType string
	data      map[string]interface{}
}

func (this *Query) Size(size int) *Query {
	this.data["size"] = size

	return this
}

func (this *Query) From(from int) *Query {
	this.data["from"] = from

	return this
}

func (this *Query) Type(t string) *Query {
	this.data["type"] = t

	return this
}

func (this *Query) Fields(fields ...*Field) *Query {
	for _, field := range fields {
		this.data[this.queryType] = field.ToMap()
	}

	return this
}

func (this *Query) String() (string, error) {
	queryBytes, err := json.Marshal(this.data)
	if err != nil {
		return "", err
	}

	return string(queryBytes), nil
}
