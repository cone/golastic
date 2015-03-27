package golastic

import "testing"

func AssertEqualString(t *testing.T, a, b string) {
	if a != b {
		mismatchError(t, a, b)
	}
}

func AssertEqualFloat(t *testing.T, a, b float64) {
	if a != b {
		mismatchError(t, a, b)
	}
}

func AssertEqualInt(t *testing.T, a, b int) {
	if a != b {
		mismatchError(t, a, b)
	}
}

func AssertEqualBool(t *testing.T, a, b bool) {
	if a != b {
		mismatchError(t, a, b)
	}
}

func Error(t *testing.T, err error) {
	t.Error("An error has ocurred: " + err.Error())
}

func mismatchError(t *testing.T, a, b interface{}) {
	t.Errorf("Mismatch.\n%s is not equal to %s\n", a, b)
}

type FakeRequesterCallback func([]byte, error) ([]byte, error)

type FakeRequester struct {
	GetResponse  []byte
	PostResponse []byte
	GetError     error
	PostError    error
	PutFunction  FakeRequesterCallback
}

func (this *FakeRequester) SetPostCallback(function FakeRequesterCallback) {
	this.PutFunction = function
}

func (this *FakeRequester) Post(url string, b []byte) ([]byte, error) {
	if this.PostError != nil {
		return []byte{}, this.PostError
	}

	return this.PostResponse, nil
}

func (this *FakeRequester) Get(url string) ([]byte, error) {
	if this.GetError != nil {
		return []byte{}, this.GetError
	}

	return this.GetResponse, nil
}

func (this *FakeRequester) Put(url string, b []byte) ([]byte, error) {
	if this.PutFunction != nil {
		return this.PutFunction(b, nil)
	}

	return []byte{}, nil
}

func (this *FakeRequester) Delete(url string) ([]byte, error) {
	return []byte{}, nil
}

type TestProduct struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
