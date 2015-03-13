package golastic

func NewQuery() *Query {
	return &Query{
		outgoing: map[string]interface{}{
			"query": struct{}{},
		},
	}
}

type Query struct {
	outgoing map[string]interface{}
}

func (this *Query) All() {
	this.outgoing["query"] = map[string]interface{}{
		"match_all": struct{}{},
	}
}

func (this *Query) Size(size int) *Query {
	this.outgoing["size"] = size

	return this
}

func (this *Query) From(from int) *Query {
	this.outgoing["from"] = from

	return this
}

func (this *Query) Source(fields ...string) *Query {
	this.outgoing["_source"] = fields

	return this
}

func (this *Query) Sort(fields ...Field) *Query {

	return this
}

type Field struct {
	Name  string
	Order string
	Mode  string
}
