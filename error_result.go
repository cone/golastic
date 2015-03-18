package golastic

type ErrorResult struct {
	Error  string `json:"error"`
	Status int    `json:"status"`
}
