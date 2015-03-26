package golastic

import (
	"testing"
)

func TestGolastic_Find(t *testing.T) {
	golastic, err := New("http://localhost:9200")
	if err != nil {
		Error(t, err)
	}

	resultItem, err := golastic.From("test", "product").Find("1")
	if err != nil {
		t.Error("An error has ocurred: " + err.Error())
	}

	AssertEqualString(t, resultItem.Index, "test")
	AssertEqualString(t, resultItem.Type, "product")
	AssertEqualString(t, resultItem.Id, "1")
	AssertEqualInt(t, resultItem.Version, 1)
	AssertEqualBool(t, resultItem.Found, true)
}

func TestGolastic_Exec(t *testing.T) {
	golastic, err := New("http://localhost:9200")
	if err != nil {
		Error(t, err)
	}

	result, err := golastic.From("test", "product").Exec(Query("match_all"))
	if err != nil {
		t.Error("An error has ocurred: " + err.Error())
	}

	if len(result.Hits.Hits) < 5 {
		t.Error("Wrong number of hits")
	}

}
