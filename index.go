package golastic

type Index struct {
	tp string
}

func (this *Index) Type(t string) *Index {
	this.tp = t

	return this
}

func (this *Index) Query() {}
