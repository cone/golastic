package golastic

import (
	"bytes"
	"testing"
)

func TestBulkDoc(t *testing.T) {
	b := Bulk()

	b.Index("1", TestProduct{1, "Name"})
	b.Delete("2", TestProduct{2, "Name2"})
	b.Add(UPDATE_ACTION, "", TestProduct{3, "Name3"})

	resBytes, err := b.Read(0, 3)
	if err != nil {
		Error(t, err)
	}

	buffer := bytes.Buffer{}
	buffer.WriteString(`{"index":{"_id":"1"}}`)
	buffer.WriteRune('\n')
	buffer.WriteString(`{"id":1,"name":"Name"}`)
	buffer.WriteRune('\n')
	buffer.WriteString(`{"delete":{"_id":"2"}}`)
	buffer.WriteRune('\n')
	buffer.WriteString(`{"id":2,"name":"Name2"}`)

	expected := string(buffer.Bytes())

	resStr := string(resBytes)

	AssertContainsSubstring(t, resStr, expected)
}
