// package ocisly contains a goroutine Wait function that could be used in
// testing environment.
//
// Name inspired by SpaceX drone ship Of Course I Still Love you
// (https://en.wikipedia.org/wiki/Autonomous_spaceport_drone_ship).
//
package ocisly

import (
	"bytes"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

var IntervalBegin = time.Millisecond * 30
var DefaultTimeout = time.Second * 10

// TODO: use exponential backoff

// Wait waits until you named gorouine finish. Using DefaultTimeout (initialized 10 seconds)
//
// If the Wait exits before your goroutine started, you could specify a large
// IntervalBegin to avoid this problem.
//
// How to figure out your function name. Basically it's {{ImportPath}}.{{FuncName}}.
// 		Case 1: function named "Func" from package (same or other) named "github.com/alice/pkg"
// 				ocisly.Wait("github.com/alice/pkg.Func")
// 		Case 2: method named "method" in type "typ" from package (same or other) named "github.com/alice/pkg"
// 				ocisly.Wait("github.com/alice/pkg.(typ).(github.com/alice/pkg.method)")
// 		Case 3: method named "method" in type "typ" from package (same or other) named "github.com/alice/pkg"
// 				ocisly.Wait("github.com/alice/pkg.(typ).(github.com/alice/pkg.method)")
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

// WaitTimeout works same as Wait, but suppots timeout specification.
func WaitTimeout(name string, timeout time.Duration) {
	var errc int
	var sleep = IntervalBegin
	for {
		runtime.Gosched()
		time.Sleep(sleep)
		sleep *= 2

		if sleep > timeout {
			println("ocisly: timeout")
			return
		}

		// PrintGoroutines()

		p := pprof.Lookup("goroutine")
		buf := bytes.NewBuffer(make([]byte, 1<<20)) // size: 1MB
		if err := p.WriteTo(buf, 1); err != nil {
			if errc > 3 {
				println("ocisly: faile to write profile")
				return
			}

			errc++
			continue
		}
		// TODO: " "+name+"+" doesn't work. Why?
		if !bytes.Contains(buf.Bytes(), []byte(name)) {
			break
		}
	}
}

// PrintGoroutines prints out all the goroutine stack.
func PrintGoroutines() {
	p := pprof.Lookup("goroutine")
	p.WriteTo(os.Stdout, 1)
}

// TODO: return suggest Wait names.
// func Suggest() string
