package golastic

func Connect(serverUrl string) *Golastic {
	return &Golastic{
		fetcher: NewFetcher(serverUrl),
	}
}

type Golastic struct {
	fetcher *Fetcher
}

func (this *Golastic) Fetcher() *Fetcher {
	return this.fetcher
}
