package golastic

import "encoding/json"

func Filter(t string) *FilterData {
	return &FilterData{
		_type: t,
		_data: map[string]map[string]interface{}{
			"filter": map[string]interface{}{
				t: struct{}{},
			},
		},
	}
}

type FilterData struct {
	_type string
	_data map[string]map[string]interface{}
}

func (this *FilterData) Set(value interface{}) *FilterData {
	this._data["filter"][this._type] = value

	return this
}

func (this *FilterData) MarshalJSON() ([]byte, error) {
	return json.Marshal(this._data)
}
