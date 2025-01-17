package main

import (
	"fmt"
	"os"
	"temp/texttable"
	texttable2 "temp/texttable/v2"
)

func main() {

	r1 := texttable2.Char('H')
	r2 := texttable2.Char('a')
	r3 := texttable2.Char('l')
	r4 := texttable2.Char('l')
	r5 := texttable2.Char('o')

	r6 := texttable2.Char('d')
	r7 := texttable2.Char('u')

	line1 := texttable2.NewDimList(texttable2.HORIZONTAL)
	line2 := texttable2.NewDimList(texttable2.HORIZONTAL)
	line1.Append(r1)
	line1.Append(r2)
	line1.Append(r3)
	line1.Append(r4)
	line1.Append(r5)
	line2.Append(r6)
	line2.Append(r7)

	cell := texttable2.NewDimList(texttable2.VERTICAL)
	cell.Append(line1)
	cell.Append(line2)

	row1 := texttable2.NewDimList(texttable2.HORIZONTAL)
	row2 := texttable2.NewDimList(texttable2.HORIZONTAL)
	row1.Append(cell)
	row1.Append(texttable2.Char(' '))
	row1.Append(texttable2.Char(' '))
	row1.Append(r4)
	row2.Append(cell)
	row2.Append(cell)
	row2.Append(cell)
	row2.Append(r2)
	row2.Append(r2)
	row2.Append(r2)

	table := texttable2.NewDimList(texttable2.VERTICAL)
	table.Append(row1)
	table.Append(row2)
	table.Append(r1)
	table.Append(r1)
	table.Append(r1)

	fmt.Println("r1", r1.Dim())
	fmt.Println("line", line1.Dim())
	fmt.Println("cell", cell.Dim())
	fmt.Println("row", row1.Dim())
	fmt.Println("table", table.Dim())
	fmt.Println("-------------------")

	matrix := texttable.NewRuneMatrix(table.Dim().W+1, table.Dim().H)
	table.PrintTo(texttable2.Pos{X: 0, Y: 0}, &matrix)
	matrix.RenderTo(os.Stdout)

}
