package golastic

//import (
//"encoding/json"
//"fmt"
//"net/http"
//"net/http/httptest"
//"testing"
//)

//func TestFetcher_Find(t *testing.T) {
//jsonStr := `{
//"_index" : "test",
//"_type" : "products",
//"_id" : "1",
//"_version" : 1,
//"found" : true,
//"_source":{"id":16,"name":"Spree Mug"}
//}`

//ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//fmt.Fprintln(w, jsonStr)
//}))
//defer ts.Close()

//fetcher := NewFetcher(ts.URL)

//resBytes, err := fetcher.doRequest("GET", ts.URL, nil)
//if err != nil {
//t.Error("An error has ocurred: " + err.Error())
//}

//result := &Result{}
//err = json.Unmarshal(resBytes, result)
//if err != nil {
//t.Error("An error has ocurred: " + err.Error())
//}

//t.Error(result)
//}
