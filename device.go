package miniaudio

/*
#include "miniaudio.h"
#include <stdlib.h>
#include <string.h>
*/
import "C"
import "unsafe"

// DeviceConfig wraps ma_device_config.
type DeviceConfig struct {
	config *C.ma_device_config
}

// NewDeviceConfig creates a new DeviceConfig and calls ma_device_config_init with the given device type.
func NewDeviceConfig(deviceType DeviceType) *DeviceConfig {
	config := C.ma_device_config_init(C.ma_device_type(deviceType))
	return &DeviceConfig{config: &config}
}

func (dc *DeviceConfig) Close() {
	if dc.config == nil {
		return
	}
	dc.config = nil
}

// SampleRate returns the sample rate.
func (dc *DeviceConfig) SampleRate() uint32 {
	if dc == nil || dc.config == nil {
		return 0
	}
	return uint32(dc.config.sampleRate)
}

// SetSampleRate sets the sample rate.
func (dc *DeviceConfig) SetSampleRate(sampleRate uint32) {
	if dc != nil && dc.config != nil {
		dc.config.sampleRate = C.ma_uint32(sampleRate)
	}
}

// PeriodSizeInFrames returns the period size in frames.
func (dc *DeviceConfig) PeriodSizeInFrames() uint32 {
	if dc == nil || dc.config == nil {
		return 0
	}
	return uint32(dc.config.periodSizeInFrames)
}

// SetPeriodSizeInFrames sets the period size in frames.
func (dc *DeviceConfig) SetPeriodSizeInFrames(periodSize uint32) {
	if dc != nil && dc.config != nil {
		dc.config.periodSizeInFrames = C.ma_uint32(periodSize)
	}
}

// PeriodSizeInMilliseconds returns the period size in milliseconds.
func (dc *DeviceConfig) PeriodSizeInMilliseconds() uint32 {
	if dc == nil || dc.config == nil {
		return 0
	}
	return uint32(dc.config.periodSizeInMilliseconds)
}

// SetPeriodSizeInMilliseconds sets the period size in milliseconds.
func (dc *DeviceConfig) SetPeriodSizeInMilliseconds(periodSize uint32) {
	if dc != nil && dc.config != nil {
		dc.config.periodSizeInMilliseconds = C.ma_uint32(periodSize)
	}
}

// Periods returns the periods count.
func (dc *DeviceConfig) Periods() uint32 {
	if dc == nil || dc.config == nil {
		return 0
	}
	return uint32(dc.config.periods)
}

// SetPeriods sets the periods count.
func (dc *DeviceConfig) SetPeriods(periods uint32) {
	if dc != nil && dc.config != nil {
		dc.config.periods = C.ma_uint32(periods)
	}
}

// PerformanceProfile returns the performance profile.
func (dc *DeviceConfig) PerformanceProfile() C.ma_performance_profile {
	if dc == nil || dc.config == nil {
		return 0
	}
	return dc.config.performanceProfile
}

// SetPerformanceProfile sets the performance profile.
func (dc *DeviceConfig) SetPerformanceProfile(profile C.ma_performance_profile) {
	if dc != nil && dc.config != nil {
		dc.config.performanceProfile = profile
	}
}

// NoPreSilencedOutputBuffer returns whether pre-silencing of output buffer is disabled.
func (dc *DeviceConfig) NoPreSilencedOutputBuffer() bool {
	if dc == nil || dc.config == nil {
		return false
	}
	return dc.config.noPreSilencedOutputBuffer != 0
}

// SetNoPreSilencedOutputBuffer sets whether pre-silencing of output buffer is disabled.
func (dc *DeviceConfig) SetNoPreSilencedOutputBuffer(disabled bool) {
	if dc != nil && dc.config != nil {
		dc.config.noPreSilencedOutputBuffer = C.ma_bool8(boolToCInt(disabled))
	}
}

// NoClip returns whether clipping is disabled.
func (dc *DeviceConfig) NoClip() bool {
	if dc == nil || dc.config == nil {
		return false
	}
	return dc.config.noClip != 0
}

