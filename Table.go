package texttable

import (
	"fmt"
	"os"
	"strings"
)

type ICell interface {
	W() int
	H() int
	RenderToMatrix(int, int, *RuneMatrix)
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
func (cell *Cell) RenderToMatrix(x int, y int, m *RuneMatrix) {
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
}

func NewTable(numberOfColums, startCapacityOfRows int) *Table {
	n := numberOfColums * startCapacityOfRows
	cells := make([]*Cell, 0, n)
	maxColW := make([]int, numberOfColums)
	maxRowH := make([]int, startCapacityOfRows)
	table := Table{cells: cells, numberOfColums: numberOfColums, maxColW: maxColW, maxRowH: maxRowH}
	return &table
}

func (table *Table) W() int {
	sum := 0
	for _, w := range table.maxColW {
		sum += w
	}
	return sum
}

func (table *Table) H() int {
	sum := 0
	for _, h := range table.maxRowH {
		sum += h
	}
	return sum
}
func (table *Table) Add(vals ...interface{}) {
	for _, val := range vals {
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
}

func (table *Table) Render() {
	table.RenderTo(os.Stdout)
}

func (table *Table) RenderTo(f *os.File) {
	m := NewRuneMatrix(table.W(), table.H())
	table.RenderToMatrix(0, 0, &m)
	m.RenderTo(f)
}

func (table *Table) RenderToMatrix(x int, y int, m *RuneMatrix) {
	index := 0
	rowNum := 0
	dy := 0
	dx := 0
	for _, maxRowH := range table.maxRowH { // alle ZEILEN durchgehen
		dx = 0
		for _, maxColW := range table.maxColW { // alle SPALTEN (der zeile) durchgehen
			cell := table.cells[index]
			cell.RenderToMatrix(x+dx, y+dy, m)
			dx += maxColW
			index++
			// if index == len(table.cells) {
			// 	return
			// }
		}
		dy += maxRowH

		rowNum++
	}
}
