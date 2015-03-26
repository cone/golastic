package golastic

import (
	"encoding/json"
)

func NewReader(requester Requester) *Reader {
	return &Reader{
		requester: requester,
	}
}

type Reader struct {
	requester Requester
	url       string
}

func (this *Reader) Exec(query *QueryData) (*Result, error) {
	queryBytes, err := json.Marshal(query)
	if err != nil {
		return nil, err
	}

	urlStr := this.url + "/_search"

	responseBytes, err := this.requester.Post(urlStr, queryBytes)
	if err != nil {
		return nil, err
	}

	result := &Result{}

	err = this.getResponse(responseBytes, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (this *Reader) Find(id string) (*ResultItem, error) {
	urlStr := this.url + "/" + id

	responseBytes, err := this.requester.Get(urlStr)
	if err != nil {
		return nil, err
	}

	result := &ResultItem{}

	return result, this.getResponse(responseBytes, result)
}

func (this *Reader) getResponse(responseBytes []byte, container interface{}) error {
	err := json.Unmarshal(responseBytes, container)
	if err != nil {
		return err
	}

	return nil
}
