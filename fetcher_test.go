package golastic

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetcher_Find(t *testing.T) {
	jsonStr := `{
		"_index" : "test",
		"_type" : "products",
		"_id" : "1",
		"_version" : 1,
		"found" : true,
		"_source":{"id":16,"name":"Spree Mug"}
	}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, jsonStr)
	}))
	defer ts.Close()

	fetcher := NewFetcher(ts.URL)

	req, err := http.NewRequest("GET", ts.URL, nil)
	if err != nil {
		t.Error("An error has ocurred: " + err.Error())
		t.Error(err)
	}

	resBytes, err := fetcher.getResponse(req)
	if err != nil {
		t.Error("An error has ocurred: " + err.Error())
		t.Error(err)
	}

	result := &ResultItem{}
	err = json.Unmarshal(resBytes, result)
	if err != nil {
		t.Error("An error has ocurred: " + err.Error())
	}

	AssertEqualString(t, result.Index, "test")
	AssertEqualString(t, result.Type, "products")
	AssertEqualString(t, result.Id, "1")
	AssertEqualInt(t, result.Version, 1)
	AssertEqualBool(t, result.Found, true)
}
