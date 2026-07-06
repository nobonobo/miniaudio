package miniaudio

/*
#include "miniaudio.h"
#include <stdlib.h>
#include <string.h>
*/
import "C"
import "unsafe"

// EngineConfig wraps ma_engine_config.
type EngineConfig struct {
	config *C.ma_engine_config
}

// NewEngineConfig creates a new EngineConfig and calls ma_engine_config_init.
func NewEngineConfig() *EngineConfig {
	config := C.ma_engine_config_init()
	return &EngineConfig{config: &config}
}

func (ec *EngineConfig) Close() {
	if ec.config == nil {
		return
	}
	ec.config = nil
}

// GetListenerCount returns the listener count.
func (ec *EngineConfig) GetListenerCount() uint32 {
	if ec == nil || ec.config == nil {
		return 0
	}
	return uint32(ec.config.listenerCount)
}

// SetListenerCount sets the listener count.
func (ec *EngineConfig) SetListenerCount(count uint32) {
	if ec != nil && ec.config != nil {
		ec.config.listenerCount = C.ma_uint32(count)
	}
}

// GetChannels returns the channel count.
func (ec *EngineConfig) GetChannels() uint32 {
	if ec == nil || ec.config == nil {
		return 0
	}
	return uint32(ec.config.channels)
}

// SetChannels sets the channel count.
func (ec *EngineConfig) SetChannels(channels uint32) {
	if ec != nil && ec.config != nil {
		ec.config.channels = C.ma_uint32(channels)
	}
}

// GetSampleRate returns the sample rate.
func (ec *EngineConfig) GetSampleRate() uint32 {
	if ec == nil || ec.config == nil {
		return 0
	}
	return uint32(ec.config.sampleRate)
}

// SetSampleRate sets the sample rate.
func (ec *EngineConfig) SetSampleRate(sampleRate uint32) {
	if ec != nil && ec.config != nil {
		ec.config.sampleRate = C.ma_uint32(sampleRate)
	}
}

// GetPeriodSizeInFrames returns the period size in frames.
func (ec *EngineConfig) GetPeriodSizeInFrames() uint32 {
	if ec == nil || ec.config == nil {
		return 0
	}
	return uint32(ec.config.periodSizeInFrames)
}

// SetPeriodSizeInFrames sets the period size in frames.
func (ec *EngineConfig) SetPeriodSizeInFrames(periodSize uint32) {
	if ec != nil && ec.config != nil {
		ec.config.periodSizeInFrames = C.ma_uint32(periodSize)
	}
}

// GetPeriodSizeInMilliseconds returns the period size in milliseconds.
func (ec *EngineConfig) GetPeriodSizeInMilliseconds() uint32 {
	if ec == nil || ec.config == nil {
		return 0
	}
	return uint32(ec.config.periodSizeInMilliseconds)
}

// SetPeriodSizeInMilliseconds sets the period size in milliseconds.
func (ec *EngineConfig) SetPeriodSizeInMilliseconds(periodSize uint32) {
	if ec != nil && ec.config != nil {
		ec.config.periodSizeInMilliseconds = C.ma_uint32(periodSize)
	}
}

// GetGainSmoothTimeInFrames returns the gain smooth time in frames.
func (ec *EngineConfig) GetGainSmoothTimeInFrames() uint32 {
	if ec == nil || ec.config == nil {
		return 0
	}
	return uint32(ec.config.gainSmoothTimeInFrames)
}

// SetGainSmoothTimeInFrames sets the gain smooth time in frames.
func (ec *EngineConfig) SetGainSmoothTimeInFrames(frames uint32) {
	if ec != nil && ec.config != nil {
		(*ec.config).gainSmoothTimeInFrames = C.ma_uint32(frames)
	}
}

// GetGainSmoothTimeInMilliseconds returns the gain smooth time in milliseconds.
func (ec *EngineConfig) GetGainSmoothTimeInMilliseconds() uint32 {
	if ec == nil || ec.config == nil {
		return 0
	}
	return uint32(ec.config.gainSmoothTimeInMilliseconds)
}

// SetGainSmoothTimeInMilliseconds sets the gain smooth time in milliseconds.
func (ec *EngineConfig) SetGainSmoothTimeInMilliseconds(milliseconds uint32) {
	if ec != nil && ec.config != nil {
		ec.config.gainSmoothTimeInMilliseconds = C.ma_uint32(milliseconds)
	}
}

