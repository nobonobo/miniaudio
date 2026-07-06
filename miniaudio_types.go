package miniaudio

/*
#include "miniaudio.h"
#include "stdlib.h"
#include "string.h"
*/
import "C"
import "unsafe"

// Types

// Format represents ma_format enum values.
type Format int32

const (
	FormatUnknown Format = 0
	FormatF32     Format = 1
	FormatF64     Format = 2
	FormatS16     Format = 3
	FormatS24     Format = 4
	FormatS32     Format = 5
	FormatU8      Format = 6
)

// DeviceType represents ma_device_type enum values.
type DeviceType int32

const (
	DeviceTypePlayback DeviceType = 1 << 0
	DeviceTypeCapture  DeviceType = 1 << 1
	DeviceTypeDuplex   DeviceType = 1 << 2
	DeviceTypeLoopback DeviceType = 1 << 3
)

// DeviceDirection represents ma_device_direction enum values.
type DeviceDirection int32

const (
	DeviceDirectionPlayback DeviceDirection = 1 << 0
	DeviceDirectionCapture  DeviceDirection = 1 << 1
)

// Channel represents ma_channel enum values.
type Channel int32

const (
	ChannelMono               Channel = 1
	ChannelFrontLeft          Channel = 2
	ChannelFrontRight         Channel = 3
	ChannelFrontCenter        Channel = 4
	ChannelBackCenter         Channel = 5
	ChannelBackLeft           Channel = 6
	ChannelBackRight          Channel = 7
	ChannelFrontLeftOfCenter  Channel = 8
	ChannelFrontRightOfCenter Channel = 9
	ChannelBackLeftOfCenter   Channel = 10
	ChannelBackRightOfCenter  Channel = 11
	ChannelSideLeft           Channel = 12
	ChannelSideRight          Channel = 13
	ChannelLFE1               Channel = 14
	ChannelStereoLeft         Channel = 15
	ChannelStereoRight        Channel = 16
)

// ResamplerQuality represents ma_resampler_quality enum values.
type ResamplerQuality int32

const (
	ResamplerQualityLow    ResamplerQuality = 0
	ResamplerQualityMedium ResamplerQuality = 1
	ResamplerQualityHigh   ResamplerQuality = 2
	ResamplerQualityHighFC ResamplerQuality = 3
)

// NodeState represents ma_node_state enum values.
type NodeState int32

const (
	NodeStateInitializing   NodeState = 0
	NodeStateReady          NodeState = 1
	NodeStateProcessing     NodeState = 2
	NodeStateStopping       NodeState = 3
	NodeStateStopped        NodeState = 4
	NodeStateUninitializing NodeState = 5
)

// SeekOrigin represents ma_seek_origin enum values.
type SeekOrigin int32

const (
	SeekOriginBegin   SeekOrigin = 0
	SeekOriginCurrent SeekOrigin = 1
	SeekOriginEnd     SeekOrigin = 2
)

// EncodingFormat represents ma_encoding_format enum values.
type EncodingFormat int32

const (
	EncodingFormatUnknown EncodingFormat = 0
	EncodingFormatWav     EncodingFormat = 1
	EncodingFormatAiff    EncodingFormat = 2
)

// DitherMode represents ma_dither_mode enum values.
type DitherMode int32

const (
	DitherModeNone        DitherMode = 0
	DitherModeRectangular DitherMode = 1
)

// ChannelMixMode represents ma_channel_mix_mode enum values.
type ChannelMixMode int32

const (
	ChannelMixModeDefault     ChannelMixMode = 0
	ChannelMixModeStrict      ChannelMixMode = 1
	ChannelMixModeAllowLosses ChannelMixMode = 2
)

// Backend represents available audio backends.
type Backend int32

const (
	BackendNone        Backend = 0
	BackendWASAPI      Backend = 1 << 0
	BackendDirectSound Backend = 1 << 1
	BackendMME         Backend = 1 << 2
	BackendAudioIO     Backend = 1 << 3
	BackendPortAudio   Backend = 1 << 4
	BackendALSA        Backend = 1 << 5
	BackendPulseAudio  Backend = 1 << 6
	BackendCoreAudio   Backend = 1 << 7
	BackendJACK        Backend = 1 << 8
	BackendOSS         Backend = 1 << 9
	BackendNaCl        Backend = 1 << 10
)

// Helper functions

