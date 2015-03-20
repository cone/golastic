package golastic

import (
	"encoding/json"
)

func Item() *ItemData {
	return &ItemData{}
}

type ItemData map[string]interface{}

func (this *ItemData) Put(key string, value interface{}) *ItemData {
	m := *this
	m[key] = value

	return this
}

func (this *ItemData) Bytes() ([]byte, error) {
	return json.Marshal(this)
}
