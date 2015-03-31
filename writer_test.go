package golastic

import (
	"errors"
	"testing"
)

func TestWriter(t *testing.T) {
	requester := &FakeRequester{}

	b := Bulk().Index("", TestProduct{})

	errs := b.Errors()

	if len(errs) > 0 {
		Error(t, errs[0])
	}

	i := NewWriter(requester)

	errs = i.Bulk(b)
	if len(errs) > 0 {
		Error(t, errs[0])
	}

	requester.PostFunction =
		func(b []byte, err error) ([]byte, error) {
			fakeError := errors.New("Some error")

			return []byte{}, fakeError
		}

	b = Bulk().Delete("1", TestProduct{}).Delete("2", TestProduct{})

	errs = i.Bulk(b)
	if len(errs) < 1 {
		t.Error("There should be some errors")
	}
}
