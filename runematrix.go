package texttable

import (
	"os"
)

type RuneMatrix struct {
	w, h  int
	Runes []rune
}

func NewRuneMatrix(w, h int) RuneMatrix {
	runes := make([]rune, (w+1)*h, (w+1)*h)
	m := RuneMatrix{w, h, runes}
	m.Fill(0, 0, w, h, ' ') // Leerzeichen schreiben
	// m.FillAll('.')           // Leerzeichen schreiben
	// m.Fill(w, 0, 1, h, '\n') // Zeilenumbrüche am rechten Rand schreiben
	return m
}

func (m *RuneMatrix) FillAll(r rune) {
	m.Fill(0, 0, m.w, m.h, r)
}
func (m *RuneMatrix) Fill(x, y, width, height int, r rune) {
	for a := x; a < x+width; a++ {
		for b := y; b < y+height; b++ {
			m.Set(a, b, r)
		}
	}
}
func (m *RuneMatrix) SmoothOpenCrossEnds() {
	for x := 0; x < m.w; x++ {
		for y := 0; y < m.h; y++ {
			if m.SmoothOpenCrossEnd(x, y) {
				// fmt.Printf("smoothed at %v/%v\n", x, y)
			}
		}
	}
}
func (m *RuneMatrix) HasFrameAt(x, y int) bool {
	r := m.Get(x, y)
	return r == '┼' || r == '│' || r == '─' || r == '┐' || r == '┘' || r == '└' || r == '┌' || r == '┬' || r == '┤' || r == '┴' || r == '├'
}
func (m *RuneMatrix) SmoothOpenCrossEnd(x, y int) bool {
	if m.Get(x, y) == '┼' {
		// left right top down Nachbarn fehlt (bool) wenn leer
		left, right, top, down := m.HasFrameAt(x-1, y), m.HasFrameAt(x+1, y), m.HasFrameAt(x, y-1), m.HasFrameAt(x, y+1)
		if top == true && right == false && down == true && left == false {
			m.Set(x, y, '│')
			return true
		}
		if top == false && right == true && down == false && left == true {
			m.Set(x, y, '─')
			return true
		}
		if top == false && right == false && down == true && left == true {
			m.Set(x, y, '┐')
			return true
		}
		if top == true && right == false && down == false && left == true {
			m.Set(x, y, '┘')
			return true
		}
		if top == true && right == true && down == false && left == false {
			m.Set(x, y, '└')
			return true
		}
		if top == false && right == true && down == true && left == false {
			m.Set(x, y, '┌')
			return true
		}
		if top == false && right == true && down == true && left == true {
			m.Set(x, y, '┬')
			return true
		}
		if top == true && right == false && down == true && left == true {
			m.Set(x, y, '┤')
			return true
		}
		if top == true && right == true && down == false && left == true {
			m.Set(x, y, '┴')
			return true
		}
		if top == true && right == true && down == true && left == false {
			m.Set(x, y, '├')
			return true
		}
	}
	return false
}

var t = `
┌─┬─┬─┐
│-│-│-│
├─┼─┼─┤
│-│-│-│
├─┼─┼─┤
│-│y│ │
└─┴─┴─┘`

func (m *RuneMatrix) Set(x, y int, r rune) {
	// alles außerhalb wird ignoriert
	if x < 0 || y < 0 || x > m.w || y > m.h {
		return
	}
	i := x + y*m.w
	m.Runes[i] = r
}
func (m *RuneMatrix) Get(x, y int) rune {
	// alles außerhalb wird als ' ' behandelt
	if x < 0 || y < 0 || x >= m.w || y >= m.h {
		return ' '
	}
	i := x + y*m.w
	return m.Runes[i]
}
func (m *RuneMatrix) HorizontalLineAt(y int) {
	var r rune
	for x := 0; x < m.w; x++ {
		r = m.Get(x, y)
		if r == ' ' {
			m.Set(x, y, '─')
		} else if r == '│' {
			m.Set(x, y, '┼')
		}
	}
}
func (m *RuneMatrix) VerticalLineAt(x int) {
	var r rune
	for y := 0; y < m.h; y++ {
		r = m.Get(x, y)
		if r == ' ' {
			m.Set(x, y, '│')
		} else if r == '─' {
			m.Set(x, y, '┼')
		}
	}
}
func WriteString(m *RuneMatrix, x, y int, s string) (int, int) {
	width, height := 0, 0
	dx, dy := 0, 1
	for _, v := range s {
		if v == '\n' || v == '\r' {
			dx = 0
			dy++
		} else {
			if m != nil {
				m.Set(x+dx, y+dy, v)
			}
			dx++
		}
		if dx > width {
			width = dx
		}
		if dy > height {
			height = dy
		}
	}
	return width, height
}

func (m *RuneMatrix) Render() {
	m.RenderTo(os.Stdout)
}
func (m *RuneMatrix) RenderTo(f *os.File) {
	for i, r := range m.Runes {
		f.WriteString(string(r))
		if (i+1)%m.w == 0 {
			f.WriteString("↵\n")
			// f.WriteString("\n")
		}
	}
}
