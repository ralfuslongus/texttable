package texttable

type ICell interface {
	RuneDim() (width, height int)
	RuneAt(x, y int) rune
	String() string
	SetAlignment(a Alignment) ICell
}
