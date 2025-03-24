package texttable

import (
	"errors"
	"io"
	"strings"
)

type Table struct {
	cells   []ICell
	columns int // fix
	index   int
	// rows         int	// grows
	borderConfig *BorderConfig
	// cached (Rune-)Width(s):
	cachedWidths  []int
	cachedHeights []int
	cachedWidth   int
	cachedHeight  int
}

func (t *Table) SetAlignment(a Alignment) ICell {
	println("!!!!! SetAlignment not implemented by Table yet !!!!!")
	return t
}

// ----------------------------------------------
func NewTable(colums, initialRows int, borderConfig *BorderConfig) *Table {
	cells := make([]ICell, 0, colums*initialRows)
	t := Table{cells, colums, 0, borderConfig, nil, nil, -1, -1}
	return &t
}

// ----------------------------------------------
func (t *Table) Append(stringOrCells ...any) ICell {
	var cell ICell
	for _, stringOrCell := range stringOrCells {
		// important: when modifieing table, clear all cached values
		t.cachedWidths = nil
		t.cachedHeights = nil
		t.cachedWidth = -1
		t.cachedHeight = -1

		cell = NewCell(stringOrCell)
		newCells := append(t.cells, cell)

		if cap(newCells) != cap(t.cells) {
			println("!!!!!warning: slice of cells changed by appending from ", cap(t.cells), "to", cap(newCells), "!!!!!")
		}
		t.cells = newCells
		t.index++
	}
	return cell
}

func (t *Table) GetAt(col, row int) ICell {
	index := col + row*t.columns
	return t.Get(index)
}
func (t *Table) Get(index int) ICell {
	if index < 0 || index >= len(t.cells) {
		return nil
	} else {
		return t.cells[index]
	}
}
func (t *Table) ReplaceAt(col, row int, stringOrCell any) (ICell, error) {
	if col < 0 || row < 0 {
		return t, errors.New("negative col or row")
	}
	index := col + row*t.columns
	return t.Replace(index, stringOrCell)
}
func (t *Table) Replace(index int, stringOrCell any) (ICell, error) {
	// important: when modifieing table, clear all cached values
	t.cachedWidths = nil
	t.cachedHeights = nil
	t.cachedWidth = -1
	t.cachedHeight = -1
	if index < 0 || index >= t.index {
		return t, errors.New("index out of range")
	}

	cell := NewCell(stringOrCell)
	t.cells[index] = cell

	return t.cells[index], nil
}

