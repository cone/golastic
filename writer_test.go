package golastic

import (
	"testing"
)

func TestWriter(t *testing.T) {
	requester := &FakeRequester{}

	i := NewWriter(requester)

	params := []interface{}{
		[]string{"Hello"},
		"World",
	}

	pBytes := i.Bulk(INDEX_ACTION, params)

	errs := i.errors
	if len(errs) > 0 {
		Error(t, errs[0])
	}

	if string(pBytes) == "" {
		t.Error("Result is empty")
	}

}

func TestWriter_getItem(t *testing.T) {
	requester := &FakeRequester{}

	i := NewWriter(requester)

	result, err := i.processItem(INDEX_ACTION, []string{"Hello", "World"})
	if err != nil {
		Error(t, err)
	}

	if string(result) == "" {
		t.Error("Result is empty")
	}
}
