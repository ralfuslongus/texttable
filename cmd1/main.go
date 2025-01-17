package main

import (
	"fmt"
	"os"
	"temp/texttable"
)

func main() {

	m := texttable.NewRuneMatrix(3, 3)
	m.HorizontalLineAt(0)
	m.HorizontalLineAt(2)
	m.VerticalLineAt(0)
	m.VerticalLineAt(2)
	w, h := texttable.WriteString(&m, 1, 1, "")
	// w, h := texttable.WriteString(&m, 1, 1, "x")
	fmt.Printf("w/h: %d/%d\n", w, h)
	m.SmoothOpenCrossEnds()
	m.RenderTo(os.Stdout)

	// t := texttable.NewTable(2)
	// t.Add("1", "2")
	// t.Add("a", "b")
	// t.Add("\nccc", "hallo\n neuen\n  Zeile")
	// t.Add("ccc", "房")
	// t.RenderTo(os.Stdout)
}
