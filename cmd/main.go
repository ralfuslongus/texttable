package main

import (
	"os"
	"time"

	. "github.com/ralfuslongus/texttable"
)

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
	t4.WriteTo(os.Stdout)

}

func main() {
	defer PrintColoredStacktraceOnError()

	for i := 0; i < 1; i++ {
		run()
		time.Sleep(100 * time.Millisecond)
	}
}