// GetDefaultVolumeSmoothTimeInPCMFrames returns the default volume smooth time in PCM frames.
func (ec *EngineConfig) GetDefaultVolumeSmoothTimeInPCMFrames() uint32 {
	if ec == nil || ec.config == nil {
		return 0
	}
	return uint32(ec.config.defaultVolumeSmoothTimeInPCMFrames)
}

// SetDefaultVolumeSmoothTimeInPCMFrames sets the default volume smooth time in PCM frames.
func (ec *EngineConfig) SetDefaultVolumeSmoothTimeInPCMFrames(frames uint32) {
	if ec != nil && ec.config != nil {
		ec.config.defaultVolumeSmoothTimeInPCMFrames = C.ma_uint32(frames)
	}
}

// GetPreMixStackSizeInBytes returns the pre-mix stack size in bytes.
func (ec *EngineConfig) GetPreMixStackSizeInBytes() uint32 {
	if ec == nil || ec.config == nil {
		return 0
	}
	return uint32(ec.config.preMixStackSizeInBytes)
}

// SetPreMixStackSizeInBytes sets the pre-mix stack size in bytes.
func (ec *EngineConfig) SetPreMixStackSizeInBytes(size uint32) {
	if ec != nil && ec.config != nil {
		ec.config.preMixStackSizeInBytes = C.ma_uint32(size)
	}
}

// GetNoAutoStart returns whether auto start is disabled.
func (ec *EngineConfig) GetNoAutoStart() bool {
	if ec == nil || ec.config == nil {
		return false
	}
	return ec.config.noAutoStart != 0
}

// SetNoAutoStart sets whether auto start is disabled.
func (ec *EngineConfig) SetNoAutoStart(noAutoStart bool) {
	if ec != nil && ec.config != nil {
		ec.config.noAutoStart = C.ma_bool32(boolToCInt(noAutoStart))
	}
}

// GetNoDevice returns whether device creation is disabled.
func (ec *EngineConfig) GetNoDevice() bool {
	if ec == nil || ec.config == nil {
		return false
	}
	return ec.config.noDevice != 0
}

// SetNoDevice sets whether device creation is disabled.
func (ec *EngineConfig) SetNoDevice(noDevice bool) {
	if ec != nil && ec.config != nil {
		(*ec.config).noDevice = C.ma_bool32(boolToCInt(noDevice))
	}
}

// GetMonoExpansionMode returns the mono expansion mode.
func (ec *EngineConfig) GetMonoExpansionMode() uint32 {
	if ec == nil || ec.config == nil {
		return 0
	}
	return uint32(ec.config.monoExpansionMode)
}

// SetMonoExpansionMode sets the mono expansion mode.
func (ec *EngineConfig) SetMonoExpansionMode(mode uint32) {
	if ec != nil && ec.config != nil {
		ec.config.monoExpansionMode = C.ma_mono_expansion_mode(mode)
	}
}

// GetPLog returns the log pointer.
func (ec *EngineConfig) GetPLog() *Log {
	if ec == nil || ec.config == nil {
		return nil
	}
	return &Log{log: ec.config.pLog}
}

// SetPLog sets the log pointer.
func (ec *EngineConfig) SetPLog(log *Log) {
	if ec != nil && ec.config != nil {
		ec.config.pLog = log.log
	}
}

// GetPContext returns the context pointer.
func (ec *EngineConfig) GetPContext() *Context {
	if ec == nil || ec.config == nil {
		return nil
	}
	return &Context{context: ec.config.pContext}
}

// SetPContext sets the context pointer.
func (ec *EngineConfig) SetPContext(ctx *Context) {
	if ec != nil && ec.config != nil {
		ec.config.pContext = ctx.context
	}
}

// GetPDevice returns the device pointer.
func (ec *EngineConfig) GetPDevice() *Device {
	if ec == nil || ec.config == nil {
		return nil
	}
	return &Device{device: ec.config.pDevice}
}

// SetPDevice sets the device pointer.
func (ec *EngineConfig) SetPDevice(device *Device) {
	if ec != nil && ec.config != nil {
		ec.config.pDevice = device.device
	}
}

// GetPPaymentDeviceID returns the playback device ID pointer.
func (ec *EngineConfig) GetPPaymentDeviceID() *DeviceID {
	if ec == nil || ec.config == nil {
		return nil
	}
	return &DeviceID{id: ec.config.pPlaybackDeviceID}
}

