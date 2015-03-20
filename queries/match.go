package queries

import (
	"encoding/json"
)

func Match() *MatchQuery {
	query := getMatchQuery()
	query._type = "match"

	return query
}

func MatchPhrase() *MatchQuery {
	query := getMatchQuery()
	query._type = "match_phrase"

	return query
}

func MatchPhrasePrefix() *MatchQuery {
	query := getMatchQuery()
	query._type = "match_phrase_prefix"

	return query
}

func getMatchQuery() *MatchQuery {
	return &MatchQuery{
		Query: &Query{},
		_data: map[string]interface{}{},
	}
}

type MatchQuery struct {
	_type string
	_data map[string]interface{}
	*Query
}

func (this *MatchQuery) Put(key string, value interface{}) *MatchQuery {
	this._data[key] = value

	return this
}

func (this *MatchQuery) Bytes() ([]byte, error) {
	m := *this.Query
	m[this._type] = this._data

	return json.Marshal(m)
}
