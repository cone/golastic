package golastic

func NewField(name string) *Field {
	return &Field{
		Name: name,
	}
}

type Field struct {
	Name        string
	OutterQuery string
	ParamSet    *Params
}

func (this *Field) Query(query string) *Field {
	this.OutterQuery = query

	return this
}

func (this *Field) Params(params *Params) *Field {
	this.ParamSet = params

	return this
}

func (this *Field) ToMap() map[string]interface{} {
	m := map[string]interface{}{}

	if this.ParamSet == nil {
		m[this.Name] = this.OutterQuery
	} else {
		m[this.Name] = this.ParamSet.ToMap()
	}

	return m
}
