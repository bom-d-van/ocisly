# ocisly

package ocisly contains a goroutine Wait function that could be used in testing environment.

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
Wait waits until you named gorouine finish.

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

Suggestion: to improve Wait accuacy, you could turn your inline/anonyomous function to a named function.

If your function case is too complicated, you could call
ocisly.PrintGoroutines() to help you figure out what name is your gorouinte
using.
```
