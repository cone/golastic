package golastic

const (
	ZERO_TERMS_QUERY_ALL = "all"
	TYPE_PHRASE_PREFIX   = "phrase_prefix"
)

type Params map[string]interface{}

func (this *Params) Analyzer(value string) *Params {
	m := *this
	m["analyzer"] = value

	return this
}

func (this *Params) CutoffFrequency(value float64) *Params {
	m := *this
	m["cutoff_frequency"] = value

	return this
}

func (this *Params) Query(value string) *Params {
	m := *this
	m["query"] = value

	return this
}

func (this *Params) Type(value string) *Params {
	m := *this
	m["type"] = value

	return this
}

func (this *Params) ZeroTermsQuery(value string) *Params {
	m := *this
	m["zero_terms_query"] = value

	return this
}

func (this *Params) Fields(fields ...string) *Params {
	m := *this
	m["fields"] = fields

	return this
}

func (this *Params) ToMap() *Params {
	return this
}