// SetNoClip sets whether clipping is disabled.
func (dc *DeviceConfig) SetNoClip(disabled bool) {
	if dc != nil && dc.config != nil {
		dc.config.noClip = C.ma_bool8(boolToCInt(disabled))
	}
}

// NoDisableDenormals returns whether denormal disabling is disabled.
func (dc *DeviceConfig) NoDisableDenormals() bool {
	if dc == nil || dc.config == nil {
		return false
	}
	return dc.config.noDisableDenormals != 0
}

// SetNoDisableDenormals sets whether denormal disabling is disabled.
func (dc *DeviceConfig) SetNoDisableDenormals(disabled bool) {
	if dc != nil && dc.config != nil {
		dc.config.noDisableDenormals = C.ma_bool8(boolToCInt(disabled))
	}
}

// NoFixedSizedCallback returns whether fixed-sized callback disabling is enabled.
func (dc *DeviceConfig) NoFixedSizedCallback() bool {
	if dc == nil || dc.config == nil {
		return false
	}
	return dc.config.noFixedSizedCallback != 0
}

// SetNoFixedSizedCallback sets whether fixed-sized callback disabling is enabled.
func (dc *DeviceConfig) SetNoFixedSizedCallback(enabled bool) {
	if dc != nil && dc.config != nil {
		dc.config.noFixedSizedCallback = C.ma_bool8(boolToCInt(enabled))
	}
}

// DataCallback returns the data callback.
func (dc *DeviceConfig) DataCallback() C.ma_device_data_proc {
	if dc == nil || dc.config == nil {
		return nil
	}
	return dc.config.dataCallback
}

// SetDataCallback sets the data callback.
func (dc *DeviceConfig) SetDataCallback(callback C.ma_device_data_proc) {
	if dc != nil && dc.config != nil {
		dc.config.dataCallback = callback
	}
}

// NotificationCallback returns the notification callback.
func (dc *DeviceConfig) NotificationCallback() C.ma_device_notification_proc {
	if dc == nil || dc.config == nil {
		return nil
	}
	return dc.config.notificationCallback
}

// SetNotificationCallback sets the notification callback.
func (dc *DeviceConfig) SetNotificationCallback(callback C.ma_device_notification_proc) {
	if dc != nil && dc.config != nil {
		dc.config.notificationCallback = callback
	}
}

// StopCallback returns the stop callback.
func (dc *DeviceConfig) StopCallback() C.ma_stop_proc {
	if dc == nil || dc.config == nil {
		return nil
	}
	return dc.config.stopCallback
}

// SetStopCallback sets the stop callback.
func (dc *DeviceConfig) SetStopCallback(callback C.ma_stop_proc) {
	if dc != nil && dc.config != nil {
		dc.config.stopCallback = callback
	}
}

// PUserData returns the user data pointer.
func (dc *DeviceConfig) PUserData() unsafe.Pointer {
	if dc == nil || dc.config == nil {
		return nil
	}
	return unsafe.Pointer(dc.config.pUserData)
}

// SetPUserData sets the user data pointer.
func (dc *DeviceConfig) SetPUserData(userData unsafe.Pointer) {
	if dc != nil && dc.config != nil {
		dc.config.pUserData = userData
	}
}

// Playback getters

// PlaybackDeviceID returns the playback device ID.
func (dc *DeviceConfig) PlaybackDeviceID() *DeviceID {
	if dc == nil || dc.config == nil {
		return nil
	}
	return &DeviceID{id: dc.config.playback.pDeviceID}
}

// SetPlaybackDeviceID sets the playback device ID.
func (dc *DeviceConfig) SetPlaybackDeviceID(id *DeviceID) {
	if dc != nil && dc.config != nil && id != nil {
		dc.config.playback.pDeviceID = id.id
	}
}

// PlaybackFormat returns the playback format.
func (dc *DeviceConfig) PlaybackFormat() Format {
	if dc == nil || dc.config == nil {
		return 0
	}
	return Format(dc.config.playback.format)
}

// SetPlaybackFormat sets the playback format.
func (dc *DeviceConfig) SetPlaybackFormat(format Format) {
	if dc != nil && dc.config != nil {
		dc.config.playback.format = C.ma_format(format)
	}
}

