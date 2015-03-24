package golastic

type Requester interface {
	Post(string, []byte) ([]byte, error)
	Get(string, string) ([]byte, error)
}
