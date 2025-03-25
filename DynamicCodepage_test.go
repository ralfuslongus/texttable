package texttable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleRunes(t *testing.T) {
	cp := NewDynamicCodepage()
	rIn := 'a'
	b := cp.encodeRune(rIn)
	rOut := cp.ecodeByte(b)
	assert.Equal(t, rOut, rIn)

	rIn = 'b'
	b = cp.encodeRune(rIn)
	rOut = cp.ecodeByte(b)
	assert.Equal(t, rOut, rIn)

	rIn = 'c'
	b = cp.encodeRune(rIn)
	rOut = cp.ecodeByte(b)
	assert.Equal(t, rOut, rIn)
}
func TestStrings(t *testing.T) {
	cp := NewDynamicCodepage()
	sIn := "abcdefghij"
	bytes := cp.Encode(sIn)
	// fmt.Printf("% x\n", bytes)
	assert.Equal(t, bytes[0], byte(1))
	assert.Equal(t, bytes[1], byte(2))
	assert.Equal(t, bytes[2], byte(3))

	sOut := cp.Decode(bytes)
	// println("sIn bytes")
	// xd.Print([]byte(sIn), 0)
	// println("sOut bytes")
	// xd.Print([]byte(sOut), 0)
	assert.Equal(t, sOut, sIn)
}
func TestOverflow(t *testing.T) {
	cp := NewDynamicCodepage()
	count := 300
	toMuchDifferentRunes := make([]rune, count)
	for i := 0; i < count; i++ {
		r := rune(i)
		toMuchDifferentRunes[i] = r
	}

	sIn := string(toMuchDifferentRunes)
	bytes := cp.Encode(sIn)
	sOut := cp.Decode(bytes)
	// // println("sIn bytes")
	// // xd.Print([]byte(sIn), 0)
	// // println("sOut bytes")
	// // xd.Print([]byte(sOut), 0)

	goodIn := sIn[0:255]
	badIn := sIn[255:]
	goodOut := sOut[0:255]
	badOut := sOut[255:]
	// println("goodIn:", goodIn)
	// println("goodOut:", goodOut)
	// println("badIn:", badIn)
	// println("badOut:", badOut)
	assert.Equal(t, goodIn, goodOut)
	assert.NotEqual(t, badIn, badOut)
}
