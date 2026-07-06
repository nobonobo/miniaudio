//go:build ignore

package main

import (
	"encoding/binary"
	"fmt"
	"math"
	"os"
)

func main() {
	file, err := os.Create("./examples/sample.wav")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	sampleRate := 44100
	duration := 1.0    // seconds
	frequency := 440.0 // A4 note
	numSamples := int(sampleRate) * int(duration)

	// Write WAV header
	writeWavHeader(file, numSamples, uint32(sampleRate))

	// Generate sine wave
	for i := 0; i < numSamples; i++ {
		t := float64(i) / float64(sampleRate)
		sample := math.Sin(2 * math.Pi * frequency * t)
		// Convert to 16-bit PCM
		intSample := int16(sample * 32767)
		binary.Write(file, binary.LittleEndian, intSample)
	}

	fmt.Println("WAV file generated: sample.wav")
}

func writeWavHeader(file *os.File, numSamples int, sampleRate uint32) {
	numChannels := uint16(1)
	bitsPerSample := uint16(16)
	dataSize := numSamples * 2 // 1 channel, 2 bytes per sample
	headerSize := 44
	totalSize := headerSize + dataSize

	// RIFF header
	file.Write([]byte("RIFF"))
	binary.Write(file, binary.LittleEndian, uint32(totalSize-8))
	file.Write([]byte("WAVE"))

	// fmt chunk
	file.Write([]byte("fmt "))
	binary.Write(file, binary.LittleEndian, uint32(16)) // Chunk size
	binary.Write(file, binary.LittleEndian, uint16(1))  // PCM format
	binary.Write(file, binary.LittleEndian, numChannels)
	binary.Write(file, binary.LittleEndian, sampleRate)
	binary.Write(file, binary.LittleEndian, uint32(sampleRate)*uint32(numChannels)*uint32(bitsPerSample)/8)
	binary.Write(file, binary.LittleEndian, uint16(numChannels*bitsPerSample/8))
	binary.Write(file, binary.LittleEndian, bitsPerSample)

	// data chunk
	file.Write([]byte("data"))
	binary.Write(file, binary.LittleEndian, uint32(dataSize))
}
