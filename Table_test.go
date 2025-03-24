package texttable

import (
	"strings"
	"testing"
)

/*
Was soll Table alles können:
1.) So wenig Speicher wie möglich verbrauchen (feather_m0: 32kB RAM!, Table bisher: ~20kB)
	- keine zwischenstrings, direkt auf io.Writer/machine.Serial schreiben
	- bytes statt runes speichern, mit DynamicCodepage, auf 255 Zeichen beschränkt
	- Rahmen nicht zeichenweise speichern, nur wo
2.) Hierarchy: Tables als Cells in anderen Tables möglich
	- ICell interface (RuneDim(), RuneAt(x,y))
3.) Dynamische Rahmen, verbinden sich mit umliegenden Rahmen
	- Umgebende Rahmen/runes mit einbeziehen (wie?), auch zwischen den Hierarchien (Table als Cell)
4.) Änderbarkeit
	- Set(col,row)
	- Clear(col/row)
	- SetColSep(col,rune)
	- SetRowSep(row,rune)
	- Gesamtgröße immer neu berechnen
*/

func TestDimChangesWithSeparators(testing *testing.T) {
	t := NewTable(2, 2, NoBorders)
	// 2x2-table with no separators
	t.Append("a")
	t.Append("b")
	t.Append("c")
	t.Append("d")

	// dim with single-char-cells = 2x2
	// without any separators
	w, h := t.RuneDim()
	AssertEqual(testing, w, 2)
	AssertEqual(testing, h, 2)

	// with all borders
	t.borderConfig = AllBorders
	w, h = t.RuneDim()
	AssertEqual(testing, w, 5)
	AssertEqual(testing, h, 5)
}
func TestGesamt(testing *testing.T) {
	var w, h int
	t := NewTable(2, 3, NoBorders)
	t.Append("c0")
	t.Append("c1")
	t.Append("c2")
	// t.Append("")

	AssertEqual(testing, t.GetAt(0, 0).String(), "c0")
	AssertEqual(testing, t.GetAt(1, 0).String(), "c1")
	AssertEqual(testing, t.GetAt(0, 1).String(), "c2")
	AssertEqual(testing, t.GetAt(1, 1), nil)

	AssertEqual(testing, t.RuneAt(0, 0), 'c')
	AssertEqual(testing, t.RuneAt(1, 0), '0')
	AssertEqual(testing, t.RuneAt(2, 0), 'c')
	AssertEqual(testing, t.RuneAt(3, 0), '1')

	AssertEqual(testing, t.RuneAt(0, 1), 'c')
	AssertEqual(testing, t.RuneAt(1, 1), '2')
	AssertEqual(testing, t.RuneAt(2, 1), ' ')
	AssertEqual(testing, t.RuneAt(3, 1), ' ')

	sb := strings.Builder{}
	t.WriteTo(&sb)
	AssertEqual(testing, sb.String(), "c0c1\nc2  ")

	// Change cells
	t.ReplaceAt(0, 0, "A")
	t.ReplaceAt(1, 0, "B")
	t.ReplaceAt(0, 1, "C")
	// t.ReplaceAt(1, 1, "D")
	AssertEqual(testing, t.GetAt(0, 0).String(), "A")
	AssertEqual(testing, t.GetAt(1, 0).String(), "B")
	AssertEqual(testing, t.GetAt(0, 1).String(), "C")
	// AssertEqual(testing, t.GetAt(1, 1).String(), "D")
	w, h = t.RuneDim()
	AssertEqual(testing, w, 2)
	AssertEqual(testing, h, 2)
	AssertEqual(testing, t.String(), "AB\nC ")

	// Changes of Dim by setting separators
	// t.SetColSep(0, SINGLE)
	// w, h = t.RuneDim()
	// AssertEqual(testing, w, 5)
	// AssertEqual(testing, h, 2)
	// AssertEqual(testing, t.String(), "|AB\n|CD")

	// t.SetColSep(1, SINGLE)
	// w, h = t.RuneDim()
	// AssertEqual(testing, w, 6)
	// AssertEqual(testing, h, 2)
	// AssertEqual(testing, t.String(), "|A|B\n|C|D")

	// t.SetColSep(2, SINGLE)
	// w, h = t.RuneDim()
	// AssertEqual(testing, w, 7)
	// AssertEqual(testing, h, 2)
	// AssertEqual(testing, t.String(), "|A|B|\n|C|D|")

	// t.SetRowSep(0, SINGLE)
	// w, h = t.RuneDim()
	// AssertEqual(testing, w, 7)
	// AssertEqual(testing, h, 3)
	// AssertEqual(testing, t.String(), "-----\n|A|B|\n|C|D|")

	// t.SetRowSep(1, SINGLE)
	// w, h = t.RuneDim()
	// AssertEqual(testing, w, 7)
	// AssertEqual(testing, h, 4)
	// AssertEqual(testing, t.String(), "-----\n|A|B|\n-----\n|C|D|")

	// t.SetRowSep(3, SINGLE)
	// w, h = t.RuneDim()
	// AssertEqual(testing, w, 7)
	// AssertEqual(testing, h, 5)
	// AssertEqual(testing, t.String(), "-----\n|A|B|\n-----\n|C|D|\n\n-----")

	// // remove all separators, H&V
	// t.SetColSep(0, 0).SetColSep(1, 0).SetColSep(2, 0).SetRowSep(0, 0).SetRowSep(1, 0).SetRowSep(2, 0)
	// w, h = t.RuneDim()
	// AssertEqual(testing, w, 4)
	// AssertEqual(testing, h, 2)
	// AssertEqual(testing, t.String(), "AB\nCD")

}
