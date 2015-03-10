package sequencer

import (
	drum "github.com/kellydunn/go-challenge-1"
	"testing"
)

func TestLoadSample(t *testing.T) {
	buf, err := LoadSample("../kits/0.808-alpha/kick.wav")
	if err != nil {
		t.Errorf("Error loading sample %v", err)
	}

	expected := 28350
	actual := len(buf)
	if expected != actual {
		t.Errorf("Unexpected buffer length after reading sample.  Expected: %d, Actual: %d", expected, actual)
	}
}

func TestPlayTrigger(t *testing.T) {
	s, err := NewSequencer()
	if err != nil {
		t.Errorf("Error creating a sequencer %v", err)
	}

	p, err := drum.DecodeFile("../patterns/pattern_1.splice")
	if err != nil {
		t.Errorf("Error decoding pattern %v", err)
	}

	s.Pattern = p
	s.PlayTrigger(0)

	for i, track := range s.Pattern.Tracks {
		if track.Playhead != 0 {
			t.Errorf("Track #%d has incorrect Playhead value.", i)
		}
	}
}