// PlaybackChannels returns the playback channel count.
func (dc *DeviceConfig) PlaybackChannels() uint32 {
	if dc == nil || dc.config == nil {
		return 0
	}
	return uint32(dc.config.playback.channels)
}

// SetPlaybackChannels sets the playback channel count.
func (dc *DeviceConfig) SetPlaybackChannels(channels uint32) {
	if dc != nil && dc.config != nil {
		dc.config.playback.channels = C.ma_uint32(channels)
	}
}

// PlaybackPChannelMap returns the playback channel map pointer.
func (dc *DeviceConfig) PlaybackPChannelMap() unsafe.Pointer {
	if dc == nil || dc.config == nil {
		return nil
	}
	return unsafe.Pointer(dc.config.playback.pChannelMap)
}

// SetPlaybackPChannelMap sets the playback channel map pointer.
func (dc *DeviceConfig) SetPlaybackPChannelMap(channelMap unsafe.Pointer) {
	if dc != nil && dc.config != nil {
		dc.config.playback.pChannelMap = (*C.ma_channel)(channelMap)
	}
}

// PlaybackChannelMixMode returns the playback channel mix mode.
func (dc *DeviceConfig) PlaybackChannelMixMode() ChannelMixMode {
	if dc == nil || dc.config == nil {
		return 0
	}
	return ChannelMixMode(dc.config.playback.channelMixMode)
}

// SetPlaybackChannelMixMode sets the playback channel mix mode.
func (dc *DeviceConfig) SetPlaybackChannelMixMode(mode ChannelMixMode) {
	if dc != nil && dc.config != nil {
		dc.config.playback.channelMixMode = C.ma_channel_mix_mode(mode)
	}
}

// PlaybackCalculateLFEFromSpatialChannels returns whether LFE calculation from spatial channels is enabled.
func (dc *DeviceConfig) PlaybackCalculateLFEFromSpatialChannels() bool {
	if dc == nil || dc.config == nil {
		return false
	}
	return dc.config.playback.calculateLFEFromSpatialChannels != 0
}

// SetPlaybackCalculateLFEFromSpatialChannels sets whether LFE calculation from spatial channels is enabled.
func (dc *DeviceConfig) SetPlaybackCalculateLFEFromSpatialChannels(enabled bool) {
	if dc != nil && dc.config != nil {
		dc.config.playback.calculateLFEFromSpatialChannels = C.ma_bool32(boolToCInt(enabled))
	}
}

// PlaybackShareMode returns the playback share mode.
func (dc *DeviceConfig) PlaybackShareMode() C.ma_share_mode {
	if dc == nil || dc.config == nil {
		return 0
	}
	return dc.config.playback.shareMode
}

// SetPlaybackShareMode sets the playback share mode.
func (dc *DeviceConfig) SetPlaybackShareMode(mode C.ma_share_mode) {
	if dc != nil && dc.config != nil {
		dc.config.playback.shareMode = mode
	}
}

// Capture getters

// CaptureDeviceID returns the capture device ID.
func (dc *DeviceConfig) CaptureDeviceID() *DeviceID {
	if dc == nil || dc.config == nil {
		return nil
	}
	return &DeviceID{id: dc.config.capture.pDeviceID}
}

// SetCaptureDeviceID sets the capture device ID.
func (dc *DeviceConfig) SetCaptureDeviceID(id *DeviceID) {
	if dc != nil && dc.config != nil && id != nil {
		dc.config.capture.pDeviceID = id.id
	}
}

// CaptureFormat returns the capture format.
func (dc *DeviceConfig) CaptureFormat() Format {
	if dc == nil || dc.config == nil {
		return 0
	}
	return Format(dc.config.capture.format)
}

// SetCaptureFormat sets the capture format.
func (dc *DeviceConfig) SetCaptureFormat(format Format) {
	if dc != nil && dc.config != nil {
		dc.config.capture.format = C.ma_format(format)
	}
}

// CaptureChannels returns the capture channel count.
func (dc *DeviceConfig) CaptureChannels() uint32 {
	if dc == nil || dc.config == nil {
		return 0
	}
	return uint32(dc.config.capture.channels)
}

