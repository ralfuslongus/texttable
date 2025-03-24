package texttable

import (
	"testing"
)

func TestBool(testing *testing.T) {
	c := NewCell(true)
	w, h := c.RuneDim()
	AssertEqual(testing, w, 4)
	AssertEqual(testing, h, 1)
	println(c.String())
}
func TestEmptyAlignment(testing *testing.T) {
	c := NewCell("")
	w, h := c.RuneDim()
	AssertEqual(testing, w, 0)
	AssertEqual(testing, h, 0)
}
func TestEmptyCell(testing *testing.T) {
	c := NewCell("")
	w, h := c.RuneDim()
	AssertEqual(testing, w, 0)
	AssertEqual(testing, h, 0)
}
func TestCell(testing *testing.T) {
	c := NewCell("a\näöü€")
	w, h := c.RuneDim()
	AssertEqual(testing, w, 4)
	AssertEqual(testing, h, 2)

	c.SetAlignment(LEFT)
	AssertEqual(testing, c.RuneAt(-1, 0), ' ')
	AssertEqual(testing, c.RuneAt(0, 0), 'a')
	AssertEqual(testing, c.RuneAt(1, 0), ' ')
	AssertEqual(testing, c.RuneAt(2, 0), ' ')
	AssertEqual(testing, c.RuneAt(3, 0), ' ')
	AssertEqual(testing, c.RuneAt(4, 0), ' ')

	AssertEqual(testing, c.RuneAt(-1, 1), ' ')
	AssertEqual(testing, c.RuneAt(0, 1), 'ä')
	AssertEqual(testing, c.RuneAt(1, 1), 'ö')
	AssertEqual(testing, c.RuneAt(2, 1), 'ü')
	AssertEqual(testing, c.RuneAt(3, 1), '€')
	AssertEqual(testing, c.RuneAt(3, 1), '€')
	AssertEqual(testing, c.RuneAt(4, 1), ' ')

}
