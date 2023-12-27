package samplecode

import (
	"time"
)

func createT() *t {
	var a = "a"
	b := 2
	dur := time.Minute * 5

	if b > 1 {
		a = "A"
	}

	return &t{
		a: &a,
		b: &b,
		c: &c,
		t: &t2{
			p: &pi,
			t: &dur,
		},
	}
}