// SetPPaymentDeviceID sets the playback device ID pointer.
func (ec *EngineConfig) SetPPaymentDeviceID(deviceID *DeviceID) {
	if ec != nil && ec.config != nil {
		ec.config.pPlaybackDeviceID = deviceID.id
	}
}

// GetDataCallback returns the data callback.
func (ec *EngineConfig) GetDataCallback() C.ma_device_data_proc {
	if ec == nil || ec.config == nil {
		return nil
	}
	return ec.config.dataCallback
}

// SetDataCallback sets the data callback.
func (ec *EngineConfig) SetDataCallback(callback C.ma_device_data_proc) {
	if ec != nil && ec.config != nil {
		ec.config.dataCallback = callback
	}
}

// GetNotificationCallback returns the notification callback.
func (ec *EngineConfig) GetNotificationCallback() C.ma_device_notification_proc {
	if ec == nil || ec.config == nil {
		return nil
	}
	return ec.config.notificationCallback
}

// SetNotificationCallback sets the notification callback.
func (ec *EngineConfig) SetNotificationCallback(callback C.ma_device_notification_proc) {
	if ec != nil && ec.config != nil {
		ec.config.notificationCallback = callback
	}
}

// GetOnProcess returns the on process callback.
func (ec *EngineConfig) GetOnProcess() C.ma_engine_process_proc {
	if ec == nil || ec.config == nil {
		return nil
	}
	return ec.config.onProcess
}

// SetOnProcess sets the on process callback.
func (ec *EngineConfig) SetOnProcess(callback C.ma_engine_process_proc) {
	if ec != nil && ec.config != nil {
		ec.config.onProcess = callback
	}
}

// GetPProcessUserData returns the user data for onProcess.
func (ec *EngineConfig) GetPProcessUserData() unsafe.Pointer {
	if ec == nil || ec.config == nil {
		return nil
	}
	return unsafe.Pointer(ec.config.pProcessUserData)
}

// SetPProcessUserData sets the user data for onProcess.
func (ec *EngineConfig) SetPProcessUserData(userData unsafe.Pointer) {
	if ec != nil && ec.config != nil {
		ec.config.pProcessUserData = userData
	}
}

// Engine is a high-level audio engine that wraps around the resource manager and node graph.
type Engine struct {
	engine *C.ma_engine
}

// NewEngine creates a new Engine instance, initializes it with ma_engine_init, and returns it.
func NewEngine(config *EngineConfig) (*Engine, error) {
	e := (*C.ma_engine)(C.malloc(C.size_t(unsafe.Sizeof(*(*C.ma_engine)(nil)))))
	if e == nil {
		return nil, ErrNilEngine
	}
	C.memset(unsafe.Pointer(e), 0, C.size_t(unsafe.Sizeof(*e)))

	var cConfig *C.ma_engine_config
	if config != nil {
		cConfig = config.config
	}

	result := C.ma_engine_init(cConfig, e)
	if result != 0 {
		C.free(unsafe.Pointer(e))
		return nil, ErrResult(result)
	}

	return &Engine{engine: e}, nil
}

// Close uninitializes and releases the engine memory.
func (e *Engine) Close() {
	if e.engine == nil {
		return
	}
	C.ma_engine_uninit(e.engine)
	C.free(unsafe.Pointer(e.engine))
	e.engine = nil
}

// Start starts the engine.
func (e *Engine) Start() error {
	if e.engine == nil {
		return ErrNilEngine
	}
	result := C.ma_engine_start(e.engine)
	if result != 0 {
		return ErrResult(result)
	}
	return nil
}

// Stop stops the engine.
func (e *Engine) Stop() error {
	if e.engine == nil {
		return ErrNilEngine
	}
	result := C.ma_engine_stop(e.engine)
	if result != 0 {
		return ErrResult(result)
	}
	return nil
}

// PlaySound plays a sound with the given file path.
func (e *Engine) PlaySound(filePath string) error {
	if e.engine == nil {
		return ErrNilEngine
	}
	cPath := C.CString(filePath)
	defer C.free(unsafe.Pointer(cPath))

	result := C.ma_engine_play_sound(e.engine, cPath, nil)
	if result != 0 {
		return ErrResult(result)
	}
	return nil
}

// GetSampleRate returns the sample rate of the engine.
func (e *Engine) GetSampleRate() uint32 {
	if e.engine == nil {
		return 0
	}
	return uint32(C.ma_engine_get_sample_rate(e.engine))
}

