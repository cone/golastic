package golastic

import (
	"testing"
)

func TestParams(t *testing.T) {
	p := NewParam()

	p.Analyzer("test_analizer")
	p.CutoffFrequency(0.001)
	p.Query("this is a test")
	p.Type(TYPE_PHRASE_PREFIX)
	p.ZeroTermsQuery(ZERO_TERMS_QUERY_ALL)

	AssertEqualString(t, p.data["analyzer"].(string), "test_analizer")
	AssertEqualString(t, p.data["query"].(string), "this is a test")
	AssertEqualString(t, p.data["type"].(string), TYPE_PHRASE_PREFIX)
	AssertEqualString(t, p.data["zero_terms_query"].(string), ZERO_TERMS_QUERY_ALL)

	AssertEqualFloat(t, p.data["cutoff_frequency"].(float64), 0.001)
}
