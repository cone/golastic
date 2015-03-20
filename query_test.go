package golastic

import (
	"testing"
)

func TestQuery_MATCH(t *testing.T) {
	q := Query("match", Item().Put("message", "this is a test"))

	qBytes, err := q.Bytes()
	if err != nil {
		Error(t, err)
	}

	qStr := string(qBytes)

	expected := `{"match":{"message":"this is a test"}}`

	AssertEqualString(t, qStr, expected)

	q = Query("match", Item().Put(
		"message", Item().Put("query", "this is a test"),
	))

	qBytes, err = q.Bytes()
	if err != nil {
		Error(t, err)
	}

	qStr = string(qBytes)

	expected = `{"match":{"message":{"query":"this is a test"}}}`

	AssertEqualString(t, qStr, expected)
}
