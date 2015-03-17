package golastic

const (
	ZERO_TERMS_QUERY_ALL = "all"
	TYPE_PHRASE_PREFIX   = "phrase_prefix"
)

func NewParam() *Params {
	return &Params{
		data: map[string]interface{}{},
	}
}

type Params struct {
	data map[string]interface{}
}

func (this *Params) Analyzer(value string) *Params {
	this.data["analyzer"] = value

	return this
}

func (this *Params) CutoffFrequency(value float64) *Params {
	this.data["cutoff_frequency"] = value

	return this
}

func (this *Params) Query(value string) *Params {
	this.data["query"] = value

	return this
}

func (this *Params) Type(value string) *Params {
	this.data["type"] = value

	return this
}

func (this *Params) ZeroTermsQuery(value string) *Params {
	this.data["zero_terms_query"] = value

	return this
}

func (this *Params) ToMap() map[string]interface{} {
	return this.data
}
