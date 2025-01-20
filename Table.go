package texttable

import (
	"fmt"
	"os"
	"strings"
)

type ICell interface {
	W() int
	H() int
	RenderToMatrix(x int, y int, w int, h int, m *RuneMatrix)
}

// ---------------------------------------------------------
type Separator rune

func NewSeparator(r rune) Separator {
	sep := Separator(r)
	return sep
}

func (sep Separator) W() int {
	return 0
}
func (sep Separator) H() int {
	return 1
}
func (sep Separator) RenderToMatrix(x int, y int, w int, h int, m *RuneMatrix) {
	for dx := 0; dx < w; dx++ {
		for dy := 0; dy < h; dy++ {
			m.Set(x+dx, y+dy, rune(sep))
		}
	}
}

// ---------------------------------------------------------
type Cell struct {
	lines           [][]rune
	maxWidthOfLines int
}

func NewCell(val interface{}) *Cell {
	multiLine := fmt.Sprintf("%v", val)
	parts := strings.Split(multiLine, "\n")
	lines := make([][]rune, len(parts))
	maxWidthOfLines := 0
	for i, part := range parts {
		line := []rune(part)
		lines[i] = line
		maxWidthOfLines = max(maxWidthOfLines, len(line))
	}
	cell := Cell{lines: lines, maxWidthOfLines: maxWidthOfLines}
	return &cell
}
func (cell *Cell) W() int {
	return cell.maxWidthOfLines
}
func (cell *Cell) H() int {
	return len(cell.lines)
}
func (cell *Cell) RenderToMatrix(x int, y int, w int, h int, m *RuneMatrix) {
	for i, line := range cell.lines {
		for j, r := range line {
			m.Set(x+j, y+i, r)
		}
	}
}

// ---------------------------------------------------------
type Table struct {
	cells          []ICell // alle Zellen in einem slice, wächst mit jedem Add
	numberOfColums int     // die fixe Anzahl an Spalten
	maxColW        []int   // die Maximalbreiten aller Spalten, wird bei jedem Add aktualisiert
	maxRowH        []int   // die Maximalhöhen aller Zeilen, wird bei jedem Add aktualisiert
}

func NewTable(numberOfColums, startCapacityOfRows int) *Table {
	n := numberOfColums * startCapacityOfRows
	cells := make([]ICell, 0, n)
	maxColW := make([]int, numberOfColums)
	maxRowH := make([]int, startCapacityOfRows)
	t := Table{cells: cells, numberOfColums: numberOfColums, maxColW: maxColW, maxRowH: maxRowH}
	return &t
}

func (t *Table) W() int {
	sum := 0
	for _, w := range t.maxColW {
		sum += w
	}
	return sum + t.numberOfColums - 1
}

func (t *Table) H() int {
	sum := 0
	for _, h := range t.maxRowH {
		sum += h
	}
	return sum
}

func (t *Table) AddSeparatorsTillEndOfRow() {
	for {
		t.AddSeparator()
		if len(t.cells)%t.numberOfColums == 0 {
			break
		}
	}
}

func (t *Table) AddSeparator() {
	// Separator erzeugen und anhängen
	sep := NewSeparator('┼')
	t.append(sep)
}

func (t *Table) Add(vals ...interface{}) {
	for _, val := range vals {
		// Type Assertion
		if cell, ok := val.(ICell); ok {
			// fmt.Println("val IS of type ICell")
			// Zelle  so wie sie ist anhängen
			t.append(cell)
		} else {
			// fmt.Println("val is not of type ICell")
			// Zelle  erzeugen und anhängen
			cell := NewCell(val)
			t.append(cell)
		}
	}
}
func (t *Table) append(cell ICell) {
	// ...Zelle anhängen
	index := len(t.cells)
	t.cells = append(t.cells, cell)

	// Zeilen- und Spaltennummern der neuen Zelle berechnen
	rowNum := index / t.numberOfColums
	colNum := index % t.numberOfColums
	// fmt.Printf("adding value to index %d at row/col %d/%d\n", index, rowNum, colNum)

	// Maximalbreite der aktuellen Spalte updaten, Maxima wurden schon im Konstruktor angelegt weil Anzahl Spalten bekannt ist
	w := cell.W()
	t.maxColW[colNum] = max(t.maxColW[colNum], w)

	// Maximalhöhe der aktuellen Zeile updaten
	h := cell.H()
	if len(t.maxRowH) <= rowNum { // Zeile hat noch kein Maximum da startCapacityOfRows überschritten wurde, also anlegen, oder...
		// fmt.Printf("Neues Zeilenmaximum für Zeile %d anlegen\n", rowNum)
		t.maxRowH = append(t.maxRowH, h)
	} else {
		t.maxRowH[rowNum] = max(t.maxRowH[rowNum], h) // .. neues Maximum ermitteln und setzten
	}
}

func (t *Table) Render(smooth bool, withOuterFrame bool) {
	t.RenderTo(os.Stdout, smooth, withOuterFrame)
}

func (t *Table) RenderTo(f *os.File, smooth bool, withOuterFrame bool) {
	var m RuneMatrix
	if withOuterFrame {
		m = NewRuneMatrix(t.W()+2, t.H()+2)
		// m.FillAll('⋅')
		m.HorizontalLineAt(0)
		m.HorizontalLineAt(m.h - 1)
		m.VerticalLineAt(0)
		m.VerticalLineAt(m.w - 1)
		t.RenderToMatrix(1, 1, t.W(), t.H(), &m)
	} else {
		m = NewRuneMatrix(t.W(), t.H())
		// m.FillAll('⋅')
		t.RenderToMatrix(0, 0, t.W(), t.H(), &m)

	}
	if smooth {
		m.SmoothOpenCrossEnds()
	}
	m.RenderTo(f)
}

func (t *Table) RenderToMatrix(x int, y int, w int, h int, m *RuneMatrix) {
	index := 0
	rowNum := 0
	dy := 0
	dx := 0
	for _, maxH := range t.maxRowH { // alle ZEILEN durchgehen
		dx = 0
		for i, maxW := range t.maxColW { // alle SPALTEN (der zeile) durchgehen
			if index >= len(t.cells) {
				break
			}
			cell := t.cells[index]
			cell.RenderToMatrix(x+dx, y+dy, maxW, maxH, m)
			index++
			if i < len(t.maxColW)-1 { // wenn noch eine Spalte kommt Separator zeichen
				for sy := 0; sy < maxH; sy++ {
					m.Set(x+dx+maxW, y+dy+sy, '│')
				}
				dx++
				dx += maxW
			}
		}
		dy += maxH
		rowNum++
	}
}