// SetCaptureChannels sets the capture channel count.
func (dc *DeviceConfig) SetCaptureChannels(channels uint32) {
	if dc != nil && dc.config != nil {
		dc.config.capture.channels = C.ma_uint32(channels)
	}
}

// CapturePChannelMap returns the capture channel map pointer.
func (dc *DeviceConfig) CapturePChannelMap() unsafe.Pointer {
	if dc == nil || dc.config == nil {
		return nil
	}
	return unsafe.Pointer(dc.config.capture.pChannelMap)
}

// SetCapturePChannelMap sets the capture channel map pointer.
func (dc *DeviceConfig) SetCapturePChannelMap(channelMap unsafe.Pointer) {
	if dc != nil && dc.config != nil {
		dc.config.capture.pChannelMap = (*C.ma_channel)(channelMap)
	}
}

// CaptureChannelMixMode returns the capture channel mix mode.
func (dc *DeviceConfig) CaptureChannelMixMode() ChannelMixMode {
	if dc == nil || dc.config == nil {
		return 0
	}
	return ChannelMixMode(dc.config.capture.channelMixMode)
}

// SetCaptureChannelMixMode sets the capture channel mix mode.
func (dc *DeviceConfig) SetCaptureChannelMixMode(mode ChannelMixMode) {
	if dc != nil && dc.config != nil {
		dc.config.capture.channelMixMode = C.ma_channel_mix_mode(mode)
	}
}

// CaptureCalculateLFEFromSpatialChannels returns whether LFE calculation from spatial channels is enabled.
func (dc *DeviceConfig) CaptureCalculateLFEFromSpatialChannels() bool {
	if dc == nil || dc.config == nil {
		return false
	}
	return dc.config.capture.calculateLFEFromSpatialChannels != 0
}

// SetCaptureCalculateLFEFromSpatialChannels sets whether LFE calculation from spatial channels is enabled.
func (dc *DeviceConfig) SetCaptureCalculateLFEFromSpatialChannels(enabled bool) {
	if dc != nil && dc.config != nil {
		dc.config.capture.calculateLFEFromSpatialChannels = C.ma_bool32(boolToCInt(enabled))
	}
}

// CaptureShareMode returns the capture share mode.
func (dc *DeviceConfig) CaptureShareMode() C.ma_share_mode {
	if dc == nil || dc.config == nil {
		return 0
	}
	return dc.config.capture.shareMode
}

// SetCaptureShareMode sets the capture share mode.
func (dc *DeviceConfig) SetCaptureShareMode(mode C.ma_share_mode) {
	if dc != nil && dc.config != nil {
		dc.config.capture.shareMode = mode
	}
}

// WASAPI specific getters

// WasapiNoAutoConvertSRC returns whether auto-convert SRC is disabled for WASAPI.
func (dc *DeviceConfig) WasapiNoAutoConvertSRC() bool {
	if dc == nil || dc.config == nil {
		return false
	}
	return dc.config.wasapi.noAutoConvertSRC != 0
}

// SetWasapiNoAutoConvertSRC sets whether auto-convert SRC is disabled for WASAPI.
func (dc *DeviceConfig) SetWasapiNoAutoConvertSRC(disabled bool) {
	if dc != nil && dc.config != nil {
		dc.config.wasapi.noAutoConvertSRC = C.ma_bool8(boolToCInt(disabled))
	}
}

// WasapiNoDefaultQualitySRC returns whether default quality SRC is disabled for WASAPI.
func (dc *DeviceConfig) WasapiNoDefaultQualitySRC() bool {
	if dc == nil || dc.config == nil {
		return false
	}
	return dc.config.wasapi.noDefaultQualitySRC != 0
}

// SetWasapiNoDefaultQualitySRC sets whether default quality SRC is disabled for WASAPI.
func (dc *DeviceConfig) SetWasapiNoDefaultQualitySRC(disabled bool) {
	if dc != nil && dc.config != nil {
		dc.config.wasapi.noDefaultQualitySRC = C.ma_bool8(boolToCInt(disabled))
	}
}

