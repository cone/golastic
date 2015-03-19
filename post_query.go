package golastic

import (
	"encoding/json"
)

const (
	MATCH_QUERY         = "match"
	MATCH_PHRASE        = "match_phrase"
	MATCH_PHRASE_PREFIX = "match_phrase_prefix"
	MULTI_MATCH         = "multi_match"
)

func NewPostQuery(t string) *PostQuery {
	return &PostQuery{
		queryType: t,
		data:      map[string]interface{}{},
	}
}

type PostQuery struct {
	queryType string
	data      map[string]interface{}
}

func (this *PostQuery) Size(size int) *PostQuery {
	this.data["size"] = size

	return this
}

func (this *PostQuery) From(from int) *PostQuery {
	this.data["from"] = from

	return this
}

func (this *PostQuery) Type(t string) *PostQuery {
	this.data["type"] = t

	return this
}

func (this *PostQuery) Element(key string, value interface{}) *PostQuery {
	this.data[this.queryType] = map[string]interface{}{key: value}

	return this
}

func (this *PostQuery) Params(params *Params) *PostQuery {
	this.data[this.queryType] = params

	return this
}

func (this *PostQuery) Bytes() ([]byte, error) {
	return json.Marshal(this.data)
}

func (this *PostQuery) String() (string, error) {
	queryBytes, err := json.Marshal(this.data)
	if err != nil {
		return "", err
	}

	return string(queryBytes), nil
}
