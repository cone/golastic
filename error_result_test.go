package golastic

import (
	"encoding/json"
	"testing"
)

func TestErrorResult(t *testing.T) {
	errorResult := &ErrorResult{}

	jsonBytes := []byte(`{
		"error":"Error",
		"status":400
	}`)

	err := json.Unmarshal(jsonBytes, errorResult)
	if err != nil {
		t.Error("An error has ocurred: " + err.Error())
	}

	AssertEqualString(t, errorResult.Error, "Error")
	AssertEqualInt(t, errorResult.Status, 400)
}
