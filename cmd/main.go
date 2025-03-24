package main

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/go-stack/stack"
	. "github.com/ralfuslongus/texttable"
)

func printInfos(t *Table) {
	w, h := t.CachedRuneDim()
	fmt.Printf("w/h of table: %d/%d\n", w, h)
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
	t4 := NewTable(3, 2, AllBorders)
	t4.Append("LEFT").SetAlignment(LEFT)
	t4.Append("CENTER").SetAlignment(CENTER)
	t4.Append("RIGHT___________________").SetAlignment(RIGHT)

	// evtl. TODO: BorderConfig ändern damit einzelne Col/Row-Separators gesetzte werden können
	//t4.AppendSeparatorsTillEndOfRow()

	t4.Append(1).SetAlignment(LEFT)
	t4.Append(true).SetAlignment(CENTER)
	t4.Append(2).SetAlignment(RIGHT)

	// evtl. TODO: BorderConfig ändern damit einzelne Col/Row-Separators gesetzte werden können
	//t4.AppendSeparatorsTillEndOfRow()

	t4.Append("Multiline-\nString-\nNr 1").SetAlignment(LEFT)
	t4.Append("Multiline-\nString-\nNr 2").SetAlignment(CENTER)
	t4.Append("Multiline-\nString-\nNr 3").SetAlignment(RIGHT)
	result := t4.String()
	println(result)
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
