package texttable

import (
	"strings"
)

type Cell struct {
	lines           [][]rune
	maxWidthOfLines int
	alignment       Alignment
	separator       bool
}

func NewCell(v interface{}) *Cell {

	switch val := v.(type) {
	case *Cell:
		// Zelle so wie sie ist anhängen
		return val
	case nil:
		// Table ohne Ränder als string anhängen
		return NewCellFromString("").WithAlignment(LEFT)
	case *Table:
		// Table ohne Ränder als string anhängen
		return NewCellFromString(val.ToString(false)).WithAlignment(LEFT)
	case string:
		// Zelle mit gegebenem string erzeugen und anhängen
		return NewCellFromString(val).WithAlignment(LEFT)
	case bool:
		// Zelle mit geSprintetem val erzeugen und anhängen
		return NewCellFromString(BoolToString(val)).WithAlignment(RIGHT)
	case int:
		// Zelle mit geSprintetem val erzeugen und anhängen
		return NewCellFromString(IntToString(val)).WithAlignment(RIGHT)
	default:
		// Zelle mit geSprintetem val erzeugen und anhängen
		return NewCellFromString("!unsupported type!").WithAlignment(CENTER)
	}
}
func NewCellFromString(multiLine string) *Cell {
	// multiLine := val //fmt.Sprintf("%v", val)
	parts := strings.Split(multiLine, "\n")
	lines := make([][]rune, len(parts))
	maxWidthOfLines := 0
	for i, part := range parts {
		line := []rune(part)
		lines[i] = line
		maxWidthOfLines = max(maxWidthOfLines, len(line))
	}
	cell := Cell{lines, maxWidthOfLines, 0, false}
	return &cell
}
func (cell *Cell) WithAlignment(alignment Alignment) *Cell {
	cell.alignment = alignment
	return cell
}
func (cell *Cell) WithMaxWidthOfLines(max int) *Cell {
	cell.maxWidthOfLines = max
	return cell
}
func (cell *Cell) AsSeparator() *Cell {
	cell.separator = true
	return cell
}
func (cell *Cell) W() int {
	return cell.maxWidthOfLines
}
func (cell *Cell) H() int {
	return len(cell.lines)
}
func (cell *Cell) RenderToMatrix(x int, y int, w int, h int, m *RuneMatrix) {
	if cell.separator {
		for i := 0; i < w; i++ {
			for j := 0; j < h; j++ {
				m.Set(x+i, y+j, '┼')
			}
		}
	} else {
		for i, line := range cell.lines {
			var xOffset int
			switch cell.alignment {
			case LEFT:
				xOffset = 0
			case CENTER:
				xOffset = (w - len(line)) / 2
			case RIGHT:
				xOffset = w - len(line)
			}

			for j, r := range line {
				m.Set(x+xOffset+j, y+i, r)
			}
		}
	}
}
func (cell *Cell) String() string {
	m := NewRuneMatrix(cell.W(), cell.H())
	// m.FillAll('⋅')
	cell.RenderToMatrix(0, 0, cell.W(), cell.H(), &m)

	return m.String()
}
