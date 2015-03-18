package golastic

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func NewFetcher(serverUrl string) *Fetcher {
	return &Fetcher{
		ServerUrl: serverUrl,
	}
}

type Fetcher struct {
	ServerUrl string
	index     string
	_type     string
}

func (this *Fetcher) Index(index string) *Fetcher {
	this.index = index

	return this
}

func (this *Fetcher) Type(t string) *Fetcher {
	this._type = t

	return this
}

func (this *Fetcher) Query(query *PostQuery) (*Result, error) {
	queryBytes, err := query.Bytes()
	if err != nil {
		return nil, err
	}

	url := this.getUrl() + "/_search"

	responseBytes, err := this.doRequest("POST", url, bytes.NewBuffer(queryBytes))
	if err != nil {
		return nil, err
	}

	result := &Result{}

	err = json.Unmarshal(responseBytes, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (this *Fetcher) Find(id int) (*Result, error) {
	url := this.getUrl() + "/" + strconv.Itoa(id)

	responseBytes, err := this.doRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	result := &Result{}

	err = json.Unmarshal(responseBytes, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (this *Fetcher) doRequest(method, url string, buffer *bytes.Buffer) ([]byte, error) {
	req, err := http.NewRequest(method, url, buffer)
	if err != nil {
		return []byte{}, err
	}

	client := http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}

	resBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}

	if ErrorResult, thereIsError := this.checkIfError(resBytes); thereIsError {
		return []byte{}, errors.New(ErrorResult.Error)
	}

	return resBytes, nil
}

func (this *Fetcher) checkIfError(jsonBytes []byte) (*ErrorResult, bool) {
	errorResult := &ErrorResult{}

	err := json.Unmarshal(jsonBytes, errorResult)
	if err != nil {
		return nil, false
	}

	return errorResult, true
}

func (this *Fetcher) getUrl() string {
	var url string

	if strings.Trim(this.index, " ") != "" {
		url = this.ServerUrl + "/" + this.index
	}

	if strings.Trim(this._type, " ") != "" {
		url = url + "/" + this._type
	}

	return url
}
