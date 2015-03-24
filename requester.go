package golastic

type Requester interface {
	Post(string, []byte) ([]byte, error)
	Get(string) ([]byte, error)
	Put(string, []byte) ([]byte, error)
	Delete(string) ([]byte, error)
}
