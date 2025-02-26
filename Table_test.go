package texttable

import (
	"testing"
)

// ANSI Escape Codes für Farben
const (
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Reset  = "\033[0m"
)

func TestMultiLineCell(t *testing.T) {
	// Eine Cell mit mehreren Zeilen
	multiLineCell := NewCell("Line1\nLine2")
	result := multiLineCell.String()
	expected := "Line1\nLine2"

	AssertEqual(t, result, expected)
}

func TestManualRuneMatrix(t *testing.T) {
	// Eine RuneMatrix manuell beschreiben
	m := NewRuneMatrix(3, 2)
	m.Set(0, 0, 'a')
	m.Set(1, 0, 'b')
	m.Set(2, 0, 'c')
	m.Set(0, 1, 'A')
	m.Set(1, 1, 'B')
	m.Set(2, 1, 'C')
	result := m.String()
	expected := "abc\nABC"

	AssertEqual(t, result, expected)
}

func TestSimpleTableWithoutOuterFrame(t *testing.T) {
	// Ein einfacher Table ohne alles
	t1 := NewTable(2, 2)
	t1.Add("v1", "v2")
	t1.Add("v3", "v4")
	result := t1.ToString(false)
	expected := "v1│v2\nv3│v4"

	AssertEqual(t, result, expected)
}

func TestSimpleTableWithHeaderSeparatorAndOuterFrame(t *testing.T) {
	// Ein einfacher Table ohne Rand, mit Header-Separatoren
	t1 := NewTable(2, 4)
	t1.Add("A1", "A2")
	t1.AddSeparatorsTillEndOfRow()
	t1.Add("b1", "b2")
	t1.Add("c1", "c2")
	t1.Add("d1", "d2")
	result := t1.ToString(false)
	expected := `A1│A2
──┼──
b1│b2
c1│c2
d1│d2`
	AssertEqual(t, result, expected)
}
func TestNestedTables(t *testing.T) {
	// Ein einfacher Table ohne Rand, mit Header-Separatoren

	t1 := NewTable(2, 2)
	t1.Add("a", "b")
	t1.AddSeparatorsTillEndOfRow()
	t1.Add("c", "d")

	// Ein geschachtelter Table mit Rand
	t2 := NewTable(2, 2)
	t2.Add("A", "B")
	t2.Add("C", t1)
	result := t2.ToString(true)
	expected := `┌─┬───┐
│A│B  │
│C│a│b│
│ ├─┼─┤
│ │c│d│
└─┴─┴─┘`

	AssertEqual(t, result, expected)
}
func TestTableWithHeaderAndFooterSeparator(t *testing.T) {
	// Ein einfacher Table mit Rand, Header- und Footer-Separatoren
	t3 := NewTable(3, 2)
	t3.Add("Über1", "Über2", "Preis")
	t3.AddSeparatorsTillEndOfRow()
	t3.Add("midA")
	t3.Add("midB")
	t3.Add(3)
	t3.Add("midA")
	t3.Add("midB", 4)
	t3.Add("midA", "midB", 5)
	t3.AddSeparatorsTillEndOfRow()
	t3.Add("Unter1", "GESAMTPREIS:", 3+4+5)
	result := t3.ToString(true)
	expected := `┌──────┬────────────┬─────┐
│Über1 │Über2       │Preis│
├──────┼────────────┼─────┤
│midA  │midB        │    3│
│midA  │midB        │    4│
│midA  │midB        │    5│
├──────┼────────────┼─────┤
│Unter1│GESAMTPREIS:│   12│
└──────┴────────────┴─────┘`

	AssertEqual(t, result, expected)
}
func TestTableWithManualAlignmentsAndMaxWidth(t *testing.T) {
	t4 := NewTable(3, 2)
	t4.Add(NewCell("LEFT").
		WithAlignment(LEFT).
		WithMaxWidthOfLines(20))
	t4.Add(NewCell("CENTER").
		WithAlignment(CENTER).
		WithMaxWidthOfLines(20))
	t4.Add(NewCell("RIGHT").
		WithAlignment(RIGHT).
		WithMaxWidthOfLines(20))

	t4.AddSeparatorsTillEndOfRow()

	t4.Add(NewCell(1).
		WithAlignment(LEFT))
	t4.Add(NewCell(true).
		WithAlignment(CENTER))
	t4.Add(NewCell(2).
		WithAlignment(RIGHT))

	t4.AddSeparatorsTillEndOfRow()

	t4.Add(NewCell("Multiline-\nString-\nNr 1").
		WithAlignment(LEFT))
	t4.Add(NewCell("Multiline-\nString-\nNr 2").
		WithAlignment(CENTER))
	t4.Add(NewCell("Multiline-\nString-\nNr 3").
		WithAlignment(RIGHT))
	result := t4.ToString(true)
	expected := `┌────────────────────┬────────────────────┬────────────────────┐
│LEFT                │       CENTER       │               RIGHT│
├────────────────────┼────────────────────┼────────────────────┤
│1                   │        true        │                   2│
├────────────────────┼────────────────────┼────────────────────┤
│Multiline-          │     Multiline-     │          Multiline-│
│String-             │      String-       │             String-│
│Nr 1                │        Nr 2        │                Nr 3│
└────────────────────┴────────────────────┴────────────────────┘`

	AssertEqual(t, result, expected)
}
func TestGrowingTable(t *testing.T) {
	t5 := NewTable(2, 2)
	t5.Add("a")
	t5.Add("b")
	t5.Add("c")
	t5.Add("d")

	AssertEqual(t, t5.ToString(false), "a│b\nc│d")
	t5.Add("e")
	t5.Add("f")
	t5.Add("g")
	t5.Add("h")
	AssertEqual(t, t5.ToString(false), "a│b\nc│d\ne│f\ng│h")
}
func TestModdingTable(t *testing.T) {
	t5 := NewTable(2, 2)
	t5.Add("a")
	t5.Add("b")
	t5.Add("c")
	t5.Add("d")

	AssertEqual(t, t5.ToString(false), "a│b\nc│d")
	t5.Set(0, 0, "AAAA")
	AssertEqual(t, t5.ToString(false), "AAAA│b\nc   │d")
}