// WasapiLoopbackProcessID returns the WASAPI loopback process ID.
func (dc *DeviceConfig) WasapiLoopbackProcessID() uint32 {
	if dc == nil || dc.config == nil {
		return 0
	}
	return uint32(dc.config.wasapi.loopbackProcessID)
}

// SetWasapiLoopbackProcessID sets the WASAPI loopback process ID.
func (dc *DeviceConfig) SetWasapiLoopbackProcessID(id uint32) {
	if dc != nil && dc.config != nil {
		dc.config.wasapi.loopbackProcessID = C.ma_uint32(id)
	}
}

// WasapiNoHardwareOffloading returns whether hardware offloading is disabled for WASAPI.
func (dc *DeviceConfig) WasapiNoHardwareOffloading() bool {
	if dc == nil || dc.config == nil {
		return false
	}
	return dc.config.wasapi.noHardwareOffloading != 0
}

// SetWasapiNoHardwareOffloading sets whether hardware offloading is disabled for WASAPI.
func (dc *DeviceConfig) SetWasapiNoHardwareOffloading(disabled bool) {
	if dc != nil && dc.config != nil {
		dc.config.wasapi.noHardwareOffloading = C.ma_bool8(boolToCInt(disabled))
	}
}

// ALSA specific getters

// AlsaNoMMap returns whether MMap is disabled for ALSA.
func (dc *DeviceConfig) AlsaNoMMap() bool {
	if dc == nil || dc.config == nil {
		return false
	}
	return dc.config.alsa.noMMap != 0
}

// SetAlsaNoMMap sets whether MMap is disabled for ALSA.
func (dc *DeviceConfig) SetAlsaNoMMap(disabled bool) {
	if dc != nil && dc.config != nil {
		dc.config.alsa.noMMap = C.ma_bool32(boolToCInt(disabled))
	}
}

// AlsaNoAutoFormat returns whether auto format is disabled for ALSA.
func (dc *DeviceConfig) AlsaNoAutoFormat() bool {
	if dc == nil || dc.config == nil {
		return false
	}
	return dc.config.alsa.noAutoFormat != 0
}

// SetAlsaNoAutoFormat sets whether auto format is disabled for ALSA.
func (dc *DeviceConfig) SetAlsaNoAutoFormat(disabled bool) {
	if dc != nil && dc.config != nil {
		dc.config.alsa.noAutoFormat = C.ma_bool32(boolToCInt(disabled))
	}
}

// AlsaNoAutoChannels returns whether auto channels is disabled for ALSA.
func (dc *DeviceConfig) AlsaNoAutoChannels() bool {
	if dc == nil || dc.config == nil {
		return false
	}
	return dc.config.alsa.noAutoChannels != 0
}

// SetAlsaNoAutoChannels sets whether auto channels is disabled for ALSA.
func (dc *DeviceConfig) SetAlsaNoAutoChannels(disabled bool) {
	if dc != nil && dc.config != nil {
		dc.config.alsa.noAutoChannels = C.ma_bool32(boolToCInt(disabled))
	}
}

// AlsaNoAutoResample returns whether auto resample is disabled for ALSA.
func (dc *DeviceConfig) AlsaNoAutoResample() bool {
	if dc == nil || dc.config == nil {
		return false
	}
	return dc.config.alsa.noAutoResample != 0
}

// SetAlsaNoAutoResample sets whether auto resample is disabled for ALSA.
func (dc *DeviceConfig) SetAlsaNoAutoResample(disabled bool) {
	if dc != nil && dc.config != nil {
		dc.config.alsa.noAutoResample = C.ma_bool32(boolToCInt(disabled))
	}
}

// Pulse specific getters

// PulseStreamNamePlayback returns the PulseAudio playback stream name.
func (dc *DeviceConfig) PulseStreamNamePlayback() string {
	if dc == nil || dc.config == nil {
		return ""
	}
	if dc.config.pulse.pStreamNamePlayback == nil {
		return ""
	}
	return C.GoString(dc.config.pulse.pStreamNamePlayback)
}

// SetPulseStreamNamePlayback sets the PulseAudio playback stream name.
func (dc *DeviceConfig) SetPulseStreamNamePlayback(name string) {
	if dc != nil && dc.config != nil {
		cStr := C.CString(name)
		defer C.free(unsafe.Pointer(cStr))
		dc.config.pulse.pStreamNamePlayback = cStr
	}
}

