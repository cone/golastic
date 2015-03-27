package golastic

import (
	"errors"
	"math/rand"
	"strconv"
	"testing"
)

func TestWriter(t *testing.T) {
	requester := &FakeRequester{}

	i := NewWriter(requester)

	params := []string{}

	for i := 0; i < 1000; i++ {
		params = append(params, strconv.Itoa(i))
	}

	errs := i.Bulk(INDEX_ACTION, params)
	if len(errs) > 0 {
		Error(t, errs[0])
	}

	requester.SetPostCallback(
		func(b []byte, err error) ([]byte, error) {
			n := rand.Intn(10)
			var fakeError error

			if n < 5 {
				fakeError = errors.New("Some error")
			}

			return []byte{}, fakeError
		},
	)

	errs = i.Bulk(INDEX_ACTION, params)
	if len(errs) < 1 {
		t.Error("There should be some errors")
	}
}

func TestWriter_getItem(t *testing.T) {
	requester := &FakeRequester{}

	i := NewWriter(requester)

	result, err := i.createItemJsonBytes(INDEX_ACTION, []string{"Hello", "World"})
	if err != nil {
		Error(t, err)
	}

	if string(result) == "" {
		t.Error("Result is empty")
	}
}
