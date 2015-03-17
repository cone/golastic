package golastic

import "testing"

func TestField_Params(t *testing.T) {
	field := NewField("test")

	if field.ParamSet != nil {
		t.Error("Error: field.ParamSet should be nil")
	}

	field.Params(NewParam().Query("this is a test"))

	if field.ParamSet == nil {
		t.Error("Error: field.ParamSet should not be nil")
	}
}

func TestField_Query(t *testing.T) {
	field := NewField("test")

	expected := "this is a test"

	field.Query(expected)

	AssertEqualString(t, field.OutterQuery, expected)
}

func TestField_ToMap(t *testing.T) {
	field := NewField("test")

	expected := "this is a test"

	field.Query(expected)

	m := field.ToMap()

	AssertEqualString(t, m["test"].(string), expected)

	field.Params(NewParam().Query(expected))

	m = field.ToMap()

	paramMap := m["test"].(map[string]interface{})

	AssertEqualString(t, paramMap["query"].(string), expected)
}
