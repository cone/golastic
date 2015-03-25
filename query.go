package golastic

import "encoding/json"

func Query(t string) *QueryData {
	return &QueryData{
		_type:      t,
		_outerData: map[string]interface{}{},
		_innerData: struct{}{},
	}
}

type QueryData struct {
	_type      string
	_outerData map[string]interface{}
	_innerData interface{}
}

func (this *QueryData) Size(size int) *QueryData {
	this._outerData["size"] = size

	return this
}

func (this *QueryData) From(from int) *QueryData {
	this._outerData["from"] = from

	return this
}

func (this *QueryData) Type(t string) *QueryData {
	this._outerData["type"] = t

	return this
}

func (this *QueryData) Set(value interface{}) *QueryData {
	this._innerData = value

	return this
}

func (this *QueryData) MarshalJSON() ([]byte, error) {
	this._outerData["query"] = map[string]interface{}{
		this._type: this._innerData,
	}

	return json.Marshal(this._outerData)
}
