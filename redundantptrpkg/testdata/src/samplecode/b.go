package samplecode

import (
	"math"
	"time"
)

const (
	x = 4.14
)

var (
	c  = x
	pi = math.Pi
)

type t2 struct {
	p *float64
	t *time.Duration
}

type t struct {
	a *string
	b *int
	c *float64
	t *t2
}