func boolToCInt(b bool) C.ma_bool32 {
	if b {
		return 1
	}
	return 0
}

const (
	MA_SUCCESS               = 0
	MA_ERROR                 = -1
	MA_INVALID_ARGS          = -2
	MA_NO_MEMORY             = -3
	MA_OUT_OF_RANGE          = -4
	MA_ALREADY_INITIALIZED   = -5
	MA_NOT_FOUND             = -6
	MA_INVALID_OPERATION     = -7
	MA_INVALID_STATE         = -8
	MA_AUDIO_NOT_AVAILABLE   = -9
	MA_BACKEND_NOT_AVAILABLE = -10
	MA_NO_DATA               = -11
	MA_ALREADY_EOS           = -12
	MA_AT_END                = 0x1
	MA_NOT_IMPLEMENTED       = -13
)

// ErrNilEngine is returned when the engine is nil.
var ErrNilEngine = &EngineError{msg: "engine is nil"}

// EngineError is an error type for engine operations.
type EngineError struct {
	msg string
}

func (e *EngineError) Error() string {
	return e.msg
}

// ErrResult converts a ma_result to a Go error.
func ErrResult(result C.ma_result) error {
	switch int(result) {
	case MA_SUCCESS:
		return nil
	case MA_ERROR:
		return &EngineError{msg: "MA_ERROR"}
	case MA_INVALID_ARGS:
		return &EngineError{msg: "MA_INVALID_ARGS"}
	case MA_NO_MEMORY:
		return &EngineError{msg: "MA_NO_MEMORY"}
	case MA_OUT_OF_RANGE:
		return &EngineError{msg: "MA_OUT_OF_RANGE"}
	case MA_ALREADY_INITIALIZED:
		return &EngineError{msg: "MA_ALREADY_INITIALIZED"}
	case MA_NOT_FOUND:
		return &EngineError{msg: "MA_NOT_FOUND"}
	case MA_INVALID_OPERATION:
		return &EngineError{msg: "MA_INVALID_OPERATION"}
	case MA_INVALID_STATE:
		return &EngineError{msg: "MA_INVALID_STATE"}
	case MA_AUDIO_NOT_AVAILABLE:
		return &EngineError{msg: "MA_AUDIO_NOT_AVAILABLE"}
	case MA_BACKEND_NOT_AVAILABLE:
		return &EngineError{msg: "MA_BACKEND_NOT_AVAILABLE"}
	case MA_NO_DATA:
		return &EngineError{msg: "MA_NO_DATA"}
	case MA_ALREADY_EOS:
		return &EngineError{msg: "MA_ALREADY_EOS"}
	case MA_AT_END:
		return &EngineError{msg: "MA_AT_END"}
	case MA_NOT_IMPLEMENTED:
		return &EngineError{msg: "MA_NOT_IMPLEMENTED"}
	default:
		return &EngineError{msg: "unknown error"}
	}
}

// mallocWrapper allocates and zero-initializes a C.ma_engine struct.
func mallocWrapper() *C.ma_engine {
	engine := (*C.ma_engine)(C.malloc(C.size_t(unsafe.Sizeof(*(*C.ma_engine)(nil)))))
	if engine != nil {
		C.memset(unsafe.Pointer(engine), 0, C.size_t(unsafe.Sizeof(*engine)))
	}
	return engine
}

// freeWrapper frees the memory allocated for a C.ma_engine struct.
func freeWrapper(engine *C.ma_engine) {
	if engine != nil {
		C.free(unsafe.Pointer(engine))
	}
}

type DataSource interface {
	dataSourcePtr() unsafe.Pointer
}

type Log struct {
	log *C.ma_log
}

type DeviceID struct {
	id *C.ma_device_id
}

type SoundFlags uint32

const (
	SoundFlagStream        SoundFlags = 0x00000001
	SoundFlagDecode        SoundFlags = 0x00000002
	SoundFlagAsync         SoundFlags = 0x00000004
	SoundFlagWaitInit      SoundFlags = 0x00000008
	SoundFlagUnknownLength SoundFlags = 0x00000010
	SoundFlagLooping       SoundFlags = 0x00000020

	SoundFlagNoDefaultAttachment SoundFlags = 0x00001000
	SoundFlagNoPitch             SoundFlags = 0x00002000
	SoundFlagNoSpatialization    SoundFlags = 0x00004000
)
