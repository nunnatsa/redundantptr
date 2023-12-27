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

	t1 := t{}

	ddd := &t1
	ddd.b = &b // want "suspect redundant pointer"

	t1.a = &a // want "suspect redundant pointer"

	g := &t{
		a: &a, // want "suspect redundant pointer"
		b: &b, // want "suspect redundant pointer"
		c: &c, // want "suspect redundant pointer"
		t: &t2{
			p: &pi,  // want "suspect redundant pointer"
			t: &dur, // want "suspect redundant pointer"
		},
	}

	return g
}
