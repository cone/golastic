package golastic

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	//"reflect"
	//"sync"

	"code.google.com/p/go-uuid/uuid"
)

const (
	indexer_tmpl  = `{"%s":{"_id":"%s"}}`
	INDEX_ACTION  = "index"
	DELETE_ACTION = "delete"
	UPDATE_ACTION = "update"
)

func Bulk() *BulkDoc {
	return &BulkDoc{}
}

type BulkDoc struct {
	Data [][]byte
}

func (this *BulkDoc) Index(id string, param interface{}) error {
	return this.Add(INDEX_ACTION, id, param)
}

func (this *BulkDoc) Delete(id string, param interface{}) error {
	return this.Add(DELETE_ACTION, id, param)
}

func (this *BulkDoc) Update(id string, param interface{}) error {
	return this.Add(UPDATE_ACTION, id, param)
}

func (this *BulkDoc) Add(action string, id string, param interface{}) error {
	if id == "" {
		id = uuid.New()
	}

	buffer := bytes.Buffer{}
	buffer.WriteString(fmt.Sprintf(indexer_tmpl, action, id))

	paramBytes, err := json.Marshal(param)
	if err != nil {
		return err
	}

	buffer.WriteRune('\n')
	buffer.WriteString(string(paramBytes))
	buffer.WriteRune('\n')

	this.Data = append(this.Data, buffer.Bytes())

	return nil
}

func (this *BulkDoc) Len() int {
	return len(this.Data)
}

func (this *BulkDoc) Read(from, to int) ([]byte, error) {
	l := this.Len()

	if to > l {
		to = l
	}

	if from > to {
		return []byte{}, errors.New("Invalid bounds.")
	}

	buffer := bytes.Buffer{}

	for i := from; i < to; i++ {
		buffer.Write(this.Data[i])
	}

	return buffer.Bytes(), nil
}
