package golastic

import (
	//"bytes"
	//"net/http"
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
	query     *PostQuery
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

func (this *Fetcher) Query(query *PostQuery) *Fetcher {
	this.query = query

	return this
}

//func (this *Fetcher) Find(id int) (interface{}, error) {

//}

//func (this *Fetcher) Result(query *PostQuery) (interface{}, error) {
//queryBytes, err := query.Bytes()
//if err != nil {
//return struct{}{}, err
//}

//req, err := http.NewRequest("POST", this.getUrl(), bytes.NewBuffer(queryBytes))
//if err != nil {
//return struct{}{}, err
//}

//return struct{}{}, nil
//}

func (this *Fetcher) getUrl() string {
	url := this.ServerUrl + "/" + this.index

	if strings.Trim(this._type, " ") != "" {
		url = url + "/" + this._type
	}

	return url
}