// GetChannels returns the number of channels of the engine.
func (e *Engine) GetChannels() uint32 {
	if e.engine == nil {
		return 0
	}
	return uint32(C.ma_engine_get_channels(e.engine))
}

// GetTimeInPCMFrames returns the current global time in PCM frames.
func (e *Engine) GetTimeInPCMFrames() uint64 {
	if e.engine == nil {
		return 0
	}
	return uint64(C.ma_engine_get_time_in_pcm_frames(e.engine))
}

// SetTimeInPCMFrames sets the current global time in PCM frames.
func (e *Engine) SetTimeInPCMFrames(time uint64) error {
	if e.engine == nil {
		return ErrNilEngine
	}
	result := C.ma_engine_set_time_in_pcm_frames(e.engine, C.ma_uint64(time))
	if result != 0 {
		return ErrResult(result)
	}
	return nil
}

// SetVolume sets the engine volume.
func (e *Engine) SetVolume(volume float32) error {
	if e.engine == nil {
		return ErrNilEngine
	}
	result := C.ma_engine_set_volume(e.engine, C.float(volume))
	if result != 0 {
		return ErrResult(result)
	}
	return nil
}

// SetGainDB sets the engine gain in dB.
func (e *Engine) SetGainDB(gainDB float32) error {
	if e.engine == nil {
		return ErrNilEngine
	}
	result := C.ma_engine_set_gain_db(e.engine, C.float(gainDB))
	if result != 0 {
		return ErrResult(result)
	}
	return nil
}

// ListenerSetPosition sets the position of the listener.
func (e *Engine) ListenerSetPosition(listenerIndex uint32, x, y, z float32) error {
	if e.engine == nil {
		return ErrNilEngine
	}
	C.ma_engine_listener_set_position(e.engine, C.ma_uint32(listenerIndex), C.float(x), C.float(y), C.float(z))
	return nil
}

// ListenerSetDirection sets the direction of the listener.
func (e *Engine) ListenerSetDirection(listenerIndex uint32, x, y, z float32) error {
	if e.engine == nil {
		return ErrNilEngine
	}
	C.ma_engine_listener_set_direction(e.engine, C.ma_uint32(listenerIndex), C.float(x), C.float(y), C.float(z))
	return nil
}

// ListenerSetVelocity sets the velocity of the listener (for Doppler effect).
func (e *Engine) ListenerSetVelocity(listenerIndex uint32, x, y, z float32) error {
	if e.engine == nil {
		return ErrNilEngine
	}
	C.ma_engine_listener_set_velocity(e.engine, C.ma_uint32(listenerIndex), C.float(x), C.float(y), C.float(z))
	return nil
}

// ListenerGetPosition returns the position of the listener.
func (e *Engine) ListenerGetPosition(listenerIndex uint32) (x, y, z float32) {
	if e.engine == nil {
		return 0, 0, 0
	}
	pos := C.ma_engine_listener_get_position(e.engine, C.ma_uint32(listenerIndex))
	return float32(pos.x), float32(pos.y), float32(pos.z)
}

// ListenerGetDirection returns the direction of the listener.
func (e *Engine) ListenerGetDirection(listenerIndex uint32) (x, y, z float32) {
	if e.engine == nil {
		return 0, 0, 0
	}
	dir := C.ma_engine_listener_get_direction(e.engine, C.ma_uint32(listenerIndex))
	return float32(dir.x), float32(dir.y), float32(dir.z)
}

// ListenerGetVelocity returns the velocity of the listener.
func (e *Engine) ListenerGetVelocity(listenerIndex uint32) (x, y, z float32) {
	if e.engine == nil {
		return 0, 0, 0
	}
	vel := C.ma_engine_listener_get_velocity(e.engine, C.ma_uint32(listenerIndex))
	return float32(vel.x), float32(vel.y), float32(vel.z)
}

// ListenerSetWorldUp sets the world up vector of the listener.
func (e *Engine) ListenerSetWorldUp(listenerIndex uint32, x, y, z float32) error {
	if e.engine == nil {
		return ErrNilEngine
	}
	C.ma_engine_listener_set_world_up(e.engine, C.ma_uint32(listenerIndex), C.float(x), C.float(y), C.float(z))
	return nil
}

// ListenerSetEnabled enables or disables the listener.
func (e *Engine) ListenerSetEnabled(listenerIndex uint32, enabled bool) error {
	if e.engine == nil {
		return ErrNilEngine
	}
	C.ma_engine_listener_set_enabled(e.engine, C.ma_bool32(boolToCInt(enabled)), 0)
	return nil
}

