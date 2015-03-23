package golastic

import (
	"testing"
)

func TestApi(t *testing.T) {
	i := NewApi("")

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

func TestApi_getItem(t *testing.T) {
	i := NewApi("")

	result, err := i.processItem(INDEX_ACTION, []string{"Hello", "World"})
	if err != nil {
		Error(t, err)
	}

	if string(result) == "" {
		t.Error("Result is empty")
	}
}
