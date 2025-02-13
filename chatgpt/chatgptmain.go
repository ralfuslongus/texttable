package main

import "fmt"

// Alignment ist ein benutzerdefinierter Typ für Alignment
type Alignment int

// Definiere einige Konstanten für Alignment
const (
	Left Alignment = iota
	Center
	Right
)

// Table repräsentiert eine einfache Tabelle, die Werte speichert
type Table struct {
	rows      [][]interface{}
	alignment Alignment // Variable für Alignment
}

// Add fügt neue Werte zur Tabelle hinzu
func (t *Table) Add(vals ...interface{}) {
	// Überprüfe, ob vals nicht leer ist
	if len(vals) > 0 {
		// Überprüfe, ob das letzte Element vom Typ Alignment ist
		if align, ok := vals[len(vals)-1].(Alignment); ok {
			t.alignment = align       // Weisen Sie die Variable alignment zu
			vals = vals[:len(vals)-1] // Entferne das letzte Element aus vals
		}
	}

	// Füge die verbleibenden Werte zur Tabelle hinzu
	t.rows = append(t.rows, vals)
}

// Print gibt die Werte der Tabelle aus
func (t *Table) Print() {
	for _, row := range t.rows {
		fmt.Println(row)
	}
}

// PrintAlignment gibt den aktuellen Alignment-Wert aus
func (t *Table) PrintAlignment() {
	fmt.Printf("Current Alignment: %d\n", t.alignment)
}

func main() {
	// Erstelle eine neue Tabelle
	table := &Table{}

	// Füge einige Werte hinzu, einschließlich eines Alignment-Werts
	table.Add(1, "Alice", 3.14, Left) // Left ist ein Alignment-Wert
	table.Add(2, "Bob", 2.71)
	table.Add(3, "Charlie", 1.41, Center) // Center ist ein Alignment-Wert

	// Gebe die Werte der Tabelle aus
	table.Print()

	// Gebe den aktuellen Alignment-Wert aus
	table.PrintAlignment()
}
