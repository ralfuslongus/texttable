package main

import (
	"fmt"
	"texttable"

	"github.com/fatih/color"
	"github.com/go-stack/stack"
)

func printInfos(t *texttable.Table) {
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

func main() {
	// Verwende defer und recover, um Panics abzufangen
	defer func() {
		if err := recover(); err != nil {
			printColoredStacktrace(err)
		}
	}()

	// Hier wird ein Panic ausgelöst
	run()
}

func run() {

	t1 := texttable.NewTable(2, 4)

	t1.Add("A1", "A2")
	t1.AddSeparatorsTillEndOfRow()
	t1.Add("b1", "b2")
	t1.Add("c1", "c2")
	t1.Add("d1", "d2")
	println()
	printInfos(t1)
	t1.Render(true, false)

	// t.AddSeparatorsTillEndOfRow()
	// printInfos(t)

	// t.Add("bb1", "bb2", "bb3")
	// printInfos(t)

	// t.Add("ccc1", "ccc2", "ccc3")
	// t.AddSeparatorsTillEndOfRow()
	// printInfos(t)

	t2 := texttable.NewTable(3, 2)
	t2.Add("a", "bbb", "c")
	t2.AddSeparatorsTillEndOfRow()
	t2.Add("ddd", "e", t1)

	// println()
	printInfos(t2)
	t2.Render(true, true)

	t3 := texttable.NewTable(3, 2)
	t3.Add("Über1", "Über2", "Über3")
	t3.AddSeparatorsTillEndOfRow()
	t3.Add(t1, "Nix gewesen\naußer\nSpeesen!", true)
	t3.AddSeparatorsTillEndOfRow()
	t3.Add("leer", t1, t2)
	t3.Render(true, true)
}
