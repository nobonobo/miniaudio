package main

import (
	"log"
	"time"

	ma "github.com/nobonobo/miniaudio"
)

func main() {
	config := ma.NewEngineConfig()
	// Create a new engine
	engine, err := ma.NewEngine(config)
	if err != nil {
		log.Fatal(err)
	}
	defer engine.Close()
	if err := engine.Start(); err != nil {
		log.Fatal(err)
	}
	snd, err := engine.PlaySoundFromFile("./examples/sample.wav", 0, nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer snd.Close()
	if err := snd.Start(); err != nil {
		log.Fatal(err)
	}
	for snd.IsPlaying() {
		time.Sleep(1 * time.Millisecond)
	}
	// Verify sample rate and channels (should be 0 in no device mode)
	sampleRate := engine.GetSampleRate()
	channels := engine.GetChannels()
	log.Println("Sample rate: ", sampleRate, "Channels: ", channels)
}
