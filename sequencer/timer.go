package sequencer

import (
	"time"
)

// Timer is a struct that defines the basic synchronization
// behavior of the step sequencer
type Timer struct {
	Pulses chan int
	Done   chan bool
	Tempo  float32
}

// NewTimer creates and returns a new pointer to a Timer.
func NewTimer() *Timer {
	t := &Timer{
		Pulses: make(chan int),
		Done:   make(chan bool),
		Tempo:  float32(DefaultTempo),
	}

	return t
}

// Sets the current Timer's tempo
func (t *Timer) SetTempo(tempo float32) {
	t.Tempo = tempo
}

func (t *Timer) Start() {
	go func() {
		for {
			select {
			case <-t.Done:
				break
			default:
				interval := microsecondsPerPulse(t.Tempo)
				time.Sleep(interval)
				t.Pulses <- 1
			}
		}
	}()
}

func microsecondsPerPulse(bpm float32) time.Duration {
	return time.Duration((float32(Minute) * float32(Microsecond)) / (float32(Ppqn) * bpm))
}
