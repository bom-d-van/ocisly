// package ocisly contains a goroutine Wait function that could be used in
// testing environment.
//
// Name inspired by SpaceX drone ship Of Course I Still Love you
// (https://en.wikipedia.org/wiki/Autonomous_spaceport_drone_ship).
//
package ocisly

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

var IntervalBegin = time.Millisecond * 30
var DefaultTimeout = time.Second * 10

// TODO: use exponential back-off

// Wait waits until you named gorouine finish. Using DefaultTimeout (initialized 10 seconds)
//
// If the Wait exits before your goroutine started, you could specify a large
// IntervalBegin to avoid this problem.
//
// How to figure out your function name. Basically it's {{ImportPath}}.{{FuncName}}.
// 		Case 1: function named "Func" from package (same or other) named "github.com/alice/pkg"
// 				ocisly.Wait("github.com/alice/pkg.Func")
// 		Case 2: method named "method" in type "*typ" from package (same or other) named "github.com/alice/pkg"
// 				ocisly.Wait("github.com/alice/pkg.(*typ).method")
// 		Case 3: first anonymous function called "foo" in function/method named "Func" in package (same or other)
// 				named "github.com/alice/pkg".
// 				ocisly.Wait("github.com/alice/pkg.Func.func1")
// Running examples could be found in file ocisly_test.go.
//
// Disclaimer: cases mentioned above isn't comprehensive, so if your function is
// complicated to figure out by human brain, you could call
// ocisly.PrintGoroutines() to help you figure out what name is your gorouinte using.
//
// Suggestion: to improve Wait accuracy, you could turn your inline/anonymous function to a named function.
func Wait(name string) {
	WaitTimeout(name, DefaultTimeout)
}

// WaitTimeout works same as Wait, but supports timeout specification.
func WaitTimeout(name string, timeout time.Duration) {
	// var errc int
	var sleep = IntervalBegin

wait:
	for {
		runtime.Gosched()
		time.Sleep(sleep)
		sleep *= 2

		if sleep > timeout {
			println("[ocisly] error: timeout")
			return
		}

		// PrintGoroutines()

		var srs = make([]runtime.StackRecord, runtime.NumGoroutine())
		_, ok := runtime.GoroutineProfile(srs)
		if !ok {
			continue
		}

		for _, sr := range srs {
			for _, pc := range sr.Stack() {
				f := runtime.FuncForPC(pc)
				// log.Printf("--> %+v\n", f.Name())
				if f.Name() == name {
					continue wait
				}
			}
		}

		break
	}
}

// PrintGoroutines prints out all the goroutine stack.
func PrintGoroutines() {
	p := pprof.Lookup("goroutine")
	p.WriteTo(os.Stdout, 1)
}

// PrintSuggestions prints function/method name suggestions that could be used in
// Wait function.
func PrintSuggestions() {
	for i := 1; ; i++ {
		pc, _, _, ok := runtime.Caller(i)
		if !ok {
			break
		}
		fn := runtime.FuncForPC(pc)
		if fn == nil {
			continue
		}
		fmt.Printf("[ocisly] suggestion %d: %s\n", i, fn.Name())
	}
}

// TODO: return suggest Wait names.
// func Suggest() string
