package golastic

//import (
//"net/http"
//)

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

//func (this *Fetcher) Result(query *PostQuery) (interface{}, error) {
//queryBytes, err := query.Bytes()
//if err != nil {
//return struct{}{}, err
//}

//return struct{}{}, nil
//}
