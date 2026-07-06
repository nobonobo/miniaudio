package main

import (
	"log"
	"math"
	"time"

	ma "github.com/nobonobo/miniaudio"
)

func main() {
	config := ma.NewEngineConfig()
	engine, err := ma.NewEngine(config)
	if err != nil {
		log.Fatal(err)
	}
	defer engine.Close()

	if err := engine.Start(); err != nil {
		log.Fatal(err)
	}

	// Load sound from file
	snd, err := engine.PlaySoundFromFile("./examples/sample.wav", ma.SoundFlagLooping, nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer snd.Close()

	// Enable spatialization for 3D positioning
	//snd.SetSpatializationEnabled(true)

	// Enable looping
	snd.ResetStartTime()

	if err := snd.Start(); err != nil {
		log.Fatal(err)
	}

	// Orbit parameters
	radius := float32(5.0)
	speed := 2.0 // radians per second

	log.Println("Orbiting sound around listener...")
	log.Println("Press Ctrl+C to stop.")

	startTime := time.Now()

	for {
		elapsed := time.Since(startTime).Seconds()

		// Calculate sound position on orbit circle
		angle := elapsed * speed
		soundX := float32(math.Cos(angle) * float64(radius))
		soundZ := float32(math.Sin(angle) * float64(radius))
		soundY := float32(1.0) // height above ground

		// Update sound position
		snd.SetPosition(soundX, soundY, soundZ)

		// Small sleep to avoid busy waiting
		time.Sleep(50 * time.Millisecond)
	}
}
