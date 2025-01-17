package texttable3

import (
	"fmt"
	"os"
	"strings"
	"temp/texttable"
)

type ICell interface {
	W() int
	H() int
	RenderToMatrix(int, int, *texttable.RuneMatrix)
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
func (cell *Cell) RenderToMatrix(x int, y int, m *texttable.RuneMatrix) {
	for i, line := range cell.lines {
		for j, r := range line {
			m.Set(x+j, y+i, r)
		}
	}
}

// ---------------------------------------------------------
type Table struct {
	cells          []*Cell // alle Zellen in einem slice, wächst mit jedem Add
	numberOfColums int     // die fixe Anzahl an Spalten
	maxColW        []int   // die Maximalbreiten aller Spalten, wird bei jedem Add aktualisiert
	maxRowH        []int   // die Maximalhöhen aller Zeilen, wird bei jedem Add aktualisiert
	borders        Borders
}

func NewTable(numberOfColums, startCapacityOfRows int, borders Borders) *Table {
	n := numberOfColums * startCapacityOfRows
	cells := make([]*Cell, 0, n)
	maxColW := make([]int, numberOfColums)
	maxRowH := make([]int, startCapacityOfRows)
	table := Table{cells: cells, numberOfColums: numberOfColums, maxColW: maxColW, maxRowH: maxRowH, borders: borders}
	return &table
}

func (table *Table) W() int {
	sum := 0
	if table.borders.Has(OUTER) {
		sum += 2
	}
	if table.borders.Has(COLSEP) {
		sum += len(table.maxColW) - 1
	}
	for _, w := range table.maxColW {
		sum += w
	}
	return sum
}

func (table *Table) H() int {
	sum := 0
	if table.borders.Has(OUTER) {
		sum += 2
	}
	if table.borders.Has(ROWSEP) {
		sum += len(table.maxRowH) - 1
	}
	for _, h := range table.maxRowH {
		sum += h
	}
	return sum
}
func (table *Table) Add(val interface{}) {
	// Zelle erzeugen und anhängen
	index := len(table.cells)
	cell := NewCell(val)
	table.cells = append(table.cells, cell)

	// Zeilen- und Spaltennummern der neuen Zelle berechnen
	rowNum := index / table.numberOfColums
	colNum := index % table.numberOfColums
	fmt.Printf("adding value to index %d at row/col %d/%d\n", index, rowNum, colNum)

	// Maximalbreite der aktuellen Spalte updaten, Maxima wurden schon im Konstruktor angelegt weil Anzahl Spalten bekannt ist
	w := cell.W()
	table.maxColW[colNum] = max(table.maxColW[colNum], w)

	// Maximalhöhe der aktuellen Zeile updaten
	h := cell.H()
	if len(table.maxRowH) <= rowNum { // Zeile hat noch kein Maximum da startCapacityOfRows überschritten wurde, also anlegen, oder...
		fmt.Printf("Neues Zeilenmaximum für Zeile %d anlegen\n", rowNum)
		table.maxRowH = append(table.maxRowH, h)
	} else {
		table.maxRowH[rowNum] = max(table.maxRowH[rowNum], h) // .. neues Maximum ermitteln und setzten
	}
}

func (table *Table) Render() {
	table.RenderTo(os.Stdout)
}

func (table *Table) RenderTo(f *os.File) {
	m := texttable.NewRuneMatrix(table.W(), table.H())
	table.RenderToMatrix(0, 0, &m)
	m.RenderTo(f)
}

func (table *Table) RenderToMatrix(x int, y int, m *texttable.RuneMatrix) {
	index := 0
	rowNum := 0
	dy := 0
	dx := 0
	if table.borders.Has(OUTER) { // oberen rand auslassen
		m.HorizontalLineAt(y+dy)
		dy++
	}
	for _, maxRowH := range table.maxRowH { // alle ZEILEN durchgehen
		dx = 0
		if table.borders.Has(OUTER) { // linken rand auslassen
			dx++
		}
		for _, maxColW := range table.maxColW { // alle SPALTEN (der zeile) durchgehen
			cell := table.cells[index]
			cell.RenderToMatrix(x+dx, y+dy, m)
			if table.borders.Has(COLSEP) { // senkrechten spaltentrenner auslassen
				dx++
			}
			dx += maxColW
			index++
			// if index == len(table.cells) {
			// 	return
			// }
		}
		if table.borders.Has(OUTER) { // rechten rand auslassen
			dx++
		}
		dy += maxRowH

		// waagerechten header-, footer- oder zeilentrenner auslassen

		// zeilentrenner
		rowtrenner := table.borders.Has(ROWSEP)
		// headertrenner
		headertrenner := rowNum == 0 && table.borders.Has(HEADER)
		// footertrenner
		footertrenner := rowNum-2 == len(table.maxRowH) && table.borders.Has(FOOTER)
		fmt.Printf("borders: %v rowNum: %d rowtrenner: %v headertrenner: %v footertrenner: %v \n", table.borders, rowNum, rowtrenner, headertrenner, footertrenner)

		if rowtrenner || headertrenner || footertrenner {
			dy++
		}
		rowNum++
	}
	if table.borders.Has(OUTER) { // unteren rand auslassen
		m.HorizontalLineAt(y+dy)
		dy++
	}
}

// ---------------------------------------------------------
type Borders int8

const (
	OUTER  Borders = 1
	HEADER Borders = 2
	FOOTER Borders = 4
	ROWSEP Borders = 8
	COLSEP Borders = 16
	INNER  Borders = HEADER | FOOTER | ROWSEP | COLSEP
	ALL    Borders = OUTER | INNER
)

func (b Borders) Has(borders Borders) bool {
	return b&borders != 0
}
func (b Borders) With(borders Borders) Borders {
	return b | borders
}
func (b Borders) Without(borders Borders) Borders {
	return b ^ borders
}
