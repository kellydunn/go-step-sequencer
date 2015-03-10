package sequencer

import (
	"testing"
)

func TestMicrosecondsPerPulse(t *testing.T) {
	timer := NewTimer()
	expected := 20833334
	actual := timer.MicrosecondsPerPulse()
	if expected != int(actual) {
		t.Errorf("Unexpected Microseconds Per pulse for 120.0 bpm.  Expected: %d, Actual: %d\n", expected, actual)
	}
}
