package golastic

import (
	"encoding/json"
	"testing"
)

func TestResultItem(t *testing.T) {
	resultItem, err := getResultItem()
	if err != nil {
		Error(t, err)
	}

	AssertEqualString(t, resultItem.Index, "test")
	AssertEqualString(t, resultItem.Type, "products")
	AssertEqualString(t, resultItem.Id, "1")
	AssertEqualInt(t, resultItem.Version, 1)
	AssertEqualBool(t, resultItem.Found, true)
}

func TestResultItem_Scan(t *testing.T) {
	resultItem, err := getResultItem()
	if err != nil {
		Error(t, err)
	}

	product := &TestProduct{}

	resultItem.Scan(&product)

	AssertEqualInt(t, product.Id, 16)
	AssertEqualString(t, product.Name, "Spree Mug")
}

func getResultItem() (*ResultItem, error) {
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
		return nil, err
	}

	return resultItem, nil
}
