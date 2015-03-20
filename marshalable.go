package golastic

type Marshalable interface {
	Bytes() ([]byte, error)
}
