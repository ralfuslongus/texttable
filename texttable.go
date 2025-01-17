package texttable

import (
	"errors"
	"fmt"
	"os"
)

type Table struct {
	width       int
	cells       []string
	cellWidths  []int
	cellHeights []int
}

func NewTable(width int) Table {
	t := Table{
		width: width,
		cells: make([]string, 0),
	}
	return t
}
func (t *Table) Height() int {
	return len(t.cells) / t.width
}
func (t *Table) Width() int {
	return t.width
}
func (t *Table) Add(row ...interface{}) error {
	// func (t *Table) Add(row ...string) error {
	if len(row) != t.width {
		return errors.New("row width must match table width")
	}
	for _, cell := range row {
		s := fmt.Sprintf("%v", cell)
		w, h := calcDimensionsOf(s)
		t.cells = append(t.cells, s)
		t.cellWidths = append(t.cellWidths, w)
		t.cellHeights = append(t.cellHeights, h)
	}
	return nil
}
func (t *Table) getCell(x, y int) string {
	return t.cells[x+y*t.width]
}
func (t *Table) getCellDimensions(x, y int) (int, int) {
	i := x + y*t.width
	return t.cellWidths[i], t.cellHeights[i]
}
func calcDimensionsOf(cell string) (int, int) {
	return WriteString(nil, 0, 0, cell)
}
func (t *Table) RenderTo(f *os.File) {
	fmt.Printf("width/height of Table: %d/%d\n", t.Width(), t.Height())
	maxH := make([]int, t.Height()) // maximal height of every row
	maxW := make([]int, t.Width())  // maximal widths of every colums

	for y := 0; y < t.Height(); y++ {
		for x := 0; x < t.Width(); x++ {
			w, h := t.getCellDimensions(x, y)
			if maxW[x] < w {
				// fmt.Printf("updating maxW[%d] from %d to %d\n", x, maxW[x], w)
				maxW[x] = w
			}
			if maxH[y] < h {
				// fmt.Printf("updating maxH[%d] from %d to %d\n", y, maxH[y], h)
				maxH[y] = h
			}
		}
	}
	fmt.Printf("maxW/maxH of Table: %v/%v\n", maxW, maxH)

	// Gesamtbreite w ist Summe aller Maximalbreiten + alle Rahmen
	w := 0
	for _, v := range maxW {
		w += v
	}
	w += 1
	w += len(maxW)

	// Gesamthöhe h ist Summe aller Maximalhöhen + Rahmen oben und unten
	h := 0
	for _, v := range maxH {
		h += v
	}
	h += 2

	rm := NewRuneMatrix(w, h)
	fmt.Printf("w/h of RuneMatrix: %v/%v\n", w, h)

	rm.HorizontalLineAt(0)
	rm.HorizontalLineAt(h - 1)

	x := 0
	rm.VerticalLineAt(x)
	for _, v := range maxW {
		x += v + 1
		rm.VerticalLineAt(x)
	}

	// Alle cells zeichen
	fmt.Printf("cells: %v\n", t.cells)
	my := 1
	for ty := 0; ty < t.Height(); ty++ {
		mx := 1
		for tx := 0; tx < t.Width(); tx++ {
			s := t.getCell(tx, ty)
			// fmt.Printf("cell at %d/%d:%v\n", tx, ty, s)
			fmt.Printf("mx/my:%d/%d\n", mx, my)
			WriteString(&rm, mx, my, s)
			mx += 1 + maxW[tx]
		}
		my += maxH[ty]
	}

	rm.SmoothOpenCrossEnds()
	rm.RenderTo(f)
}
