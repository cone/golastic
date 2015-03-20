package queries

const (
	ZERO_TERMS_QUERY_ALL = "all"
	TYPE_PHRASE_PREFIX   = "phrase_prefix"
)

func MatchParams() *MatchQueryParams {
	return &MatchQueryParams{}
}

type MatchQueryParams map[string]interface{}

func (this *MatchQueryParams) Analyzer(value string) *MatchQueryParams {
	this.Put("analyzer", value)

	return this
}

func (this *MatchQueryParams) CutoffFrequency(value float64) *MatchQueryParams {
	this.Put("cutoff_frequency", value)

	return this
}

func (this *MatchQueryParams) Query(value string) *MatchQueryParams {
	this.Put("query", value)

	return this
}

func (this *MatchQueryParams) Type(value string) *MatchQueryParams {
	this.Put("type", value)

	return this
}

func (this *MatchQueryParams) ZeroTermsQuery(value string) *MatchQueryParams {
	this.Put("zero_terms_query", value)

	return this
}

func (this *MatchQueryParams) Put(key string, value interface{}) *MatchQueryParams {
	m := *this
	m[key] = value

	return this
}
