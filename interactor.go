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
	return this.withBody("POST", url, body)
}

func (this *Interactor) Put(url string, body []byte) ([]byte, error) {
	return this.withBody("PUT", url, body)
}

func (this *Interactor) DeleteWithBody(url string, body []byte) ([]byte, error) {
	return this.withBody("DELETE", url, body)
}

func (this *Interactor) withBody(method, url string, body []byte) ([]byte, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	return this.getResponse(req)
}

func (this *Interactor) Get(url string) ([]byte, error) {
	return this.withoutBody("GET", url)
}

func (this *Interactor) Delete(url string) ([]byte, error) {
	return this.withoutBody("DELETE", url)
}

func (this *Interactor) withoutBody(method, url string) ([]byte, error) {
	req, err := http.NewRequest(method, url, nil)
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
