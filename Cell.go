package texttable

import (
	"io"
	"strings"
)

// ----------------------------------------------
type Alignment byte

const (
	LEFT   Alignment = 0
	CENTER Alignment = iota
	RIGHT  Alignment = iota
)

// ----------------------------------------------
type Cell struct {
	lines     [][]byte
	alignment Alignment
	W, H      int
}

func (c *Cell) SetAlignment(a Alignment) ICell {
	c.alignment = a
	return c
}

func NewCell(v any) ICell {
	var s string
	var a Alignment
	switch val := v.(type) {
	case nil:
		return nil
	case string:
		a = LEFT
		s = val
	case bool:
		a = RIGHT
		s = BoolToString(val)
	case int:
		a = RIGHT
		s = IntToString(val)
	case ICell:
		return val
	default:
		s = "!unsupported type!"
	}

	s = strings.ReplaceAll(s, "\r\n", "\n") // win to linux
	s = strings.ReplaceAll(s, "\r", "\n")   // mac to linux
	var parts []string
	if s == "" {
		parts = nil
	} else {
		parts = strings.Split(s, "\n")
	}
	w := 0
	h := len(parts)
	lines := make([][]byte, 0, h)
	for _, part := range parts {
		line := DEFAULT_CODEPAGE.Encode(part)
		// fmt.Printf("line %d: %v\n", i, line)
		lines = append(lines, line)
		if len(line) > w {
			w = len(line)
		}
	}
	// println("alignment: ", a)
	c := Cell{lines: lines, alignment: a, W: w, H: h}
	// fmt.Printf("lines: %v\n", lines)
	return &c
}

func BoolToString(b bool) string {
	if b {
		return "true"
	} else {
		return "false"
	}
}
func IntToString(num int) string {
	if num == 0 {
		return "0"
	}

	// Bestimme das Vorzeichen
	isNegative := num < 0
	if isNegative {
		num = -num
	}

	// Erstelle einen Slice von runes für die Ziffern
	var digits []rune
	for num > 0 {
		digit := num % 10
		digits = append(digits, rune('0'+digit)) // '0' ist der ASCII-Wert für die Ziffer 0
		num /= 10
	}

	// Wenn die Zahl negativ ist, füge das Minuszeichen hinzu
	if isNegative {
		digits = append(digits, '-')
	}

	// Umkehren der Ziffern, da sie in umgekehrter Reihenfolge hinzugefügt wurden
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}

	return string(digits)
}

// ----------------------------------------------
func (c *Cell) RuneDim() (width, height int) {
	return c.W, c.H
}
func (c *Cell) RuneAt(x, y, w, h int) rune {
	// out of range
	if y < 0 || y >= c.H {

		return ' '
		// fmt.Printf("cell %v: Cell.RuneAt x/y:", c, x, y)
		// panic("Cell.RuneAt, out of range")
	}
	line := c.lines[y]

	var alignOffset int
	switch c.alignment {
	case LEFT:
		alignOffset = 0
	case CENTER:
		alignOffset = (w - len(line)) / 2
	case RIGHT:
		alignOffset = w - len(line)
	default:
		alignOffset = 0
	}
	x = x - alignOffset

	if x < 0 || x >= len(line) {
		return ' '
	}
	b := line[x]
	r := DEFAULT_CODEPAGE.ecodeByte(b)
	return r
}

// ----------------------------------------------
func (c *Cell) String() string {
	sb := strings.Builder{}
	c.WriteTo(&sb)
	return sb.String()
}

func (c *Cell) WriteTo(writer io.Writer) (int, error) {
	bytesWritten := 0
	for _, line := range c.lines {
		s := DEFAULT_CODEPAGE.Decode(line)

		i, err := writer.Write([]byte(s))
		bytesWritten += i
		if err != nil {
			return bytesWritten, err
		}
	}
	// i, err := writer.Write([]byte("\n"))
	// bytesWritten += i
	// if err != nil {
	// 	return bytesWritten, err
	// }
	return bytesWritten, nil
}
