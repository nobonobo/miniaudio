package miniaudio

/*
#include "miniaudio.h"
#include <stdlib.h>
#include <string.h>
*/
import "C"
import "unsafe"

// DecoderConfig wraps ma_decoder_config.
type DecoderConfig struct {
	config *C.ma_decoder_config
}

// NewDecoderConfig creates a new DecoderConfig and calls ma_decoder_config_init.
func NewDecoderConfig(format Format, channels uint32, sampleRate uint32) *DecoderConfig {
	config := C.ma_decoder_config_init(C.ma_format(format), C.ma_uint32(channels), C.ma_uint32(sampleRate))
	return &DecoderConfig{config: &config}
}

func (dc *DecoderConfig) Close() {
	if dc.config == nil {
		return
	}
	dc.config = nil
}

// Format returns the output format.
func (dc *DecoderConfig) Format() Format {
	if dc == nil || dc.config == nil {
		return 0
	}
	return Format(dc.config.format)
}

// SetFormat sets the output format.
func (dc *DecoderConfig) SetFormat(format Format) {
	if dc != nil && dc.config != nil {
		dc.config.format = C.ma_format(format)
	}
}

// Channels returns the output channel count.
func (dc *DecoderConfig) Channels() uint32 {
	if dc == nil || dc.config == nil {
		return 0
	}
	return uint32(dc.config.channels)
}

// SetChannels sets the output channel count.
func (dc *DecoderConfig) SetChannels(channels uint32) {
	if dc != nil && dc.config != nil {
		dc.config.channels = C.ma_uint32(channels)
	}
}

// SampleRate returns the output sample rate.
func (dc *DecoderConfig) SampleRate() uint32 {
	if dc == nil || dc.config == nil {
		return 0
	}
	return uint32(dc.config.sampleRate)
}

// SetSampleRate sets the output sample rate.
func (dc *DecoderConfig) SetSampleRate(sampleRate uint32) {
	if dc != nil && dc.config != nil {
		dc.config.sampleRate = C.ma_uint32(sampleRate)
	}
}

// PChannelMap returns the channel map pointer.
func (dc *DecoderConfig) PChannelMap() unsafe.Pointer {
	if dc == nil || dc.config == nil {
		return nil
	}
	return unsafe.Pointer(dc.config.pChannelMap)
}

// SetPChannelMap sets the channel map pointer.
func (dc *DecoderConfig) SetPChannelMap(channelMap unsafe.Pointer) {
	if dc != nil && dc.config != nil {
		dc.config.pChannelMap = (*C.ma_channel)(channelMap)
	}
}

// ChannelMixMode returns the channel mix mode.
func (dc *DecoderConfig) ChannelMixMode() ChannelMixMode {
	if dc == nil || dc.config == nil {
		return 0
	}
	return ChannelMixMode(dc.config.channelMixMode)
}

// SetChannelMixMode sets the channel mix mode.
func (dc *DecoderConfig) SetChannelMixMode(mode ChannelMixMode) {
	if dc != nil && dc.config != nil {
		dc.config.channelMixMode = C.ma_channel_mix_mode(mode)
	}
}

// DitherMode returns the dither mode.
func (dc *DecoderConfig) DitherMode() DitherMode {
	if dc == nil || dc.config == nil {
		return 0
	}
	return DitherMode(dc.config.ditherMode)
}

// SetDitherMode sets the dither mode.
func (dc *DecoderConfig) SetDitherMode(mode DitherMode) {
	if dc != nil && dc.config != nil {
		dc.config.ditherMode = C.ma_dither_mode(mode)
	}
}

// ResamplingFormat returns the resampling format.
func (dc *DecoderConfig) ResamplingFormat() Format {
	if dc == nil || dc.config == nil {
		return 0
	}
	return Format(dc.config.resampling.format)
}

// SetResamplingFormat sets the resampling format.
func (dc *DecoderConfig) SetResamplingFormat(format Format) {
	if dc != nil && dc.config != nil {
		dc.config.resampling.format = C.ma_format(format)
	}
}

// ResamplingChannels returns the resampling channel count.
func (dc *DecoderConfig) ResamplingChannels() uint32 {
	if dc == nil || dc.config == nil {
		return 0
	}
	return uint32(dc.config.resampling.channels)
}

