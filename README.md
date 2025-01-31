!!!!!!! WARNING !!!!!!!
PRE-PRE-ALPHA-VERSION
DO NOT USE, JUST FOR TESTING AND GET USED TO GO AND GITHUB

# texttable
Texttable lets you create simple tables with unicode-separators like this:
```
┌─────┬───────────┬─────────────┐
│Über1│Über2      │Über3        │
├──┬──┼───────────┼─────────────┤
│A1│A2│Nix gewesen│true         │
├──┼──┤außer      │             │
│b1│b2│Speesen!   │             │
│c1│c2│           │             │
│d1│d2│           │             │
├──┴──┼──┬────────┼───┬───┬─────┤
│leer │A1│A2      │a  │bbb│c    │
│     ├──┼──      ├───┼───┼──┬──┤
│     │b1│b2      │ddd│e  │A1│A2│
│     │c1│c2      │   │   ├──┼──┤
│     │d1│d2      │   │   │b1│b2│
│     │           │   │   │c1│c2│
│     │           │   │   │d1│d2│
└─────┴───────────┴───┴───┴──┴──┘
```

## Install
```bash
go get github.com/ralfuslongus/texttable
```

## Examples
```go
	t1 := texttable.NewTable(2, 4)
	t1.Add("A1", "A2")
	t1.AddSeparatorsTillEndOfRow()
	t1.Add("b1", "b2")
	t1.Add("c1", "c2")
	t1.Add("d1", "d2")
	t1.Render(true, false)
```
