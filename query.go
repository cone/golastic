package golastic

func Query(t string, item interface{}) *ItemData {
	return &ItemData{
		t: item,
	}
}

type QueryData ItemData

func (this *QueryData) Size(size int) *QueryData {
	this.Put("size", size)

	return this
}

func (this *QueryData) From(from int) *QueryData {
	this.Put("from", from)

	return this
}

func (this *QueryData) Type(t string) *QueryData {
	this.Put("type", t)

	return this
}

func (this *QueryData) Put(key string, value interface{}) *QueryData {
	m := *this
	m[key] = value

	return this
}
