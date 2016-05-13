package ocisly

import (
	"testing"
	"time"

	"github.com/bom-d-van/ocisly/testdata/pkg"
	"github.com/bom-d-van/sidekick"
)

type s struct{}

func (*s) method() {
	time.Sleep(time.Millisecond * 100)
	pkg.Love = true
}

func function() {
	time.Sleep(time.Millisecond * 200)
	pkg.Love = true
}

func TestWait(t *testing.T) {
	inline := func() {
		time.Sleep(time.Millisecond * 100)
		pkg.Love = true
	}

	var i s
	cases := []struct {
		f    func()
		name string
	}{
		0: {f: pkg.Nop, name: "github.com/bom-d-van/ocisly/testdata/pkg.Nop"},
		1: {f: inline, name: "github.com/bom-d-van/ocisly.TestWait.func1"},
		2: {f: function, name: "github.com/bom-d-van/ocisly.function"},
		3: {f: i.method, name: "github.com/bom-d-van/ocisly.(*s).(github.com/bom-d-van/ocisly.method)"},
	}

	for i, c := range cases {
		if sidekick.SkipCase(i) {
			continue
		}

		pkg.Love = false
		go func() {
			c.f()
		}()
		Wait(c.name)
		if !pkg.Love {
			t.Errorf("case %d: pkg.Love = %t; want true", i, pkg.Love)
		}
	}
}

// func TestSuggestName(t *testing.T) {
// 	SuggestName()
// }
