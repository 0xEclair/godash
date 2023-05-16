package main

import "time"

type Timing struct {
	time.Time
}

func (t *Timing) Start() {
	t.Time = time.Now()
}

func (t *Timing) Passed() time.Duration {
	dur := time.Since(t.Time)
	t.Time = time.Now()

	return dur
}

func (t *Timing) End() time.Duration {
	return time.Since(t.Time)
}
