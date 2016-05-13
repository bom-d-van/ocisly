package pkg

import "time"

var Love bool

func Nop() {
	time.Sleep(time.Millisecond * 100)
	Love = true
}
