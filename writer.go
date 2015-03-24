package golastic

import (
	"bytes"
	"code.google.com/p/go-uuid/uuid"
	"encoding/json"
	"fmt"
	"reflect"
)

const (
	indexer_tmpl  = `{"%s":{"_id":"%s"}}`
	INDEX_ACTION  = "index"
	DELETE_ACTION = "delete"
	UPDATE_ACTION = "update"
)

func NewWriter(requester Requester) *Writer {
	return &Writer{
		requester: requester,
	}
}

type Writer struct {
	requester Requester
	url       string
	errors    []error
}

func (this *Writer) Index(param interface{}) ([]byte, error) {
	paramBytes, err := json.Marshal(param)
	if err != nil {
		return []byte{}, err
	}

	return this.requester.Put(this.url, paramBytes)
}

func (this *Writer) Update(id string, param interface{}) ([]byte, error) {
	paramBytes, err := json.Marshal(param)
	if err != nil {
		return []byte{}, err
	}

	urlStr := this.url + "/" + id

	return this.requester.Put(urlStr, paramBytes)
}

func (this *Writer) Delete(id string) ([]byte, error) {
	var urlStr string

	if id != "" {
		urlStr = this.url + "/" + id
	} else {
		urlStr = this.url
	}

	return this.requester.Delete(urlStr)
}

func (this *Writer) Bulk(action string, param interface{}) []byte {
	v := reflect.ValueOf(param)
	k := v.Kind()
	var b []byte
	var err error
	this.errors = []error{}

	if k == reflect.Array || k == reflect.Slice {
		b = this.processBulk(action, v, param)
	} else {
		b, err = this.processItem(action, param)
		if err != nil {
			this.errors = append(this.errors, err)
		}
	}

	return b
}

func (this *Writer) processBulk(action string, v reflect.Value, param interface{}) []byte {
	buffer := bytes.Buffer{}

	for i := 0; i < v.Len(); i++ {
		pBytes, err := this.processItem(action, v.Index(i).Interface())
		if err != nil {
			this.errors = append(this.errors, err)
			continue
		}
		buffer.Write(pBytes)
	}

	return buffer.Bytes()
}

func (this *Writer) processItem(action string, param interface{}) ([]byte, error) {
	id := uuid.New()

	buffer := bytes.Buffer{}
	buffer.WriteString(fmt.Sprintf(indexer_tmpl, action, id))

	paramBytes, err := json.Marshal(param)
	if err != nil {
		return []byte{}, err
	}

	buffer.WriteRune('\n')
	buffer.WriteString(string(paramBytes))
	buffer.WriteRune('\n')

	return buffer.Bytes(), nil
}
