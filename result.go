package golastic

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

type Result struct {
	Took     int       `json:"took"`
	TimedOut bool      `json:"timed_out"`
	Shards   ShardData `json:"_shards"`
	Hits     HitsData  `json:"hits"`
}

type ShardData struct {
	Total      int `json:"total"`
	Successful int `json:"successful"`
	Failed     int `json:"failed"`
}

type HitsData struct {
	Total    int          `json:"total"`
	MaxScore float64      `json:"max_score"`
	Hits     []ResultItem `json:"hits"`
}

func (this *Result) Scan(i interface{}) error {
	value := reflect.ValueOf(i)

	indirect := reflect.Indirect(value)

	if value.Type().Kind() == reflect.Ptr && indirect.Type().Kind() == reflect.Slice {

		for _, hit := range this.Hits.Hits {

			itemType := indirect.Type().Elem()

			elem := reflect.New(itemType).Elem()

			sourceBytes, err := json.Marshal(hit.Source)
			if err != nil {
				return err
			}

			err = json.Unmarshal(sourceBytes, elem.Addr().Interface())
			if err != nil {
				return err
			}

			indirect.Set(reflect.Append(indirect, elem))
		}

	} else {
		errorMsg := fmt.Sprintf(
			"The parameter must be a pointer to a slice, but it was %v : %v",
			value.Type().Kind(),
			indirect.Type().Kind(),
		)

		return errors.New(errorMsg)
	}

	return nil
}
