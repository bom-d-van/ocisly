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

// TODO: use exponential backoff

// Wait waits until you named gorouine finish.
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
// Suggestion: to improve Wait accuacy, you could turn your inline/anonyomous function to a named function.
//
// If your function case is too complicated, you could call
// ocisly.PrintGoroutines() to help you figure out what name is your gorouinte
// using.
func Wait(name string) {
	var errc int
	var sleep = IntervalBegin
	for {
		runtime.Gosched()
		time.Sleep(sleep)
		sleep *= 2

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
