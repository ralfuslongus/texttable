package texttable

import (
	"testing"
)

func AssertNotEqual(t *testing.T, result, expected any) {
	if result == expected {
		t.Errorf("result and expected are equal")
	}
}
func AssertEqual(t *testing.T, result, expected any) {
	if result != expected {
		t.Errorf("Expected %T:'%v', but got %T:'%v'\n", expected, expected, result, result)
	}
}
