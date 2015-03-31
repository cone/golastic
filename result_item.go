package golastic

import "encoding/json"

type ResultItem struct {
	Index   string      `json:"_index"`
	Type    string      `json:"_type"`
	Id      string      `json:"_id"`
	Version int         `json:"_version"`
	Found   bool        `json:"found"`
	Source  interface{} `json:"_source"`
	Created bool        `json:"created"`
	Status  int         `json:"status"`
	Error   string      `json:"error"`
}

func (this *ResultItem) Scan(i interface{}) error {
	sourceBytes, err := json.Marshal(this.Source)
	if err != nil {
		return err
	}

	return json.Unmarshal(sourceBytes, i)
}
