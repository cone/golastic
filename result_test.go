package golastic

import (
	"encoding/json"
	"testing"
)

func TestResult(t *testing.T) {
	result := &Result{}

	jsonBytes := []byte(`{
  	"took" : 193,
  	"timed_out" : false,
  	"_shards" : {
    	"total" : 5,
    	"successful" : 5,
    	"failed" : 0
  	},
  	"hits" : {
    	"total" : 1,
    	"max_score" : 1.0,
    	"hits" : [ {
      	"_index" : "test",
      	"_type" : "products",
      	"_id" : "4",
      	"_score" : 1.0,
      	"_source":{"id":13,"name":"Ruby on Rails Mug"}
    	} ]
  	}
	}`)

	err := json.Unmarshal(jsonBytes, result)
	if err != nil {
		t.Error("An error has ocurred: " + err.Error())
	}

	AssertEqualInt(t, result.Took, 193)
	AssertEqualBool(t, result.TimedOut, false)

	AssertEqualInt(t, result.Shards.Total, 5)
	AssertEqualInt(t, result.Shards.Successful, 5)
	AssertEqualInt(t, result.Shards.Failed, 0)

	AssertEqualInt(t, result.Hits.Total, 1)
	AssertEqualFloat(t, result.Hits.MaxScore, 1.0)
	AssertEqualInt(t, len(result.Hits.Hits), 1)
}
