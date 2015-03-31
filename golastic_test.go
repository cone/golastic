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

	totalHits := len(result.Hits.Hits)

	if totalHits < 5 {
		t.Errorf("Wrong number of hits: %d\n", totalHits)
	}

}

//func TestGolastic_Bulk(t *testing.T) {
//golastic, err := New("http://localhost:9200")
//if err != nil {
//Error(t, err)
//}

//errs := golastic.From("test", "words").Bulk(INDEX_ACTION, []TestProduct{{1, "Hello"}})
//if len(errs) > 0 {
//t.Error("An error has ocurred: " + errs[0].Error())
//}

//result, err := golastic.From("test", "product").Exec(Query("match_all"))
//if err != nil {
//t.Error("An error has ocurred: " + err.Error())
//}

//totalHits := len(result.Hits.Hits)

//if totalHits < 2 {
//t.Errorf("Wrong number of hits: %d\n", totalHits)
//}

//errs = golastic.From("test", "words").Bulk(DELETE_ACTION, []TestProduct{{1, "Hello"}, {2, "World"}})
//if len(errs) > 0 {
//t.Error("An error has ocurred: " + errs[0].Error())
//}

//result, err = golastic.From("test", "words").Exec(Query("match_all"))
//if err != nil {
//t.Error("An error has ocurred: " + err.Error())
//}

//totalHits = len(result.Hits.Hits)

//if totalHits > 0 {
//t.Errorf("Wrong number of hits: %d\n", totalHits)
//}
//}
