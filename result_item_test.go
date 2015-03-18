package golastic

import (
	"encoding/json"
	"testing"
)

func TestResultItem(t *testing.T) {
	jsonBytes := []byte(`{
  	"_index" : "test",
  	"_type" : "products",
  	"_id" : "1",
  	"_version" : 1,
  	"found" : true,
  	"_source":{"id":16,"name":"Spree Mug"}
	}`)

	resultItem := &ResultItem{}

	err := json.Unmarshal(jsonBytes, resultItem)
	if err != nil {
		t.Error("An error has ocurred: " + err.Error())
	}

	AssertEqualString(t, resultItem.Index, "test")
	AssertEqualString(t, resultItem.Type, "products")
	AssertEqualString(t, resultItem.Id, "1")
	AssertEqualInt(t, resultItem.Version, 1)
	AssertEqualBool(t, resultItem.Found, true)
}
