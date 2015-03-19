package golastic

import (
	"testing"
)

func TestPostQuery_Match(t *testing.T) {

	got, err := NewPostQuery(MATCH_QUERY).Element("message", "this is a test").String()
	if err != nil {
		t.Errorf("An error has ocurred: " + err.Error())
		return
	}

	expected := `{"match":{"message":"this is a test"}}`

	if got != expected {
		t.Errorf("Mismatch.\ngot: %s\nexpecting:%s\n", got, expected)
		return
	}

	got, err = NewPostQuery(MATCH_QUERY).Element(
		"message",
		map[string]string{"query": "this is a test"},
	).String()
	if err != nil {
		t.Errorf("An error has ocurred: " + err.Error())
		return
	}

	expected = `{"match":{"message":{"query":"this is a test"}}}`

	if got != expected {
		t.Errorf("Mismatch.\ngot: %s\nexpecting:%s\n", got, expected)
		return
	}

}
