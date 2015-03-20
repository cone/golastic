package queries

type Query map[string]interface{}

func (this *Query) Size(size int) *Query {
	this.Put("size", size)
	return this
}

func (this *Query) From(from int) *Query {
	this.Put("from", from)

	return this
}

func (this *Query) Type(t string) *Query {
	this.Put("type", t)

	return this
}

func (this *Query) Put(key string, value interface{}) *Query {
	m := *this
	m[key] = value

	return this
}
