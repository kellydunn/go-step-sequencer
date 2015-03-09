package sequencer

import (
	"time"
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
		BPM:    float32(DefaultBPM),
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
	return time.Duration((float32(Minute) * float32(Microsecond)) / (float32(Ppqn) * bpm))
}
