# ocisly

package ocisly contains a goroutine Wait function that could be used in testing environment.

Name inspired by SpaceX drone ship [Of Course I Still Love you](https://en.wikipedia.org/wiki/Autonomous_spaceport_drone_ship).

[![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](http://godoc.org/github.com/bom-d-van/ocisly)

A simple example

```go
func f() {}

func main() {
	go f()
	ocisly.Wait("github.com/bom-d-van/pkg.f")
}
```

Explanations:

```
Wait waits until you named gorouine finish. Using DefaultTimeout (initialized 10 seconds)

If the Wait exits before your goroutine started, you could specify a large
IntervalBegin to avoid this problem.

How to figure out your function name. Basically it's {{ImportPath}}.{{FuncName}}.
		Case 1: function named "Func" from package (same or other) named "github.com/alice/pkg"
				ocisly.Wait("github.com/alice/pkg.Func")
		Case 2: method named "method" in type "typ" from package (same or other) named "github.com/alice/pkg"
				ocisly.Wait("github.com/alice/pkg.(typ).(github.com/alice/pkg.method)")
		Case 3: method named "method" in type "typ" from package (same or other) named "github.com/alice/pkg"
				ocisly.Wait("github.com/alice/pkg.(typ).(github.com/alice/pkg.method)")
Running examples could be found in file ocisly_test.go.

Disclaimer: cases mentioned above isn't comprehensive, so if your function is
complicated to figure out by human brain, you could call
ocisly.PrintGoroutines() to help you figure out what name is your gorouinte using.

Suggestion: to improve Wait accuracy, you could turn your inline/anonymous function to a named function.
```
