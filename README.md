# miniaudio

A Go wrapper for [miniaudio](https://github.com/mackron/miniaudio), providing native audio playback and capture capabilities through CGO bindings. This project enables Go applications to leverage the powerful miniaudio library for cross-platform audio processing with support for 3D spatialization, node-based effects graphs, and multiple audio backends.

## Features

- **Cross-platform audio**: Supports Windows (WASAPI, DirectSound, MME), Linux (ALSA, PulseAudio, JACK, OSS), and macOS/iOS (Core Audio)
- **High-level audio engine**: Simple API for common audio operations
- **3D spatial audio**: Positional audio with listener positioning, Doppler effect, and attenuation models
- **Node-based effects graph**: Build complex audio processing pipelines with customizable nodes
- **Sound playback**: Play sounds from files, memory buffers, or custom data sources
- **Decoder support**: Decode various audio formats (WAV, MP3, OGG, FLAC, etc.)
- **Low-level device access**: Direct control of audio devices when needed

## Design Principles

This library follows a minimal wrapper approach:

- All Go functions call miniaudio's C API directly
- CGO bindings are kept minimal and focused
- Structs wrapping C pointers follow the pattern: `NewXxx()` for creation, `Close()` for cleanup
- One-to-one mapping between C API functions and Go methods
- No duplicate functionality beyond what miniaudio provides

## Installation

```bash
go get github.com/nobonobo/miniaudio
```

### Prerequisites

- Go 1.23 or later
- A C compiler (GCC, Clang, or MSVC)
- miniaudio.h and miniaudio.c files in the project directory (included)

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/nobonobo/miniaudio"
)

func main() {
    // Create and initialize the engine
    config := miniaudio.NewEngineConfig()
    defer config.Close()
    
    engine, err := miniaudio.NewEngine(config)
    if err != nil {
        fmt.Printf("Failed to create engine: %v\n", err)
        return
    }
    defer engine.Close()
    
    // Start the engine
    if err := engine.Start(); err != nil {
        fmt.Printf("Failed to start engine: %v\n", err)
        return
    }
    
    // Play a sound file
    if err := engine.PlaySound("path/to/sound.wav"); err != nil {
        fmt.Printf("Failed to play sound: %v\n", err)
    }
}
```

## API Overview

### Engine

The high-level audio engine for most audio needs:

```go
// Create engine with custom configuration
config := miniaudio.NewEngineConfig()
config.SetSampleRate(48000)
config.SetChannels(2)
engine, err := miniaudio.NewEngine(config)

// Engine methods
engine.Start()
engine.Stop()
engine.PlaySound("sound.wav")
engine.SetVolume(0.5)
```

### Sound Playback

```go
// Play from file
sound, err := engine.PlaySoundFromFile("sound.wav", miniaudio.SoundFlagLooping, nil)

// Control sound
sound.SetVolume(0.8)
sound.SetPan(0.5)
sound.SetPitch(1.0)
sound.Play()
sound.Stop()
```

### 3D Audio

```go
// Set listener position and orientation
engine.ListenerSetPosition(0, 0, 0, 0)
engine.ListenerSetDirection(0, 0, 0, -1)
engine.ListenerSetVelocity(0, 0, 0, 0)

// Spatial sounds automatically attenuate based on distance
```

### Device Management

```go
// Access underlying device for advanced control
device := engine.GetDevice()
```

## Project Structure

```
.
├── miniaudio.go          # Main package and type definitions
├── miniaudio_types.go    # C struct wrappers and type constants
├── engine.go             # Engine and EngineConfig implementations
├── device.go             # Device and DeviceConfig implementations
├── decoder.go            # Decoder implementation
├── sound.go              # Sound implementation
├── nodegraph.go          # Node graph implementation
├── context.go            # Context implementation
├── miniaudio.h           # miniaudio header (included)
├── miniaudio.c           # miniaudio source (included)
├── miniaudio_callbacks.c # CGO callback bridge implementations
├── miniaudio_callbacks.h # Callback header definitions
└── docs/
    ├── api.md            # API reference
    └── memo.md           # Development notes
```

## Examples

### Generate and Play a Sine Wave

```go
// See examples/wav_generator.go for WAV file generation
go run examples/wav_generator.go
```

### Basic Sound Playback

```go
// Run the test example
go run miniaudio_test.go
```

## Supported Backends

| Backend | Platforms |
|---------|-----------|
| WASAPI | Windows 10+ |
| DirectSound | Windows |
| MME | Windows |
| Core Audio | macOS, iOS |
| ALSA | Linux |
| PulseAudio | Linux |
| JACK | Linux |
| OSS | BSD, Linux |
| AudioIO | iOS, tvOS |

## Constants and Types

### Sound Flags

```go
const (
    SoundFlagStream        SoundFlags = 0x00000001  // Stream from source
    SoundFlagDecode        SoundFlags = 0x00000002  // Decode on the fly
    SoundFlagAsync         SoundFlags = 0x00000004  // Async loading
    SoundFlagWaitInit      SoundFlags = 0x00000008  // Wait for init
    SoundFlagUnknownLength SoundFlags = 0x00000010  // Unknown length
    SoundFlagLooping       SoundFlags = 0x00000020  // Loop playback
)
```

### Audio Formats

```go
const (
    FormatUnknown Format = 0
    FormatF32     Format = 1
    FormatF64     Format = 2
    FormatS16     Format = 3
    FormatS24     Format = 4
    FormatS32     Format = 5
    FormatU8      Format = 6
)
```

### Device Types

```go
const (
    DeviceTypePlayback DeviceType = 1 << 0
    DeviceTypeCapture  DeviceType = 1 << 1
    DeviceTypeDuplex   DeviceType = 1 << 2
    DeviceTypeLoopback DeviceType = 1 << 3
)
```

## Error Handling

All functions return Go `error` types. Custom error types include:

- `ErrNilEngine`: Returned when the engine is nil
- `ErrResult`: Converted from miniaudio's `ma_result` codes

```go
if err := engine.PlaySound("sound.wav"); err != nil {
    fmt.Printf("Error: %v\n", err)
}
```

## Development

### Building

```bash
# Build the package
go build ./...

# Run tests
go test -v ./...
```

### Architecture

The wrapper follows a strict pattern:

1. C structs are wrapped by Go structs containing only the C pointer
2. `NewXxx()` functions allocate C memory, initialize, and return the wrapper
3. `Close()` methods uninitialize and free C memory
4. Methods on Go structs call corresponding C API functions
5. Callbacks are bridged through dedicated `.c`/`.h` files

## License

This project uses the miniaudio library which is licensed under the MIT License. See the included `LICENSE` file for details.

## Links

- [GitHub Repository](git@github.com:nobonobo/miniaudio.git)
- [miniaudio Documentation](https://github.com/mackron/miniaudio)