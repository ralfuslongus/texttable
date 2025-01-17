package texttable2

import (
	"fmt"
	"temp/texttable"
)

type Direction bool

const (
	VERTICAL   Direction = false
	HORIZONTAL Direction = true
)

type IDim interface {
	Dim() Dim
	PrintTo(pos Pos, matrix *texttable.RuneMatrix)
}

type Pos struct {
	X, Y int
}

func (pos Pos) String() string {
	return fmt.Sprintf("x/y: %d/%d", pos.X, pos.Y)
}
func (pos Pos) Move(dim Dim, direction Direction) Pos {
	if direction == HORIZONTAL {
		return Pos{pos.X + dim.W, pos.Y}
	} else {
		return Pos{pos.X, pos.Y + dim.H}
	}
}

type Dim struct {
	W, H int
}

func (dim Dim) String() string {
	return fmt.Sprintf("w/h: %d/%d", dim.W, dim.H)
}
func (dim Dim) Enlarge(other Dim, direction Direction) Dim {
	if direction == HORIZONTAL {
		return Dim{dim.W + other.W, max(dim.H, other.H)}
	} else {
		return Dim{max(dim.W, other.W), dim.H + other.H}
	}
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