// ----------------------------------------------
func (t *Table) CachedRuneDim() (width, height int) {
	if t.cachedWidth == -1 || t.cachedHeight == -1 {
		t.cachedWidth, t.cachedHeight = t.RuneDim()
	}
	return t.cachedWidth, t.cachedHeight
}
func (t *Table) RuneDim() (width, height int) {
	widths, heights := t.GetCachedWidthsAndHeights()
	width, height = 0, 0
	var col, row int
	var w, h int
	for col, w = range widths {
		width += w
		if t.borderConfig.GetSeparatorLeftOf(col, t.columns) != 0 {
			width++
		}
	}
	if t.borderConfig.GetSeparatorRightOf(col, t.columns) != 0 {
		width++
	}
	rows := t.GetNumberOfUsedRows()
	for row, h = range heights {
		height += h
		if t.borderConfig.GetSeparatorAbove(row, rows) != 0 {
			height++
		}
	}
	if t.borderConfig.GetSeparatorBelow(row, rows) != 0 {
		height++
	}

	return width, height
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func (t *Table) GetCachedWidthsAndHeights() ([]int, []int) {
	if t.cachedWidths == nil || t.cachedHeights == nil {
		// println("!!!!!calculating widths and heights!!!!!")
		t.cachedWidths, t.cachedHeights = t.CalcWidthsAndHeights()
	} else {
		// println("using cached widths and heights")
	}
	return t.cachedWidths, t.cachedHeights
}
func (t *Table) GetNumberOfUsedRows() int {
	usedRows := len(t.cells) / t.columns
	// handle incomplete rows
	if len(t.cells)%t.columns > 0 {
		usedRows++
	}
	return usedRows

}
func (t *Table) CalcWidthsAndHeights() ([]int, []int) {
	rows := t.GetNumberOfUsedRows()
	widths := make([]int, t.columns)
	heights := make([]int, rows)
	for index, cell := range t.cells {
		col := index % t.columns
		row := index / t.columns
		if cell != nil {
			w, h := cell.RuneDim()
			widths[col] = max(widths[col], w)
			heights[row] = max(heights[row], h)
		}
	}
	return widths, heights
}

func (t *Table) SmoothedRuneAt(x, y int) rune {
	r0 := t.RuneAt(x, y)
	b0 := IsBorderRune(r0)
	if !b0 {
		return r0
	} else {
		rAbove := t.RuneAt(x, y-1)
		rRight := t.RuneAt(x+1, y)
		rBelow := t.RuneAt(x, y+1)
		rLeft := t.RuneAt(x-1, y)
		bAbove := IsBorderRune(rAbove)
		bRight := IsBorderRune(rRight)
		bBelow := IsBorderRune(rBelow)
		bLeft := IsBorderRune(rLeft)

		switch {
		case bAbove && bRight && bBelow && bLeft:
			return '┼'
		case bAbove && bRight && bBelow && !bLeft:
			return '├'
		case bAbove && bRight && !bBelow && bLeft:
			return '┴'
		case bAbove && bRight && !bBelow && !bLeft:
			return '└'
		case bAbove && !bRight && bBelow && bLeft:
			return '┤'
		case bAbove && !bRight && bBelow && !bLeft:
			return '│'
		case bAbove && !bRight && !bBelow && bLeft:
			return '┘'
		case bAbove && !bRight && !bBelow && !bLeft:
			return '╵'
		case !bAbove && bRight && bBelow && bLeft:
			return '┬'
		case !bAbove && bRight && bBelow && !bLeft:
			return '┌'
		case !bAbove && bRight && !bBelow && bLeft:
			return '─'
		case !bAbove && bRight && !bBelow && !bLeft:
			return '╶'
		case !bAbove && !bRight && bBelow && bLeft:
			return '┐'
		case !bAbove && !bRight && bBelow && !bLeft:
			return '╷'
		case !bAbove && !bRight && !bBelow && bLeft:
			return '╴'
		case !bAbove && !bRight && !bBelow && !bLeft:
			return ' '
		default:
			return ' '
		}
	}
}

// ╴ ╵ ╶ ╷
// ┌─┬─┬─┐
// │H│d│r│
// ├─┼─┼─┤
// │F│t│r│
// └─┴─┴─┘
func (t *Table) RuneAt(x, y int) rune {
	/*
		Was muss ich wissen:
		Wie breit sind alle Spalten
		Wie hoch sind alle Zeilen
		Befinde ich mich in einer Zelle oder in einem Rahmen?
		Welcher Rahmen?
		Welche Zelle?
		Wo in der Zelle?

		Algorithmus:
		cx/cy heißt calculated-x/-y
		gehe in x-Richtung durch, bis cx/cy gleich x/y ist

		Die beiden Richtungen unabhängig behandeln
		1.) Breite aller Spalten, wo in welchem Rahmen, in welcher Spalte, relative Cell-X-Position
		2.) Höhe aller Zeilen, wo in welchem Rahmen, in welcher Zeile, relative Cell-Y-Position
	*/
	w, h := t.CachedRuneDim()
	if x < 0 || y < 0 || x >= w || y >= h {
		// out of range
		return ' '
	}
	cx := 0 // calculated-x, zum durchgehen durch die Zeile bis ich die richtige Spalte oder Separator gefunden habe
	dx := 0 // delta-x, relativ zur Zelle
	var selectedCol = -1
	var selectedColSep rune
	widths, heights := t.GetCachedWidthsAndHeights()

ColLoop:
	for col := 0; col < t.columns; col++ {
		// sep davor?
		sep := t.borderConfig.GetSeparatorLeftOf(col, t.columns)
		if sep != 0 {
			if cx == x {
				selectedColSep = sep
				break ColLoop
			}
			cx++
		}
		// Spalte durchlaufen
		for dx = 0; dx < widths[col]; dx++ {
			if cx == x {
				selectedCol = col
				break ColLoop
			}
			cx++
		}
	}
	if selectedColSep == -0 && selectedCol == -1 {
		// sep dahinter?
		sep := t.borderConfig.GetSeparatorLeftOf(t.columns, t.columns)
		if sep != 0 {
			if cx == x {
				selectedColSep = sep
			}
		}
	}
	cy := 0 // calculated-y, zum durchgehen durch die Spalte bis ich die richtige Zeile oder Separator gefunden habe
	dy := 0 // delta-y, relativ zur Zelle
	var selectedRow = -1
	var selectedRowSep rune

	rows := t.GetNumberOfUsedRows()
RowLoop:
	for row := 0; row < rows; row++ {
		// sep davor?
		sep := t.borderConfig.GetSeparatorAbove(row, rows)
		if sep != 0 {
			if cy == y {
				selectedRowSep = sep
				break RowLoop
			}
			cy++
		}
		// Zeile durchlaufen
		for dy = 0; dy < heights[row]; dy++ {
			if cy == y {
				selectedRow = row
				break RowLoop
			}
			cy++
		}
	}
	if selectedRowSep == -0 && selectedRow == -1 {
		// sep dahinter?
		sep := t.borderConfig.GetSeparatorAbove(rows, rows)
		if sep != 0 {
			if cy == y {
				selectedRowSep = sep
			}
		}
	}
	if selectedRowSep != 0 {
		return rune(selectedRowSep)
	}
	if selectedColSep != 0 {
		return rune(selectedColSep)
	}
	if selectedCol >= 0 && selectedRow >= 0 {
		cell := t.GetAt(selectedCol, selectedRow)
		// s := cell.String()
		if cell == nil {
			return ' '
		} else {
			r := cell.RuneAt(dx, dy)
			// println("selected: x/y", x, "/", y, "col/row", selectedCol, "/", selectedRow, "dx/dy:", dx, "/", dy, "rune:", string(r), "cell-string:", s)
			return r
		}
	} else {
		// println("!!!!!!!!!!RuneAt(x,y):", x, y, "out of range !!!!!!!!!!!!!")
		return ' '
		// panic("out of range")
	}
}

// ----------------------------------------------

// conveenient method for debugging, wastes memory on strings.Builder
func (t *Table) String() string {
	sb := strings.Builder{}
	t.WriteTo(&sb)
	return sb.String()
}

// used to write directly to machine.serial when running on microcontroller
// no need to waste memory with strings.Builder
func (t *Table) WriteTo(writer io.Writer) (int, error) {
	// calc max columns-width
	// calc total table-width (sum of column-width + column-separators)
	// calc max row-heights
	// calc total table-height (sum of row-heights + row-separators)
	// render rune-by-rune with surrounding runes
	bytesWritten := 0
	w, h := t.RuneDim()
	// fmt.Printf("t.RuneDim: %v/%v\n", w, h)
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			// r := t.RuneAt(x, y)
			r := t.SmoothedRuneAt(x, y)
			// fmt.Printf("RuneAt(%d/%d): %v\n", x, y, r)
			i, err := writer.Write([]byte(string(r)))
			bytesWritten += i
			if err != nil {
				return bytesWritten, err
			}
		}
		if y < h-1 {
			i, err := writer.Write([]byte("\n"))
			bytesWritten += i
			if err != nil {
				return bytesWritten, err
			}
		}

	}
	return bytesWritten, nil
}
