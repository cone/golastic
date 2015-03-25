package golastic

import (
	"encoding/json"
	"testing"
)

func TestFilter(t *testing.T) {
	f := Filter("and").Set([]interface{}{
		Item().Put(
			"range",
			Item().Put(
				"postDate",
				Item().
					Put("from", "2010-03-01").
					Put("to", "2010-04-01"),
			),
		),
	})

	fBytes, err := json.Marshal(f)
	if err != nil {
		Error(t, err)
	}

	fStr := string(fBytes)

	expected := `{"filter":{"and":[{"range":{"postDate":{"from":"2010-03-01","to":"2010-04-01"}}}]}}`

	AssertEqualString(t, fStr, expected)
}