// PulseStreamNameCapture returns the PulseAudio capture stream name.
func (dc *DeviceConfig) PulseStreamNameCapture() string {
	if dc == nil || dc.config == nil {
		return ""
	}
	if dc.config.pulse.pStreamNameCapture == nil {
		return ""
	}
	return C.GoString(dc.config.pulse.pStreamNameCapture)
}

// SetPulseStreamNameCapture sets the PulseAudio capture stream name.
func (dc *DeviceConfig) SetPulseStreamNameCapture(name string) {
	if dc != nil && dc.config != nil {
		cStr := C.CString(name)
		defer C.free(unsafe.Pointer(cStr))
		dc.config.pulse.pStreamNameCapture = cStr
	}
}

// PulseChannelMap returns the PulseAudio channel map setting.
func (dc *DeviceConfig) PulseChannelMap() int32 {
	if dc == nil || dc.config == nil {
		return 0
	}
	return int32(dc.config.pulse.channelMap)
}

// SetPulseChannelMap sets the PulseAudio channel map setting.
func (dc *DeviceConfig) SetPulseChannelMap(channelMap int32) {
	if dc != nil && dc.config != nil {
		dc.config.pulse.channelMap = C.int(channelMap)
	}
}

// CoreAudio specific getters

// CoreaudioAllowNominalSampleRateChange returns whether nominal sample rate change is allowed for CoreAudio.
func (dc *DeviceConfig) CoreaudioAllowNominalSampleRateChange() bool {
	if dc == nil || dc.config == nil {
		return false
	}
	return dc.config.coreaudio.allowNominalSampleRateChange != 0
}

// SetCoreaudioAllowNominalSampleRateChange sets whether nominal sample rate change is allowed for CoreAudio.
func (dc *DeviceConfig) SetCoreaudioAllowNominalSampleRateChange(allowed bool) {
	if dc != nil && dc.config != nil {
		dc.config.coreaudio.allowNominalSampleRateChange = C.ma_bool32(boolToCInt(allowed))
	}
}

// OpenSL specific getters

// OpenslStreamType returns the OpenSL stream type.
func (dc *DeviceConfig) OpenslStreamType() C.ma_opensl_stream_type {
	if dc == nil || dc.config == nil {
		return 0
	}
	return dc.config.opensl.streamType
}

// SetOpenslStreamType sets the OpenSL stream type.
func (dc *DeviceConfig) SetOpenslStreamType(streamType C.ma_opensl_stream_type) {
	if dc != nil && dc.config != nil {
		dc.config.opensl.streamType = streamType
	}
}

// OpenslRecordingPreset returns the OpenSL recording preset.
func (dc *DeviceConfig) OpenslRecordingPreset() C.ma_opensl_recording_preset {
	if dc == nil || dc.config == nil {
		return 0
	}
	return dc.config.opensl.recordingPreset
}

// SetOpenslRecordingPreset sets the OpenSL recording preset.
func (dc *DeviceConfig) SetOpenslRecordingPreset(preset C.ma_opensl_recording_preset) {
	if dc != nil && dc.config != nil {
		dc.config.opensl.recordingPreset = preset
	}
}

// OpenslEnableCompatibilityWorkarounds returns whether compatibility workarounds are enabled for OpenSL.
func (dc *DeviceConfig) OpenslEnableCompatibilityWorkarounds() bool {
	if dc == nil || dc.config == nil {
		return false
	}
	return dc.config.opensl.enableCompatibilityWorkarounds != 0
}

// SetOpenslEnableCompatibilityWorkarounds sets whether compatibility workarounds are enabled for OpenSL.
func (dc *DeviceConfig) SetOpenslEnableCompatibilityWorkarounds(enabled bool) {
	if dc != nil && dc.config != nil {
		dc.config.opensl.enableCompatibilityWorkarounds = C.ma_bool32(boolToCInt(enabled))
	}
}

// AAudio specific getters

