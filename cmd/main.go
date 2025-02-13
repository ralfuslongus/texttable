package main

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/go-stack/stack"
	. "github.com/ralfuslongus/texttable"
)

func printInfos(t *Table) {
	fmt.Printf("w/h of table: %d/%d\n", t.W(), t.H())
}

func printColoredStacktrace(err interface{}) {
	// Farbcodes definieren
	redBold := color.New(color.FgRed, color.Bold).SprintFunc() // Rote, fette Schrift
	red := color.New(color.FgRed).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()

	// Fehler ausgeben
	fmt.Printf("%s %s\n", redBold("Error:"), red(err))

	// Stacktrace abrufen
	stackTrace := stack.Trace()

	// Zeilenweise durch den Stacktrace iterieren
	for _, stackCall := range stackTrace[4 : len(stackTrace)-2] {
		line := fmt.Sprintf("\tat %s:%d", stackCall.Frame().File, stackCall.Frame().Line)
		fmt.Println(blue(line)) // Link in Blau ausgeben
	}
}

func run() {

	// t1 := texttable.NewTable(2, 4)

	// t1.Add("A1", "A2")
	// t1.AddSeparatorsTillEndOfRow()
	// t1.Add("b1", "b2")
	// t1.Add("c1", "c2")
	// t1.Add("d1", "d2")
	// println()
	// printInfos(t1)
	// t1.Render(true, false)

	// t2 := texttable.NewTable(3, 2)
	// t2.Add("a", "bbb", "c")
	// t2.AddSeparatorsTillEndOfRow()
	// t2.Add("ddd", "e", t1)

	// printInfos(t2)
	// t2.Render(true, true)

	// t3 := texttable.NewTable(3, 2)
	// t3.Add("Über1", "Über2", "Preis")
	// t3.AddSeparatorsTillEndOfRow()
	// t3.Add("midA")
	// t3.Add("midB")
	// t3.Add(3)
	// t3.Add("midA")
	// t3.Add("midB", 4)
	// t3.Add("midA", "midB", 5)
	// t3.AddSeparatorsTillEndOfRow()
	// t3.Add("Unter1", "GESAMTPREIS:", 3+4+5)
	// t3.Render(true, true)

	t4 := NewTable(3, 2)

	t4.Add("LEFT", "CENTER", NewCell("RIGHT").WithAlignment(RIGHT))
	t4.Add(NewCell(1).WithAlignment(LEFT))
	t4.Add(NewCell(true).WithAlignment(LEFT))
	t4.Add(NewCell(2).WithAlignment(LEFT))
	t4.Add(NewCell("a").WithAlignment(RIGHT), NewCell("b").WithAlignment(CENTER), "c")
	t4.Render(true, true)
}

func main() {
	// Verwende defer und recover, um Panics abzufangen
	defer func() {
		if err := recover(); err != nil {
			printColoredStacktrace(err)
		}
	}()

	for i := 0; i < 1; i++ {
		run()
		time.Sleep(100 * time.Millisecond)
	}
}
