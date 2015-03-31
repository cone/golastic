package golastic

import (
	"fmt"
	"strconv"
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
}

func TestGolastic_Exec(t *testing.T) {
	golastic, err := New("http://localhost:9200")
	if err != nil {
		Error(t, err)
	}

	result, err := golastic.From("test", "product").Exec("", Query("match_all"))
	if err != nil {
		t.Error("An error has ocurred: " + err.Error())
	}

	totalHits := len(result.Hits.Hits)

	if totalHits < 5 {
		t.Errorf("Wrong number of hits: %d\n", totalHits)
	}

}

func TestGolastic_Bulk(t *testing.T) {
	golastic, err := New("http://localhost:9200")
	if err != nil {
		Error(t, err)
	}

	b := Bulk()

	for i := 1; i <= 10; i++ {
		b.Index(strconv.Itoa(i), TestProduct{i, fmt.Sprintf("Product %d", i)})
	}

	errs := golastic.From("test", "words").Bulk(b)
	if len(errs) > 0 {
		t.Error("An error has ocurred: " + errs[0].Error())
	}

	result, err := golastic.From("test", "product").Exec(POST_METHOD, Query("match_all"))
	if err != nil {
		t.Error("An error has ocurred: " + err.Error())
	}

	totalHits := len(result.Hits.Hits)

	if totalHits < 2 {
		t.Errorf("Wrong number of hits: %d\n", totalHits)
	}

	result, err = golastic.From("test", "words").Exec(DELETE_METHOD, Query("match_all"))
	if err != nil {
		t.Error("An error has ocurred: " + err.Error())
	}

	totalHits = len(result.Hits.Hits)

	if totalHits > 0 {
		t.Errorf("Wrong number of hits: %d\n", totalHits)
	}
}
