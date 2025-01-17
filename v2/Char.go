package texttable2

import "temp/texttable"

type Char rune

func (char Char) Dim() Dim {
	return Dim{1, 1}
}
func (char Char) PrintTo(pos Pos, matrix *texttable.RuneMatrix) {
	matrix.Set(pos.X, pos.Y, rune(char))
}
