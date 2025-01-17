package main

import (
	"fmt"
	"os"
	"strings"
	"temp/texttable"
	"unicode/utf8"
)

// --------------------------------------
type Box interface {
	Dim() (int, int)
	RenderTo(int, int, *texttable.RuneMatrix)
}

// --------------------------------------
type HBox []Box

func (hBox HBox) Dim() (int, int) {
	gesamtW := 0
	maxH := 0
	for _, box := range hBox {
		w, h := box.Dim()
		// Alle Breiten summieren sich zur Gesamtbreite auf
		gesamtW += w
		// Die Höhe ist das Maximum aller Boxen
		if h > maxH {
			maxH = h
		}
	}
	return gesamtW, maxH
}
func (hBox HBox) RenderTo(x int, y int, m *texttable.RuneMatrix) {
	for _, box := range hBox {
		w, _ := box.Dim()
		box.RenderTo(x, y, m)
		x += w
	}
}

// --------------------------------------
type VBox struct {
	boxes            []Box
	withInnerBorders bool
}

func (vBox VBox) Dim() (int, int) {
	gesamtH := 0
	maxW := 0
	for _, box := range vBox.boxes {
		w, h := box.Dim()
		// Die Breite ist das Maximum aller Boxen
		if w > maxW {
			maxW = w
		}
		// Alle Höhen summieren sich zur Gesamthöhe auf
		gesamtH += h
	}
	return maxW, gesamtH
}
func (vBox *VBox) Append(box Box) {
	vBox.boxes = append(vBox.boxes, box)
}

func (vBox VBox) RenderTo(x int, y int, m *texttable.RuneMatrix) {
	for _, box := range vBox.boxes {
		_, h := box.Dim()
		box.RenderTo(x, y, m)
		y += h
	}
}

// ------ Line soll nur eine Zeile ohne Umbrüche haben -------
type Line struct {
	Content   string
	Length    int
	Alignment Alignment
}

func (line Line) Dim() (int, int) {
	return line.Length, 1
}
func (line Line) RenderTo(x int, y int, m *texttable.RuneMatrix) {
	for i, r := range line.Content {
		m.Set(x+i, y, r)
	}
}

func (l Line) String() string {
	return l.Content + " (alignment=" + string(l.Alignment+")")
}

func LineFrom(input interface{}) Line {
	alignment := left
	// TrimWhitespace entfernt führende und nachfolgende Leerzeichen
	switch input.(type) {
	case int:
		alignment = right
	case float32:
		alignment = right
	case float64:
		alignment = right
	case bool:
		alignment = right
	}
	s := fmt.Sprintf("%v", input)
	trimmed := strings.TrimSpace(s)
	length := utf8.RuneCountInString(trimmed)
	line := Line{Content: trimmed, Length: length, Alignment: alignment}
	return line
}
func LinesFrom(multiLine string) VBox {
	parts := strings.Split(multiLine, "\n")
	// lines := make(VBox, 0, 1)
	lines := VBox{}

	// Entferne Zeilenumbrüche und gib das Ergebnis aus
	for _, part := range parts {
		line := LineFrom(part)
		lines.Append(line)
	}
	return lines
}

// --------------------------------------
type Alignment string

const (
	left   Alignment = "L"
	right  Alignment = "R"
	center Alignment = "C"
)

// --------------------------------------
func PrintDim(box Box) {
	w, h := box.Dim()
	fmt.Printf("DIM(w/h) of %T: %d/%d, value: %v\n", box, w, h, box)
}

// --------------------------------------
func main() {
	line := LineFrom("hallo")
	PrintDim(line)

	cell1 := LinesFrom("Einzeiler")
	cell1.Dim()
	cell2 := LinesFrom("Zwei-\nzeiler")
	cell2.Dim()
	PrintDim(&cell1)
	PrintDim(&cell2)

	// hBox := HBox{}
	// hBox = append(hBox, line)
	// hBox = append(hBox, cell1)
	// hBox = append(hBox, cell2)
	// PrintDim(hBox)

	vBox := VBox{withInnerBorders: true}
	vBox.Append(line)
	vBox.Append(cell1)
	vBox.Append(cell2)
	PrintDim(vBox)

	m := texttable.NewRuneMatrix(50, 10)
	vBox.RenderTo(4, 4, &m)
	m.RenderTo(os.Stdout)

}

func main2() {
	line := LineFrom("hhh")
	PrintDim(line)
	m := texttable.NewRuneMatrix(50, 10)
	m.VerticalLineAt(0)
	m.VerticalLineAt(9)
	m.HorizontalLineAt(0)
	m.HorizontalLineAt(2)
	line.RenderTo(1, 1, &m)
	m.RenderTo(os.Stdout)
}