// SetResamplingChannels sets the resampling channel count.
func (dc *DecoderConfig) SetResamplingChannels(channels uint32) {
	if dc != nil && dc.config != nil {
		dc.config.resampling.channels = C.ma_uint32(channels)
	}
}

// ResamplingSampleRateIn returns the resampling input sample rate.
func (dc *DecoderConfig) ResamplingSampleRateIn() uint32 {
	if dc == nil || dc.config == nil {
		return 0
	}
	return uint32(dc.config.resampling.sampleRateIn)
}

// SetResamplingSampleRateIn sets the resampling input sample rate.
func (dc *DecoderConfig) SetResamplingSampleRateIn(sampleRateIn uint32) {
	if dc != nil && dc.config != nil {
		dc.config.resampling.sampleRateIn = C.ma_uint32(sampleRateIn)
	}
}

// ResamplingSampleRateOut returns the resampling output sample rate.
func (dc *DecoderConfig) ResamplingSampleRateOut() uint32 {
	if dc == nil || dc.config == nil {
		return 0
	}
	return uint32(dc.config.resampling.sampleRateOut)
}

// SetResamplingSampleRateOut sets the resampling output sample rate.
func (dc *DecoderConfig) SetResamplingSampleRateOut(sampleRateOut uint32) {
	if dc != nil && dc.config != nil {
		dc.config.resampling.sampleRateOut = C.ma_uint32(sampleRateOut)
	}
}

// ResamplingLPFOrder returns the resampling LPF order.
func (dc *DecoderConfig) ResamplingLPFOrder() uint32 {
	if dc == nil || dc.config == nil {
		return 0
	}
	return uint32(dc.config.resampling.linear.lpfOrder)
}

// SetResamplingLPFOrder sets the resampling LPF order.
func (dc *DecoderConfig) SetResamplingLPFOrder(order uint32) {
	if dc != nil && dc.config != nil {
		dc.config.resampling.linear.lpfOrder = C.ma_uint32(order)
	}
}

// EncodingFormat returns the encoding format.
func (dc *DecoderConfig) EncodingFormat() EncodingFormat {
	if dc == nil || dc.config == nil {
		return 0
	}
	return EncodingFormat(dc.config.encodingFormat)
}

// SetEncodingFormat sets the encoding format.
func (dc *DecoderConfig) SetEncodingFormat(format EncodingFormat) {
	if dc != nil && dc.config != nil {
		dc.config.encodingFormat = C.ma_encoding_format(format)
	}
}

// SeekPointCount returns the seek point count.
func (dc *DecoderConfig) SeekPointCount() uint32 {
	if dc == nil || dc.config == nil {
		return 0
	}
	return uint32(dc.config.seekPointCount)
}

// SetSeekPointCount sets the seek point count.
func (dc *DecoderConfig) SetSeekPointCount(count uint32) {
	if dc != nil && dc.config != nil {
		dc.config.seekPointCount = C.ma_uint32(count)
	}
}

// PPCustomBackendVTables returns the custom backend vtable pointers.
func (dc *DecoderConfig) PPCustomBackendVTables() unsafe.Pointer {
	if dc == nil || dc.config == nil {
		return nil
	}
	return unsafe.Pointer(dc.config.ppCustomBackendVTables)
}

// SetPPCustomBackendVTables sets the custom backend vtable pointers.
func (dc *DecoderConfig) SetPPCustomBackendVTables(vtables unsafe.Pointer) {
	if dc != nil && dc.config != nil {
		dc.config.ppCustomBackendVTables = (**C.ma_decoding_backend_vtable)(vtables)
	}
}

// CustomBackendCount returns the custom backend count.
func (dc *DecoderConfig) CustomBackendCount() uint32 {
	if dc == nil || dc.config == nil {
		return 0
	}
	return uint32(dc.config.customBackendCount)
}

// SetCustomBackendCount sets the custom backend count.
func (dc *DecoderConfig) SetCustomBackendCount(count uint32) {
	if dc != nil && dc.config != nil {
		dc.config.customBackendCount = C.ma_uint32(count)
	}
}

// PCustomBackendUserData returns the custom backend user data pointer.
func (dc *DecoderConfig) PCustomBackendUserData() unsafe.Pointer {
	if dc == nil || dc.config == nil {
		return nil
	}
	return unsafe.Pointer(dc.config.pCustomBackendUserData)
}

