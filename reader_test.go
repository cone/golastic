package golastic

import (
	"testing"
)

func TestFetcher_Find(t *testing.T) {
	jsonBytes := []byte(`{
		"_index" : "test",
		"_type" : "products",
		"_id" : "1",
		"_version" : 1,
		"found" : true,
		"_source":{"id":16,"name":"Spree Mug"}
	}`)

	requester := &FakeRequester{
		GetResponse: jsonBytes,
	}

	reader := NewReader(requester)

	result, err := reader.Find("test")
	if err != nil {
		Error(t, err)
	}

	AssertEqualString(t, result.Index, "test")
	AssertEqualString(t, result.Type, "products")
	AssertEqualString(t, result.Id, "1")
	AssertEqualInt(t, result.Version, 1)
	AssertEqualBool(t, result.Found, true)
}
