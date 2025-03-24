package texttable

import (
	"fmt"
	"testing"

	"github.com/dottedmag/xd"
)

func TestSingleRunes(t *testing.T) {

	rIn := 'a'
	b := DEFAULT_CODEPAGE.encodeRune(rIn)
	rOut := DEFAULT_CODEPAGE.ecodeByte(b)
	AssertEqual(t, rOut, rIn)

	rIn = 'b'
	b = DEFAULT_CODEPAGE.encodeRune(rIn)
	rOut = DEFAULT_CODEPAGE.ecodeByte(b)
	AssertEqual(t, rOut, rIn)

	rIn = 'c'
	b = DEFAULT_CODEPAGE.encodeRune(rIn)
	rOut = DEFAULT_CODEPAGE.ecodeByte(b)
	AssertEqual(t, rOut, rIn)
}
func TestStrings(t *testing.T) {
	sIn := "abcdefghij"
	bytes := DEFAULT_CODEPAGE.Encode(sIn)
	fmt.Printf("% x\n", bytes)
	AssertEqual(t, bytes[0], byte(1))
	AssertEqual(t, bytes[1], byte(2))
	AssertEqual(t, bytes[2], byte(3))

	sOut := DEFAULT_CODEPAGE.Decode(bytes)
	println("sIn bytes")
	xd.Print([]byte(sIn), 0)
	println("sOut bytes")
	xd.Print([]byte(sOut), 0)
	AssertEqual(t, sOut, sIn)
}
func TestOverflow(t *testing.T) {
	count := 300
	toMuchDifferentRunes := make([]rune, count)
	for i := 0; i < count; i++ {
		r := rune(i)
		toMuchDifferentRunes[i] = r
	}

	sIn := string(toMuchDifferentRunes)
	bytes := DEFAULT_CODEPAGE.Encode(sIn)
	sOut := DEFAULT_CODEPAGE.Decode(bytes)
	// println("sIn bytes")
	// xd.Print([]byte(sIn), 0)
	// println("sOut bytes")
	// xd.Print([]byte(sOut), 0)

	goodIn := sIn[0:255]
	badIn := sIn[255:]
	goodOut := sOut[0:255]
	badOut := sOut[255:]
	println("goodIn:", goodIn)
	println("goodOut:", goodOut)
	println("badIn:", badIn)
	println("badOut:", badOut)
	AssertEqual(t, goodIn, goodOut)
	AssertNotEqual(t, badIn, badOut)
}
