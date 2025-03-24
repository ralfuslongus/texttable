package texttable

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/go-stack/stack"
)

func PrintColoredStacktraceOnError() {
	if err := recover(); err != nil {
		// Farbcodes definieren
		redBold := color.New(color.FgRed, color.Bold).SprintFunc() // Rote, fette Schrift
		red := color.New(color.FgRed).SprintFunc()
		blue := color.New(color.FgBlue).SprintFunc()

		// Fehler ausgeben
		fmt.Printf("%s %s\n", redBold("Error:"), red(err))

		// Stacktrace abrufen
		stackTrace := stack.Trace()

		// Zeilenweise durch den Stacktrace iterieren
		for _, stackCall := range stackTrace[3 : len(stackTrace)-2] {
			line := fmt.Sprintf("\tat %s:%d", stackCall.Frame().File, stackCall.Frame().Line)
			fmt.Println(blue(line)) // Link in Blau ausgeben
		}
	}
}