// AaudioUsage returns the AAudio usage.
func (dc *DeviceConfig) AaudioUsage() C.ma_aaudio_usage {
	if dc == nil || dc.config == nil {
		return 0
	}
	return dc.config.aaudio.usage
}

// SetAaudioUsage sets the AAudio usage.
func (dc *DeviceConfig) SetAaudioUsage(usage C.ma_aaudio_usage) {
	if dc != nil && dc.config != nil {
		dc.config.aaudio.usage = usage
	}
}

// AaudioContentType returns the AAudio content type.
func (dc *DeviceConfig) AaudioContentType() C.ma_aaudio_content_type {
	if dc == nil || dc.config == nil {
		return 0
	}
	return dc.config.aaudio.contentType
}

// SetAaudioContentType sets the AAudio content type.
func (dc *DeviceConfig) SetAaudioContentType(contentType C.ma_aaudio_content_type) {
	if dc != nil && dc.config != nil {
		dc.config.aaudio.contentType = contentType
	}
}

// AaudioInputPreset returns the AAudio input preset.
func (dc *DeviceConfig) AaudioInputPreset() C.ma_aaudio_input_preset {
	if dc == nil || dc.config == nil {
		return 0
	}
	return dc.config.aaudio.inputPreset
}

// SetAaudioInputPreset sets the AAudio input preset.
func (dc *DeviceConfig) SetAaudioInputPreset(preset C.ma_aaudio_input_preset) {
	if dc != nil && dc.config != nil {
		dc.config.aaudio.inputPreset = preset
	}
}

// AaudioAllowedCapturePolicy returns the AAudio allowed capture policy.
func (dc *DeviceConfig) AaudioAllowedCapturePolicy() C.ma_aaudio_allowed_capture_policy {
	if dc == nil || dc.config == nil {
		return 0
	}
	return dc.config.aaudio.allowedCapturePolicy
}

// SetAaudioAllowedCapturePolicy sets the AAudio allowed capture policy.
func (dc *DeviceConfig) SetAaudioAllowedCapturePolicy(policy C.ma_aaudio_allowed_capture_policy) {
	if dc != nil && dc.config != nil {
		dc.config.aaudio.allowedCapturePolicy = policy
	}
}

// AaudioNoAutoStartAfterReroute returns whether auto start after reroute is disabled for AAudio.
func (dc *DeviceConfig) AaudioNoAutoStartAfterReroute() bool {
	if dc == nil || dc.config == nil {
		return false
	}
	return dc.config.aaudio.noAutoStartAfterReroute != 0
}

// SetAaudioNoAutoStartAfterReroute sets whether auto start after reroute is disabled for AAudio.
func (dc *DeviceConfig) SetAaudioNoAutoStartAfterReroute(disabled bool) {
	if dc != nil && dc.config != nil {
		dc.config.aaudio.noAutoStartAfterReroute = C.ma_bool32(boolToCInt(disabled))
	}
}

// AaudioEnableCompatibilityWorkarounds returns whether compatibility workarounds are enabled for AAudio.
func (dc *DeviceConfig) AaudioEnableCompatibilityWorkarounds() bool {
	if dc == nil || dc.config == nil {
		return false
	}
	return dc.config.aaudio.enableCompatibilityWorkarounds != 0
}

// SetAaudioEnableCompatibilityWorkarounds sets whether compatibility workarounds are enabled for AAudio.
func (dc *DeviceConfig) SetAaudioEnableCompatibilityWorkarounds(enabled bool) {
	if dc != nil && dc.config != nil {
		dc.config.aaudio.enableCompatibilityWorkarounds = C.ma_bool32(boolToCInt(enabled))
	}
}

// AaudioAllowSetBufferCapacity returns whether setting buffer capacity is allowed for AAudio.
func (dc *DeviceConfig) AaudioAllowSetBufferCapacity() bool {
	if dc == nil || dc.config == nil {
		return false
	}
	return dc.config.aaudio.allowSetBufferCapacity != 0
}

// SetAaudioAllowSetBufferCapacity sets whether setting buffer capacity is allowed for AAudio.
func (dc *DeviceConfig) SetAaudioAllowSetBufferCapacity(allowed bool) {
	if dc != nil && dc.config != nil {
		dc.config.aaudio.allowSetBufferCapacity = C.ma_bool32(boolToCInt(allowed))
	}
}