// SetPCustomBackendUserData sets the custom backend user data pointer.
func (dc *DecoderConfig) SetPCustomBackendUserData(userData unsafe.Pointer) {
	if dc != nil && dc.config != nil {
		dc.config.pCustomBackendUserData = userData
	}
}

// Decoder wraps ma_decoder.
type Decoder struct {
	decoder *C.ma_decoder
}

// NewDecoderFromFile creates a new Decoder from a file path and initializes it with ma_decoder_init_file.
func NewDecoderFromFile(filePath string, config *DecoderConfig) (*Decoder, error) {
	d := (*C.ma_decoder)(C.malloc(C.size_t(unsafe.Sizeof(*(*C.ma_decoder)(nil)))))
	if d == nil {
		return nil, ErrNilEngine
	}
	C.memset(unsafe.Pointer(d), 0, C.size_t(unsafe.Sizeof(*d)))

	var cFilePath *C.char
	cFilePathStr := C.CString(filePath)
	defer C.free(unsafe.Pointer(cFilePathStr))
	cFilePath = cFilePathStr

	var cConfig *C.ma_decoder_config
	if config != nil {
		cConfig = config.config
	}

	result := C.ma_decoder_init_file(cFilePath, cConfig, d)
	if result != 0 {
		C.free(unsafe.Pointer(d))
		return nil, ErrResult(result)
	}

	return &Decoder{decoder: d}, nil
}

// NewDecoderFromMemory creates a new Decoder from a byte slice and initializes it with ma_decoder_init_memory.
func NewDecoderFromMemory(data []byte, config *DecoderConfig) (*Decoder, error) {
	d := (*C.ma_decoder)(C.malloc(C.size_t(unsafe.Sizeof(*(*C.ma_decoder)(nil)))))
	if d == nil {
		return nil, ErrNilEngine
	}
	C.memset(unsafe.Pointer(d), 0, C.size_t(unsafe.Sizeof(*d)))

	var cConfig *C.ma_decoder_config
	if config != nil {
		cConfig = config.config
	}

	result := C.ma_decoder_init_memory(unsafe.Pointer(&data[0]), C.size_t(len(data)), cConfig, d)
	if result != 0 {
		C.free(unsafe.Pointer(d))
		return nil, ErrResult(result)
	}

	return &Decoder{decoder: d}, nil
}

// Close uninitializes and releases the decoder memory.
func (d *Decoder) Close() {
	if d.decoder == nil {
		return
	}
	C.ma_decoder_uninit(d.decoder)
	C.free(unsafe.Pointer(d.decoder))
	d.decoder = nil
}

// ReadPCMFrames reads PCM frames from the decoder.
func (d *Decoder) ReadPCMFrames(framesOut []float32) (uint64, error) {
	if d.decoder == nil {
		return 0, ErrNilEngine
	}
	var framesRead C.ma_uint64
	result := C.ma_decoder_read_pcm_frames(d.decoder, unsafe.Pointer(&framesOut[0]), C.size_t(len(framesOut)), &framesRead)
	if result != 0 {
		return uint64(framesRead), ErrResult(result)
	}
	return uint64(framesRead), nil
}

// LengthInPCMFrames returns the length in PCM frames of the decoded data.
func (d *Decoder) LengthInPCMFrames() (uint64, error) {
	if d.decoder == nil {
		return 0, ErrNilEngine
	}
	var length C.ma_uint64
	result := C.ma_decoder_get_length_in_pcm_frames(d.decoder, &length)
	if result != 0 {
		return 0, ErrResult(result)
	}
	return uint64(length), nil
}

// SeekToPCMFrame seeks to a specific PCM frame position.
func (d *Decoder) SeekToPCMFrame(frameNumber uint64) error {
	if d.decoder == nil {
		return ErrNilEngine
	}
	result := C.ma_decoder_seek_to_pcm_frame(d.decoder, C.ma_uint64(frameNumber))
	if result != 0 {
		return ErrResult(result)
	}
	return nil
}

// PositionInPCMFrames returns the current position in PCM frames.
func (d *Decoder) PositionInPCMFrames() (uint64, error) {
	if d.decoder == nil {
		return 0, ErrNilEngine
	}
	var position C.ma_uint64
	result := C.ma_decoder_get_cursor_in_pcm_frames(d.decoder, &position)
	if result != 0 {
		return 0, ErrResult(result)
	}
	return uint64(position), nil
}
