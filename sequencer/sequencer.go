package drum

import (
	portaudio "code.google.com/p/portaudio-go/portaudio"
	"fmt"
	drum "github.com/kellydunn/go-challenge-1"
	"github.com/mkb218/gosndfile/sndfile"
)

const (
	SampleRate     = 44100
	InputChannels  = 0
	OutputChannels = 2
)

type Sequencer struct {
	Timer   *Timer
	Bar     int
	Beat    int
	Pattern *drum.Pattern
	Stream  *portaudio.Stream
}

func NewSequencer() (*Sequencer, error) {
	err := portaudio.Initialize()
	if err != nil {
		return nil, err
	}

	s := &Sequencer{
		Timer: NewTimer(),
		Bar:   0,
		Beat:  0,
	}

	stream, err := portaudio.OpenDefaultStream(
		InputChannels,
		OutputChannels,
		SampleRate,
		portaudio.FramesPerBufferUnspecified,
		s.ProcessAudio,
	)

	if err != nil {
		return nil, err
	}

	s.Stream = stream

	return s, nil
}

func (s *Sequencer) Start() {
	go func() {
		ppqnCount := 0

		for {
			select {
			case <-s.Timer.Pulses:
				ppqnCount += 1

				// TODO add in time signatures
				if ppqnCount%(int(Ppqn)/4) == 0 {
					go s.PlayTrigger()

					s.Beat += 1
					s.Beat = s.Beat % 4
				}

				// TODO Add in time signatures
				if ppqnCount%int(Ppqn) == 0 {
					s.Bar += 1
					s.Bar = s.Bar % 4
				}

				// 4 bars of quarter notes
				if ppqnCount == (int(Ppqn) * 4) {
					ppqnCount = 0
				}

			}
		}
	}()

	s.Timer.Start()
	s.Stream.Start()
}

func (s *Sequencer) ProcessAudio(out []float32) {
	for i := range out {
		var data float32

		for _, track := range s.Pattern.Tracks {
			if track.Playhead < len(track.Buffer) {
				data += track.Buffer[track.Playhead]
				track.Playhead++
			}
		}

		if data > 1.0 {
			data = 1.0
		}

		out[i] = data
	}
}

func (s *Sequencer) PlayTrigger() {
	index := (s.Bar * 4) + s.Beat

	for _, track := range s.Pattern.Tracks {
		if track.StepSequence.Steps[index] == byte(1) {
			track.Playhead = 0
		}
	}
}

func LoadSample(filename string) ([]float32, error) {
	var info sndfile.Info
	soundFile, err := sndfile.Open(filename, sndfile.Read, &info)
	if err != nil {
		fmt.Printf("Could not open file: %s\n", filename)
		return nil, err
	}

	buffer := make([]float32, 10*info.Samplerate*info.Channels)
	numRead, err := soundFile.ReadItems(buffer)
	if err != nil {
		fmt.Printf("Error reading data from file: %s\n", filename)
		return nil, err
	}

	defer soundFile.Close()

	return buffer[:numRead], nil
}
