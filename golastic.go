package golastic

import (
	"net/url"
)

func New(serverUrl string) (*Golastic, error) {
	baseUrl, err := url.Parse(serverUrl)
	if err != nil {
		return nil, err
	}

	i := NewInteractor()

	return &Golastic{
		Reader:  NewReader(i),
		BaseUrl: baseUrl,
	}, nil
}

type Golastic struct {
	BaseUrl *url.URL
	*Reader
	interactor *Interactor
	index      string
}

func (this *Golastic) From(index, t string) *Golastic {
	this.completeUrl(index + "/" + t)

	return this
}

func (this *Golastic) Index(index string) *Golastic {
	this.index = index
	this.completeUrl(index)

	return this
}

func (this *Golastic) Type(t string) *Golastic {
	this.completeUrl(this.index + "/" + t)

	return this
}

func (this *Golastic) completeUrl(urlTail string) {
	this.Reader.url = this.BaseUrl.String() + "/" + urlTail
}
