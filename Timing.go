package main

import "time"

type Timing struct {
	time.Time
}

func (t *Timing) Start() {
	t.Time = time.Now()
}

func (t *Timing) End() time.Duration {
	return time.Since(t.Time)
}
