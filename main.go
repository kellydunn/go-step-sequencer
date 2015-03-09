package main

import (
	"flag"
	"fmt"
	drum "github.com/kellydunn/go-challenge-1"
	sequencer "github.com/kellydunn/go-step-sequencer/sequencer"
	"time"
)

func main() {
	var patternPath string
	var kitPath string

	flag.StringVar(&patternPath, "pattern", "patterns/pattern_1.splice", "-pattern=path/to/pattern.splice")
	flag.StringVar(&kitPath, "kit", "kits", "-kit=path/to/kits")
	flag.Parse()

	pattern, err := drum.DecodeFile(patternPath)
	if err != nil {
		panic(err)
	}

	for _, track := range pattern.Tracks {
		filepath := kitPath + "/" + pattern.Version + "/" + track.Name + ".wav"

		track.Buffer, err = sequencer.LoadSample(filepath)
		if err != nil {
			fmt.Printf("Error obtaining sample: %v\n", err)
			panic(err)
		}

		track.Playhead = len(track.Buffer)

		fmt.Printf("loaded sample: %s\n", filepath)
	}

	fmt.Printf("%s\n", pattern)

	s, err := sequencer.NewSequencer()
	if err != nil {
		panic(err)
	}

	s.Pattern = pattern

	s.Start()

	for {
		time.Sleep(time.Second)
	}
}
