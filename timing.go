package main

import (
	"time"
)

type Timing struct {
	// timer
	time.Time

	// all passed time
	dur time.Duration
}

func (t *Timing) Start() {
	t.Time = time.Now()
}

func (t *Timing) Passed() time.Duration {
	dur := time.Since(t.Time)
	t.dur = t.dur + dur
	t.Time = time.Now()

	return dur
}

func (t *Timing) End() time.Duration {
	return time.Since(t.Time)
}

func (t *Timing) Duration() time.Duration {
	return t.dur
}