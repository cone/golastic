package golastic

import (
	"testing"
)

func TestParams(t *testing.T) {
	p := Params{}

	p.Analyzer("test_analizer")
	p.CutoffFrequency(0.001)
	p.Query("this is a test")
	p.Type(TYPE_PHRASE_PREFIX)
	p.ZeroTermsQuery(ZERO_TERMS_QUERY_ALL)

	AssertEqualString(t, p["analyzer"].(string), "test_analizer")
	AssertEqualString(t, p["query"].(string), "this is a test")
	AssertEqualString(t, p["type"].(string), TYPE_PHRASE_PREFIX)
	AssertEqualString(t, p["zero_terms_query"].(string), ZERO_TERMS_QUERY_ALL)

	AssertEqualFloat(t, p["cutoff_frequency"].(float64), 0.001)
}
