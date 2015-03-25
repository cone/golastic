package golastic

import (
	"encoding/json"
	"testing"
)

func TestQuery_MATCH(t *testing.T) {
	q := Query("match").Set(
		Item().Put("message", "this is a test"),
	)

	qBytes, err := json.Marshal(q)
	if err != nil {
		Error(t, err)
	}

	qStr := string(qBytes)

	expected := `{"query":{"match":{"message":"this is a test"}}}`

	AssertEqualString(t, qStr, expected)

	q = Query("match").Set(
		Item().Put("message", Item().Put("query", "this is a test")),
	)

	qBytes, err = json.Marshal(q)
	if err != nil {
		Error(t, err)
	}

	qStr = string(qBytes)

	expected = `{"query":{"match":{"message":{"query":"this is a test"}}}}`

	AssertEqualString(t, qStr, expected)
}
