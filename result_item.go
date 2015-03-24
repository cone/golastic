package golastic

type ResultItem struct {
	Index   string      `json:"_index"`
	Type    string      `json:"_type"`
	Id      string      `json:"_id"`
	Version int         `json:"_version"`
	Found   bool        `json:"found"`
	Source  interface{} `json:"_source"`
	Created bool        `json:"created"`
}
