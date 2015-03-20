package golastic

import (
	"testing"
)

func TestGolastic(t *testing.T) {
	golastic := New("http://127.0.0.1:9200")

	resultItem, err := golastic.Fetcher().From("test", "products").Find(1)
	if err != nil {
		t.Error("An error has ocurred: " + err.Error())
	}

	AssertEqualString(t, resultItem.Index, "test")
	AssertEqualString(t, resultItem.Type, "products")
	AssertEqualString(t, resultItem.Id, "1")
	AssertEqualInt(t, resultItem.Version, 1)
	AssertEqualBool(t, resultItem.Found, true)
}
