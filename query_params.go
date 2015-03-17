package golastic

type QueryParams struct {
	Query           string `json:"query,omitempty"`
	Operator        string `json:"operator,omitempty"`
	ZeroTermsQuery  string `json:"zero_terms_query,omitempty"`
	CutoffFrequency string `json:"cutoff_frequency,omitempty"`
	Type            string `json:"type,omitempty"`
	Analyzer        string `json:"analyzer,omitempty"`
}
