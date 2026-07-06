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

// GetSampleRate returns the sample rate.
func (dc *DeviceConfig) GetSampleRate() uint32 {
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

// GetPeriodSizeInFrames returns the period size in frames.
func (dc *DeviceConfig) GetPeriodSizeInFrames() uint32 {
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

// GetPeriodSizeInMilliseconds returns the period size in milliseconds.
func (dc *DeviceConfig) GetPeriodSizeInMilliseconds() uint32 {
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

// GetPeriods returns the periods count.
func (dc *DeviceConfig) GetPeriods() uint32 {
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

// GetPerformanceProfile returns the performance profile.
func (dc *DeviceConfig) GetPerformanceProfile() C.ma_performance_profile {
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

// GetNoPreSilencedOutputBuffer returns whether pre-silencing of output buffer is disabled.
func (dc *DeviceConfig) GetNoPreSilencedOutputBuffer() bool {
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

// GetNoClip returns whether clipping is disabled.
func (dc *DeviceConfig) GetNoClip() bool {
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

// GetNoDisableDenormals returns whether denormal disabling is disabled.
func (dc *DeviceConfig) GetNoDisableDenormals() bool {
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

// GetNoFixedSizedCallback returns whether fixed-sized callback disabling is enabled.
func (dc *DeviceConfig) GetNoFixedSizedCallback() bool {
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

// GetDataCallback returns the data callback.
func (dc *DeviceConfig) GetDataCallback() C.ma_device_data_proc {
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

// GetNotificationCallback returns the notification callback.
func (dc *DeviceConfig) GetNotificationCallback() C.ma_device_notification_proc {
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

// GetStopCallback returns the stop callback.
func (dc *DeviceConfig) GetStopCallback() C.ma_stop_proc {
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

// GetPUserData returns the user data pointer.
func (dc *DeviceConfig) GetPUserData() unsafe.Pointer {
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

// GetPlaybackDeviceID returns the playback device ID.
func (dc *DeviceConfig) GetPlaybackDeviceID() *DeviceID {
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

// GetPlaybackFormat returns the playback format.
func (dc *DeviceConfig) GetPlaybackFormat() Format {
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

// GetPlaybackChannels returns the playback channel count.
func (dc *DeviceConfig) GetPlaybackChannels() uint32 {
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

// GetPlaybackPChannelMap returns the playback channel map pointer.
func (dc *DeviceConfig) GetPlaybackPChannelMap() unsafe.Pointer {
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

// GetPlaybackChannelMixMode returns the playback channel mix mode.
func (dc *DeviceConfig) GetPlaybackChannelMixMode() ChannelMixMode {
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

// GetPlaybackCalculateLFEFromSpatialChannels returns whether LFE calculation from spatial channels is enabled.
func (dc *DeviceConfig) GetPlaybackCalculateLFEFromSpatialChannels() bool {
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

// GetPlaybackShareMode returns the playback share mode.
func (dc *DeviceConfig) GetPlaybackShareMode() C.ma_share_mode {
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

// GetCaptureDeviceID returns the capture device ID.
func (dc *DeviceConfig) GetCaptureDeviceID() *DeviceID {
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

// GetCaptureFormat returns the capture format.
func (dc *DeviceConfig) GetCaptureFormat() Format {
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

// GetCaptureChannels returns the capture channel count.
func (dc *DeviceConfig) GetCaptureChannels() uint32 {
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

// GetCapturePChannelMap returns the capture channel map pointer.
func (dc *DeviceConfig) GetCapturePChannelMap() unsafe.Pointer {
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

// GetCaptureChannelMixMode returns the capture channel mix mode.
func (dc *DeviceConfig) GetCaptureChannelMixMode() ChannelMixMode {
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

// GetCaptureCalculateLFEFromSpatialChannels returns whether LFE calculation from spatial channels is enabled.
func (dc *DeviceConfig) GetCaptureCalculateLFEFromSpatialChannels() bool {
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

// GetCaptureShareMode returns the capture share mode.
func (dc *DeviceConfig) GetCaptureShareMode() C.ma_share_mode {
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

// GetWasapiNoAutoConvertSRC returns whether auto-convert SRC is disabled for WASAPI.
func (dc *DeviceConfig) GetWasapiNoAutoConvertSRC() bool {
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

// GetWasapiNoDefaultQualitySRC returns whether default quality SRC is disabled for WASAPI.
func (dc *DeviceConfig) GetWasapiNoDefaultQualitySRC() bool {
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

// GetWasapiLoopbackProcessID returns the WASAPI loopback process ID.
func (dc *DeviceConfig) GetWasapiLoopbackProcessID() uint32 {
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

// GetWasapiNoHardwareOffloading returns whether hardware offloading is disabled for WASAPI.
func (dc *DeviceConfig) GetWasapiNoHardwareOffloading() bool {
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

// GetAlsaNoMMap returns whether MMap is disabled for ALSA.
func (dc *DeviceConfig) GetAlsaNoMMap() bool {
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

// GetAlsaNoAutoFormat returns whether auto format is disabled for ALSA.
func (dc *DeviceConfig) GetAlsaNoAutoFormat() bool {
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

// GetAlsaNoAutoChannels returns whether auto channels is disabled for ALSA.
func (dc *DeviceConfig) GetAlsaNoAutoChannels() bool {
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

// GetAlsaNoAutoResample returns whether auto resample is disabled for ALSA.
func (dc *DeviceConfig) GetAlsaNoAutoResample() bool {
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

// GetPulseStreamNamePlayback returns the PulseAudio playback stream name.
func (dc *DeviceConfig) GetPulseStreamNamePlayback() string {
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

// GetPulseStreamNameCapture returns the PulseAudio capture stream name.
func (dc *DeviceConfig) GetPulseStreamNameCapture() string {
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

// GetPulseChannelMap returns the PulseAudio channel map setting.
func (dc *DeviceConfig) GetPulseChannelMap() int32 {
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

// GetCoreaudioAllowNominalSampleRateChange returns whether nominal sample rate change is allowed for CoreAudio.
func (dc *DeviceConfig) GetCoreaudioAllowNominalSampleRateChange() bool {
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

// GetOpenslStreamType returns the OpenSL stream type.
func (dc *DeviceConfig) GetOpenslStreamType() C.ma_opensl_stream_type {
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

// GetOpenslRecordingPreset returns the OpenSL recording preset.
func (dc *DeviceConfig) GetOpenslRecordingPreset() C.ma_opensl_recording_preset {
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

// GetOpenslEnableCompatibilityWorkarounds returns whether compatibility workarounds are enabled for OpenSL.
func (dc *DeviceConfig) GetOpenslEnableCompatibilityWorkarounds() bool {
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

// GetAaudioUsage returns the AAudio usage.
func (dc *DeviceConfig) GetAaudioUsage() C.ma_aaudio_usage {
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

// GetAaudioContentType returns the AAudio content type.
func (dc *DeviceConfig) GetAaudioContentType() C.ma_aaudio_content_type {
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

// GetAaudioInputPreset returns the AAudio input preset.
func (dc *DeviceConfig) GetAaudioInputPreset() C.ma_aaudio_input_preset {
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

// GetAaudioAllowedCapturePolicy returns the AAudio allowed capture policy.
func (dc *DeviceConfig) GetAaudioAllowedCapturePolicy() C.ma_aaudio_allowed_capture_policy {
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

// GetAaudioNoAutoStartAfterReroute returns whether auto start after reroute is disabled for AAudio.
func (dc *DeviceConfig) GetAaudioNoAutoStartAfterReroute() bool {
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

// GetAaudioEnableCompatibilityWorkarounds returns whether compatibility workarounds are enabled for AAudio.
func (dc *DeviceConfig) GetAaudioEnableCompatibilityWorkarounds() bool {
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

// GetAaudioAllowSetBufferCapacity returns whether setting buffer capacity is allowed for AAudio.
func (dc *DeviceConfig) GetAaudioAllowSetBufferCapacity() bool {
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

// GetContext returns the context of the device.
func (d *Device) GetContext() *Context {
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

// GetCaptureSampleRate returns the capture sample rate of the device.
func (d *Device) GetCaptureSampleRate() uint32 {
	if d.device == nil {
		return 0
	}
	return uint32(d.device.capture.internalSampleRate)
}

// GetChannels returns the playback channel count of the device.
func (d *Device) GetChannels() uint32 {
	if d.device == nil {
		return 0
	}
	return uint32(d.device.playback.channels)
}

// GetCaptureChannels returns the capture channel count of the device.
func (d *Device) GetCaptureChannels() uint32 {
	if d.device == nil {
		return 0
	}
	return uint32(d.device.capture.channels)
}

// GetFormat returns the playback format of the device.
func (d *Device) GetFormat() Format {
	if d.device == nil {
		return 0
	}
	return Format(d.device.playback.format)
}

// GetCaptureFormat returns the capture format of the device.
func (d *Device) GetCaptureFormat() Format {
	if d.device == nil {
		return 0
	}
	return Format(d.device.capture.format)
}
