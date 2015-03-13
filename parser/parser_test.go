package parser

import (
	"testing"
)

func TestParser_Parse(t *testing.T) {
	resultGenerator := newFakeResultGenerator()

	parser := NewParser(resultGenerator)

	result := ""

	parser.Parse("this_thing_test", 0, &result)

	expected := "this_thing is a test"

	if result != expected {
		t.Errorf("Not matching, got: %s, expecting: %s", result, expected)
	}

	parser.Parse("the_long_long_string_test", 0, &result)

	expected = "the_long_long_string is a test"

	if result != expected {
		t.Errorf("Not matching, got: %s, expecting: %s", result, expected)
	}
}