// Device wraps ma_device.
type Device struct {
	device *C.ma_device
}

// NewDeviceWithContext creates a new Device instance using the given context, initializes it with ma_device_init, and returns it.
func NewDeviceWithContext(ctx *Context, config *DeviceConfig) (*Device, error) {
	if ctx == nil || ctx.context == nil {
		return nil, ErrNilEngine
	}

	var cConfig *C.ma_device_config
	if config != nil {
		cConfig = config.config
	}

	d := (*C.ma_device)(C.malloc(C.size_t(unsafe.Sizeof(*(*C.ma_device)(nil)))))
	if d == nil {
		return nil, ErrResult(MA_NO_MEMORY)
	}
	C.memset(unsafe.Pointer(d), 0, C.size_t(unsafe.Sizeof(*d)))

	result := C.ma_device_init(ctx.context, cConfig, d)
	if result != 0 {
		C.free(unsafe.Pointer(d))
		return nil, ErrResult(result)
	}

	return &Device{device: d}, nil
}

// NewDevice creates a new Device instance without a context, initializes it with ma_device_init_ex, and returns it.
func NewDevice(backends []C.ma_backend, config *DeviceConfig) (*Device, error) {
	var cBackends *C.ma_backend
	var backendCount C.ma_uint32
	if len(backends) > 0 {
		cBackends = &backends[0]
		backendCount = C.ma_uint32(len(backends))
	}

	d := (*C.ma_device)(C.malloc(C.size_t(unsafe.Sizeof(*(*C.ma_device)(nil)))))
	if d == nil {
		return nil, ErrResult(MA_NO_MEMORY)
	}
	C.memset(unsafe.Pointer(d), 0, C.size_t(unsafe.Sizeof(*d)))

	var cConfig *C.ma_device_config
	if config != nil {
		cConfig = config.config
	}

	result := C.ma_device_init_ex(cBackends, backendCount, nil, cConfig, d)
	if result != 0 {
		C.free(unsafe.Pointer(d))
		return nil, ErrResult(result)
	}

	return &Device{device: d}, nil
}

// Close uninitializes and releases the device memory.
func (d *Device) Close() {
	if d.device == nil {
		return
	}
	C.ma_device_uninit(d.device)
	C.free(unsafe.Pointer(d.device))
	d.device = nil
}

// Start starts the device.
func (d *Device) Start() error {
	if d.device == nil {
		return ErrNilEngine
	}
	C.ma_device_start(d.device)
	return nil
}

// Stop stops the device.
func (d *Device) Stop() error {
	if d.device == nil {
		return ErrNilEngine
	}
	C.ma_device_stop(d.device)
	return nil
}

// Context returns the context of the device.
func (d *Device) Context() *Context {
	if d.device == nil {
		return nil
	}
	ctx := C.ma_device_get_context(d.device)
	return &Context{context: ctx}
}

// GetSampleRate returns the sample rate of the device.
func (d *Device) GetSampleRate() uint32 {
	if d.device == nil {
		return 0
	}
	return uint32(d.device.sampleRate)
}

// CaptureSampleRate returns the capture sample rate of the device.
func (d *Device) CaptureSampleRate() uint32 {
	if d.device == nil {
		return 0
	}
	return uint32(d.device.capture.internalSampleRate)
}

// Channels returns the playback channel count of the device.
func (d *Device) Channels() uint32 {
	if d.device == nil {
		return 0
	}
	return uint32(d.device.playback.channels)
}

// CaptureChannels returns the capture channel count of the device.
func (d *Device) CaptureChannels() uint32 {
	if d.device == nil {
		return 0
	}
	return uint32(d.device.capture.channels)
}

// Format returns the playback format of the device.
func (d *Device) Format() Format {
	if d.device == nil {
		return 0
	}
	return Format(d.device.playback.format)
}

// CaptureFormat returns the capture format of the device.
func (d *Device) CaptureFormat() Format {
	if d.device == nil {
		return 0
	}
	return Format(d.device.capture.format)
}
