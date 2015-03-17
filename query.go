package golastic

import (
	"encoding/json"
)

type Query map[string]map[string]interface{}

func (this *Query) Match(fields ...QueryField) *Query {
	this.addField("match", fields)

	return this
}

func (this *Query) MatchPhrase(fields ...QueryField) *Query {
	this.addField("match_phrase", fields)

	return this
}

func (this *Query) MatchPhrasePrefix(fields ...QueryField) *Query {
	this.addField("match_phrase_prefix", fields)

	return this
}

func (this *Query) addField(parent string, fields []QueryField) {
	m := *this
	m[parent] = map[string]interface{}{}

	for _, field := range fields {
		for key, value := range field {
			m[parent][key] = value
		}
	}
}

func (this *Query) MultiMatch(ops map[string]interface{}) *Query {
	m := *this

	m["multi_match"] = ops

	return this
}

func (this *Query) String() (string, error) {
	queryBytes, err := json.Marshal(this)
	if err != nil {
		return "", err
	}

	return string(queryBytes), nil
}
