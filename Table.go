package texttable

import (
	"strings"
)

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
	sep := NewCell(nil).AsSeparator()
	t.append(sep)
}
func (t *Table) Add(vals ...interface{}) {
	for _, v := range vals {
		t.append(NewCell(v))
	}
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
func (t *Table) append(cell *Cell) {
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
	// fmt.Printf("w/h of cell %v = %d/%d\n", cell.String(), cell.W(), cell.H())
}
func (t *Table) RenderTo(f *strings.Builder, smooth bool, withOuterFrame bool) {
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
func (t *Table) String() string {
	return t.ToString(true)

}
func (t *Table) ToString(withOuterFrame bool) string {
	sb := strings.Builder{}
	t.RenderTo(&sb, true, withOuterFrame)
	return sb.String()

}
