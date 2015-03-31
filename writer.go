package golastic

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"
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

func (this *Writer) Bulk(doc *BulkDoc) []error {
	c := make(chan error)
	errors := []error{}
	var wg sync.WaitGroup

	go func() {
		for err := range c {
			errors = append(errors, err)
		}
	}()

	ln := doc.Len()
	from := 0
	to := this.chunkLength

	for {

		if from >= ln {
			break
		}

		chunk, err := doc.Read(from, to)
		if err != nil {
			c <- err
			continue
		}

		wg.Add(1)

		go func(chunk []byte) {

			this.sendChunk(chunk, c)
			wg.Done()

		}(chunk)

		from = from + this.chunkLength
		to = to + this.chunkLength

	}

	wg.Wait()

	return errors
}

func (this *Writer) sendChunk(body []byte, c chan error) {
	urlStr := this.url + "/_bulk"
	res, err := this.requester.Post(urlStr, body)
	if err != nil {
		c <- err
	}

	result := &Result{}
	err = json.Unmarshal(res, result)
	if err != nil {
		return
	}

	if result.Errors {
		for _, m := range result.Items {
			for action, item := range m {
				if item.Error != "" {
					c <- errors.New(fmt.Sprintf("Error for %s action: %s", action, item.Error))
				}
			}
		}
	}
}
