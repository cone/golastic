package golastic

import (
	"bytes"
	"code.google.com/p/go-uuid/uuid"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

const (
	indexer_tmpl  = `{"%s":{"_id":"%s"}}`
	INDEX_ACTION  = "index"
	DELETE_ACTION = "delete"
	UPDATE_ACTION = "update"
)

func NewApi(serverUrl string) *Api {
	return &Api{
		ServerUrl: serverUrl,
	}
}

type Api struct {
	ServerUrl string
	index     string
	_type     string
	errors    []error
}

func (this *Api) From(index, t string) *Api {
	this.index = index
	this._type = t

	return this
}

func (this *Api) Bulk(action string, param interface{}) []byte {
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

func (this *Api) processBulk(action string, v reflect.Value, param interface{}) []byte {
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

func (this *Api) processItem(action string, param interface{}) ([]byte, error) {
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

func (this *Api) getUrl() string {
	var url string

	if strings.Trim(this.index, " ") != "" {
		url = this.ServerUrl + "/" + this.index
	}

	if strings.Trim(this._type, " ") != "" {
		url = url + "/" + this._type
	}

	return url
}