// ListenerIsEnabled returns whether the listener is enabled.
func (e *Engine) ListenerIsEnabled(listenerIndex uint32) bool {
	if e.engine == nil {
		return false
	}
	return C.ma_engine_listener_is_enabled(e.engine, C.ma_uint32(listenerIndex)) != 0
}

// GetDevice returns the underlying device.
func (e *Engine) GetDevice() *C.ma_device {
	if e.engine == nil {
		return nil
	}
	return C.ma_engine_get_device(e.engine)
}

// PlaySoundFromFile creates and starts a sound from a file path.
// Note: This function requires ma_sound_init_from_file which is defined in miniaudio.h
// under #ifndef MA_NO_RESOURCE_MANAGER. Currently, this function is not implemented
// because miniaudio.c does not include the implementation.
func (e *Engine) PlaySoundFromFile(filePath string, flags SoundFlags, group *Sound, doneFence unsafe.Pointer) (*Sound, error) {
	if e.engine == nil {
		return nil, ErrNilEngine
	}
	cPath := C.CString(filePath)
	defer C.free(unsafe.Pointer(cPath))

	s := (*C.ma_sound)(C.malloc(C.size_t(unsafe.Sizeof(*(*C.ma_sound)(nil)))))
	if s == nil {
		return nil, ErrNilEngine
	}
	C.memset(unsafe.Pointer(s), 0, C.size_t(unsafe.Sizeof(*s)))

	var cGroup *C.ma_sound
	if group != nil && group.sound != nil {
		cGroup = group.sound
	}

	var cFence *C.ma_fence
	if doneFence != nil {
		cFence = (*C.ma_fence)(doneFence)
	}

	result := C.ma_sound_init_from_file(e.engine, cPath, C.ma_uint32(flags), cGroup, cFence, s)
	if result != 0 {
		C.free(unsafe.Pointer(s))
		return nil, ErrResult(result)
	}

	return &Sound{sound: s}, nil
}

// PlaySoundFromDataSource creates and starts a sound from a data source.
// Note: This function requires ma_sound_init_from_data_source which is defined in miniaudio.h
// under #ifndef MA_NO_RESOURCE_MANAGER. Currently, this function is not implemented
// because miniaudio.c does not include the implementation.
func (e *Engine) PlaySoundFromDataSource(dataSource DataSource, flags SoundFlags, group *Sound) (*Sound, error) {
	if e.engine == nil {
		return nil, ErrNilEngine
	}

	s := (*C.ma_sound)(C.malloc(C.size_t(unsafe.Sizeof(*(*C.ma_sound)(nil)))))
	if s == nil {
		return nil, ErrNilEngine
	}
	C.memset(unsafe.Pointer(s), 0, C.size_t(unsafe.Sizeof(*s)))

	var cGroup *C.ma_sound
	if group != nil && group.sound != nil {
		cGroup = group.sound
	}

	result := C.ma_sound_init_from_data_source(e.engine, dataSource.dataSourcePtr(), C.ma_uint32(flags), cGroup, s)
	if result != 0 {
		C.free(unsafe.Pointer(s))
		return nil, ErrResult(result)
	}

	return &Sound{sound: s}, nil
}

// CopySound creates a copy of an existing sound.
// Note: This function requires ma_sound_init_copy which is defined in miniaudio.h
// under #ifndef MA_NO_RESOURCE_MANAGER. Currently, this function is not implemented
// because miniaudio.c does not include the implementation.
func (e *Engine) CopySound(existing *Sound, flags SoundFlags, group *Sound) (*Sound, error) {
	if e.engine == nil {
		return nil, ErrNilEngine
	}
	if existing == nil || existing.sound == nil {
		return nil, ErrNilEngine
	}

	s := (*C.ma_sound)(C.malloc(C.size_t(unsafe.Sizeof(*(*C.ma_sound)(nil)))))
	if s == nil {
		return nil, ErrNilEngine
	}
	C.memset(unsafe.Pointer(s), 0, C.size_t(unsafe.Sizeof(*s)))

	var cGroup *C.ma_sound
	if group != nil && group.sound != nil {
		cGroup = group.sound
	}

	result := C.ma_sound_init_copy(e.engine, existing.sound, C.ma_uint32(flags), cGroup, s)
	if result != 0 {
		C.free(unsafe.Pointer(s))
		return nil, ErrResult(result)
	}

	return &Sound{sound: s}, nil
}
