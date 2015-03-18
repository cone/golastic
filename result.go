package golastic

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
