package drum

import (
	"time"
)

const (
	Ppqn        = 24.0
	Minute      = 60.0
	Microsecond = 1000000000
	DefaultBPM  = 120.0
)

type Timer struct {
	Pulses chan int
	Done   chan bool
	BPM    float32
}

func NewTimer() *Timer {
	t := &Timer{
		Pulses: make(chan int),
		Done:   make(chan bool),
		BPM:    DefaultBPM,
	}

	return t
}

func (t *Timer) SetBPM(bpm float32) {
	t.BPM = bpm
}

func (t *Timer) Start() {
	go func() {
		for {
			select {
			case <-t.Done:
				break
			default:
				interval := microsecondsPerPulse(t.BPM)
				time.Sleep(interval)
				t.Pulses <- 1
			}
		}
	}()
}

func microsecondsPerPulse(bpm float32) time.Duration {
	return time.Duration((Minute * Microsecond) / (Ppqn * bpm))
}
