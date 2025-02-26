package texttable

import (
	"testing"
)

func TestNewRuneMatrix(t *testing.T) {
	m := NewRuneMatrix(3, 3)
	m.FillAll('x')
	AssertEqual(t, m.String(), "xxx\nxxx\nxxx")
	AssertEqual(t, m.w, 3)
	AssertEqual(t, m.h, 3)
}
func AssertEqual(t *testing.T, result, expected interface{}) {
	if result != expected {
		t.Errorf("Expected:\n'%v'\n, but got:\n'%v'\n", expected, result)
	}
}
