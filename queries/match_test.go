package queries

import (
	"testing"
)

func TestMatch(t *testing.T) {
	m := Match()

	dm := *m.Query

	m.Put("message", "this is a test")

	mBytes, err := m.Bytes()
	if err != nil {
		t.Error("An error has ocurred: " + err.Error())
		return
	}

	mStr := string(mBytes)

	m.Size(10)
	m.From(5)
	m.Type("type")

	if mStr != `{"match":{"message":"this is a test"}}` {
		t.Error(mStr)
	}

	if dm["size"] != 10 {
		t.Error("Mismatch")
	}

	if dm["from"] != 5 {
		t.Error("Mismatch")
	}

	if dm["type"] != "type" {
		t.Error("Mismatch")
	}

}
