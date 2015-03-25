package golastic

import (
	"testing"
)

func TestGolastic_Find(t *testing.T) {
	golastic, err := New("http://localhost:9200")
	if err != nil {
		Error(t, err)
	}

	resultItem, err := golastic.From("test", "products").Find("1")
	if err != nil {
		t.Error("An error has ocurred: " + err.Error())
	}

	AssertEqualString(t, resultItem.Index, "test")
	AssertEqualString(t, resultItem.Type, "products")
	AssertEqualString(t, resultItem.Id, "1")
}

//func TestGolastic_Exec(t *testing.T) {
//golastic, err := New("http://localhost:9200")
//if err != nil {
//Error(t, err)
//}

//resultItem, err := golastic.Exec(Query("match"))
//if err != nil {
//t.Error("An error has ocurred: " + err.Error())
//}

//AssertEqualString(t, resultItem.Index, "test")
//AssertEqualString(t, resultItem.Type, "products")
//AssertEqualString(t, resultItem.Id, "1")
//}
