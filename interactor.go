package golastic

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func NewInteractor() *Interactor {
	return &Interactor{}
}

type Interactor struct {
	Client http.Client
}

func (this *Interactor) Post(url string, body []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	return this.getResponse(req)
}

func (this *Interactor) Get(url string, id string) ([]byte, error) {
	url = url + "/" + id

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	return this.getResponse(req)
}

func (this *Interactor) getResponse(req *http.Request) ([]byte, error) {
	res, err := this.Client.Do(req)
	if err != nil {
		return []byte{}, err
	}

	resBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}

	err = this.checkIfError(resBytes)
	if err != nil {
		return []byte{}, err
	}

	return resBytes, nil
}

func (this *Interactor) checkIfError(jsonBytes []byte) error {
	errorResult := &ErrorResult{}

	err := json.Unmarshal(jsonBytes, errorResult)
	if err != nil {
		return nil
	}

	if errorResult.Error == "" {
		return nil
	}

	return errors.New(errorResult.Error)
}
