package golastic

const (
	ZERO_TERMS_QUERY_ALL = "all"
	TYPE_PHRASE_PREFIX   = "phrase_prefix"
)

func NewParam() *FieldParams {
	return &FieldParams{
		data: map[string]interface{}{},
	}
}

type FieldParams struct {
	data map[string]interface{}
}

func (this *FieldParams) Analyzer(value string) *FieldParams {
	this.data["analyzer"] = value

	return this
}

func (this *FieldParams) CutoffFrequency(value float64) *FieldParams {
	this.data["cutoff_frequency"] = value

	return this
}

func (this *FieldParams) Query(value string) *FieldParams {
	this.data["query"] = value

	return this
}

func (this *FieldParams) Type(value string) *FieldParams {
	this.data["type"] = value

	return this
}

func (this *FieldParams) ZeroTermsQuery(value string) *FieldParams {
	this.data["zero_terms_query"] = value

	return this
}

func (this *FieldParams) ToMap() map[string]interface{} {
	return this.data
}
