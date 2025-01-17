package texttable2

import "temp/texttable"

type DimList struct {
	dim       Dim
	direction Direction
	list      []IDim
}

func (dimList DimList) Dim() Dim {
	return dimList.dim
}
func (dimList DimList) PrintTo(pos Pos, matrix *texttable.RuneMatrix) {
	for _, dimElement := range dimList.list {
		dimElement.PrintTo(pos, matrix)
		pos = pos.Move(dimElement.Dim(), dimList.direction)
	}
}
func NewDimList(direction Direction) DimList {
	return DimList{dim: Dim{0, 0}, direction: direction}
}
func (dimList *DimList) Append(dimElement IDim) {
	dimList.list = append(dimList.list, dimElement)
	dimList.dim = dimList.dim.Enlarge(dimElement.Dim(), dimList.direction)
}
