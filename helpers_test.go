package golastic

import "testing"

func AssertEqualString(t *testing.T, a, b string) {
	if a != b {
		mismatchError(t, a, b)
	}
}

func AssertEqualFloat(t *testing.T, a, b float64) {
	if a != b {
		mismatchError(t, a, b)
	}
}

func mismatchError(t *testing.T, a, b interface{}) {
	t.Errorf("Mismatch.\n%s is not equal to %s\n", a, b)
}
