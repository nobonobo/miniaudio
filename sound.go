package miniaudio

/*
#include "miniaudio.h"
#include <stdlib.h>
#include <string.h>
*/
import "C"
import "unsafe"

// SoundConfig wraps ma_sound_config.
type SoundConfig struct {
	config *C.ma_sound_config
}

// NewSoundConfig creates a new SoundConfig and calls ma_sound_config_init_2.
func NewSoundConfig(engine *Engine) *SoundConfig {
	if engine == nil || engine.engine == nil {
		return nil
	}
	config := C.ma_sound_config_init_2(engine.engine)
	return &SoundConfig{config: &config}
}

func (sc *SoundConfig) Close() {
	if sc.config == nil {
		return
	}
	sc.config = nil
}

// Sound wraps ma_sound.
type Sound struct {
	sound *C.ma_sound
}

// NewSound creates a new Sound instance and initializes it with ma_sound_init_ex.
func NewSound(engine *Engine, config *SoundConfig) (*Sound, error) {
	if engine == nil || engine.engine == nil {
		return nil, ErrNilEngine
	}

	s := (*C.ma_sound)(C.malloc(C.size_t(unsafe.Sizeof(*(*C.ma_sound)(nil)))))
	if s == nil {
		return nil, ErrNilEngine
	}
	C.memset(unsafe.Pointer(s), 0, C.size_t(unsafe.Sizeof(*s)))

	var cConfig *C.ma_sound_config
	if config != nil {
		cConfig = config.config
	}

	result := C.ma_sound_init_ex(engine.engine, cConfig, s)
	if result != 0 {
		C.free(unsafe.Pointer(s))
		return nil, ErrResult(result)
	}

	return &Sound{sound: s}, nil
}

// Close uninitializes and releases the sound memory.
func (s *Sound) Close() {
	if s.sound == nil {
		return
	}
	C.ma_sound_uninit(s.sound)
	C.free(unsafe.Pointer(s.sound))
	s.sound = nil
}

// Start starts the sound.
func (s *Sound) Start() error {
	if s.sound == nil {
		return ErrNilEngine
	}
	result := C.ma_sound_start(s.sound)
	if result != 0 {
		return ErrResult(result)
	}
	return nil
}

// Stop stops the sound.
func (s *Sound) Stop() error {
	if s.sound == nil {
		return ErrNilEngine
	}
	result := C.ma_sound_stop(s.sound)
	if result != 0 {
		return ErrResult(result)
	}
	return nil
}

// SetVolume sets the sound volume.
func (s *Sound) SetVolume(volume float32) {
	if s.sound == nil {
		return
	}
	C.ma_sound_set_volume(s.sound, C.float(volume))
}

// Volume gets the sound volume.
func (s *Sound) Volume() float32 {
	if s.sound == nil {
		return 0
	}
	return float32(C.ma_sound_get_volume(s.sound))
}

// SetPan sets the sound pan.
func (s *Sound) SetPan(pan float32) {
	if s.sound == nil {
		return
	}
	C.ma_sound_set_pan(s.sound, C.float(pan))
}

// Pan gets the sound pan.
func (s *Sound) Pan() float32 {
	if s.sound == nil {
		return 0
	}
	return float32(C.ma_sound_get_pan(s.sound))
}

// SetPitch sets the sound pitch.
func (s *Sound) SetPitch(pitch float32) {
	if s.sound == nil {
		return
	}
	C.ma_sound_set_pitch(s.sound, C.float(pitch))
}

// Pitch gets the sound pitch.
func (s *Sound) Pitch() float32 {
	if s.sound == nil {
		return 0
	}
	return float32(C.ma_sound_get_pitch(s.sound))
}

// SetPosition sets the sound position.
func (s *Sound) SetPosition(x, y, z float32) {
	if s.sound == nil {
		return
	}
	C.ma_sound_set_position(s.sound, C.float(x), C.float(y), C.float(z))
}

// Position gets the sound position.
func (s *Sound) Position() (x, y, z float32) {
	if s.sound == nil {
		return 0, 0, 0
	}
	C.ma_sound_get_position(s.sound)
	return
}

// SetSpatializationEnabled sets the sound spatialization enabled.
func (s *Sound) SetSpatializationEnabled(enabled bool) {
	if s.sound == nil {
		return
	}
	C.ma_sound_set_spatialization_enabled(s.sound, C.ma_bool32(boolToCInt(enabled)))
}

// IsSpatializationEnabled returns whether the sound spatialization is enabled.
func (s *Sound) IsSpatializationEnabled() bool {
	if s.sound == nil {
		return false
	}
	return C.ma_sound_is_spatialization_enabled(s.sound) != 0
}

// IsPlaying returns whether the sound is playing.
func (s *Sound) IsPlaying() bool {
	if s.sound == nil {
		return false
	}
	return C.ma_sound_is_playing(s.sound) != 0
}

// Engine returns the engine this sound belongs to.
func (s *Sound) Engine() *Engine {
	if s.sound == nil {
		return nil
	}
	engine := (*C.ma_engine)(C.ma_sound_get_engine(s.sound))
	return &Engine{engine: engine}
}

// DataSource returns the data source of this sound.
func (s *Sound) DataSource() unsafe.Pointer {
	if s.sound == nil {
		return nil
	}
	return unsafe.Pointer(C.ma_sound_get_data_source(s.sound))
}

// ResetStartTime resets the sound start time.
func (s *Sound) ResetStartTime() {
	if s.sound == nil {
		return
	}
	C.ma_sound_reset_start_time(s.sound)
}

// ResetStopTime resets the sound stop time.
func (s *Sound) ResetStopTime() {
	if s.sound == nil {
		return
	}
	C.ma_sound_reset_stop_time(s.sound)
}
