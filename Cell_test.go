package texttable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBool(testing *testing.T) {
	c := NewCell(true)
	w, h := c.RuneDim()
	assert.Equal(testing, w, 4)
	assert.Equal(testing, h, 1)
}
func TestEmptyAlignment(testing *testing.T) {
	c := NewCell("")
	w, h := c.RuneDim()
	assert.Equal(testing, w, 0)
	assert.Equal(testing, h, 0)
}
func TestEmptyCell(testing *testing.T) {
	c := NewCell("")
	w, h := c.RuneDim()
	assert.Equal(testing, w, 0)
	assert.Equal(testing, h, 0)
}
func TestCell(testing *testing.T) {
	c := NewCell("a\näöü€")
	w, h := c.RuneDim()
	assert.Equal(testing, w, 4)
	assert.Equal(testing, h, 2)

	c.SetAlignment(LEFT)
	assert.Equal(testing, c.RuneAt(-1, 0, w, h), ' ')
	assert.Equal(testing, c.RuneAt(0, 0, w, h), 'a')
	assert.Equal(testing, c.RuneAt(1, 0, w, h), ' ')
	assert.Equal(testing, c.RuneAt(2, 0, w, h), ' ')
	assert.Equal(testing, c.RuneAt(3, 0, w, h), ' ')
	assert.Equal(testing, c.RuneAt(4, 0, w, h), ' ')

	assert.Equal(testing, c.RuneAt(-1, 1, w, h), ' ')
	assert.Equal(testing, c.RuneAt(0, 1, w, h), 'ä')
	assert.Equal(testing, c.RuneAt(1, 1, w, h), 'ö')
	assert.Equal(testing, c.RuneAt(2, 1, w, h), 'ü')
	assert.Equal(testing, c.RuneAt(3, 1, w, h), '€')
	assert.Equal(testing, c.RuneAt(3, 1, w, h), '€')
	assert.Equal(testing, c.RuneAt(4, 1, w, h), ' ')

}
