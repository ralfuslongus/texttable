package main

import (
	"fmt"
	texttable3 "temp/texttable/v3"

	"github.com/fatih/color"
	"github.com/go-stack/stack"
)

func main2() {
	t := texttable3.NewTable(3, 1, texttable3.ALL)
	t.Add("a1")
	t.Add("a2")
	t.Add("a3")
	fmt.Printf("w/h of table: %d/%d\n", t.W(), t.H())

	t.Add("bb1")
	t.Add("bb2")
	t.Add("bb3")
	fmt.Printf("w/h of table: %d/%d\n", t.W(), t.H())

	t.Add("ccc1")
	t.Add("ccc2")
	t.Add("ccc3")
	fmt.Printf("w/h of table: %d/%d\n", t.W(), t.H())

	// t.Add("cccc\n1")
	// t.Add("dddd\n2")
	// t.Add("eeee\n3")
	// fmt.Printf("w/h of table: %d/%d\n", t.W(), t.H())

	// t.Add(nil)
	// t.Add("x")
	// fmt.Printf("w/h of table: %d/%d\n", t.W(), t.H())

	println()
	t.Render()
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
	main2()
}
