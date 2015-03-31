package golastic

import (
	//"bytes"
	//"code.google.com/p/go-uuid/uuid"
	"encoding/json"
	//"errors"
	//"fmt"
	//"reflect"
	//"sync"
)

func NewWriter(requester Requester) *Writer {
	return &Writer{
		requester:   requester,
		chunkLength: 50,
	}
}

type Writer struct {
	requester   Requester
	url         string
	errors      []error
	chunkLength int
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

//func (this *Writer) Bulk(action string, param interface{}) []error {
//v := reflect.ValueOf(param)
//k := v.Kind()

//if k == reflect.Array || k == reflect.Slice {
//this.processBulk(action, v, param)
//} else {
//this.Index(param)
//}

//return this.errors
//}

//func (this *Writer) processBulk(action string, v reflect.Value, param interface{}) {
//chunk := bytes.Buffer{}
//count := 0
//c := make(chan error)
//var wg sync.WaitGroup

//this.errors = []error{}

//go func() {
//for err := range c {
//this.appendError(err)
//}
//}()

//for i := 0; i < v.Len(); i++ {
//if count >= this.chunkLength || (i+1) == v.Len() {
//count = 0

//wg.Add(1)

//go func(action string, body []byte, chn chan error) {
//this.sendChunk(action, body, c)
//wg.Done()
//}(action, chunk.Bytes(), c)

//chunk.Reset()
//}

//count++

//itemBytes, err := this.createItemJsonBytes(action, v.Index(i).Interface())
//if err != nil {
//this.appendError(err)
//}

//chunk.Write(itemBytes)
//}

//wg.Wait()
//}

//func (this *Writer) sendChunk(action string, body []byte, c chan error) {
//urlStr := this.url + "/_bulk"
//res, err := this.requester.Post(urlStr, body)
//if err != nil {
//c <- err
//}

//result := &Result{}
//err = json.Unmarshal(res, result)
//if err != nil {
//return
//}

//if result.Errors {
//for _, m := range result.Items {
//item := m[action]

//if item.Error != "" {
//c <- errors.New(item.Error)
//}
//}
//}
//}

//func (this *Writer) createItemJsonBytes(action string, param interface{}) ([]byte, error) {
//id := uuid.New()

//buffer := bytes.Buffer{}
//buffer.WriteString(fmt.Sprintf(indexer_tmpl, action, id))

//paramBytes, err := json.Marshal(param)
//if err != nil {
//return []byte{}, err
//}

//buffer.WriteRune('\n')
//buffer.WriteString(string(paramBytes))
//buffer.WriteRune('\n')

//return buffer.Bytes(), nil
//}

//func (this *Writer) appendError(err error) {
//this.errors = append(this.errors, err)
//}
