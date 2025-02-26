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
