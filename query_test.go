package golastic

import (
	"testing"
)

func TestQuery_Match(t *testing.T) {
	q := &Query{}

	got, err := q.Match(QueryField{"message": "this is a test"}).String()
	if err != nil {
		t.Errorf("An error has ocurred: " + err.Error())
		return
	}

	expected := `{"match":{"message":"this is a test"}}`

	if got != expected {
		t.Errorf("Mismatch.\ngot: %s\nexpecting:%s\n", got, expected)
		return
	}

	got, err = q.Match(
		QueryField{
			"message": &QueryParams{
				Query: "this is a test",
			},
		},
	).String()
	if err != nil {
		t.Errorf("An error has ocurred: " + err.Error())
		return
	}

	expected = `{"match":{"message":{"query":"this is a test"}}}`

	if got != expected {
		t.Errorf("Mismatch.\ngot: %s\nexpecting:%s\n", got, expected)
		return
	}

}

//func TestQuery_MultiMatch(t *testing.T) {
//q := &Query{}

//_, err := q.MultiMatch("this is a test", []string{"subject", "message"}).String()
//if err != nil {
//t.Errorf("An error has ocurred: " + err.Error())
//return
//}
//}
