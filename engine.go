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

// ListenerCount returns the listener count.
func (ec *EngineConfig) ListenerCount() uint32 {
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

// Channels returns the channel count.
func (ec *EngineConfig) Channels() uint32 {
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

// SampleRate returns the sample rate.
func (ec *EngineConfig) SampleRate() uint32 {
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

// PeriodSizeInFrames returns the period size in frames.
func (ec *EngineConfig) PeriodSizeInFrames() uint32 {
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

// PeriodSizeInMilliseconds returns the period size in milliseconds.
func (ec *EngineConfig) PeriodSizeInMilliseconds() uint32 {
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

// GainSmoothTimeInFrames returns the gain smooth time in frames.
func (ec *EngineConfig) GainSmoothTimeInFrames() uint32 {
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

// GainSmoothTimeInMilliseconds returns the gain smooth time in milliseconds.
func (ec *EngineConfig) GainSmoothTimeInMilliseconds() uint32 {
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

// DefaultVolumeSmoothTimeInPCMFrames returns the default volume smooth time in PCM frames.
func (ec *EngineConfig) DefaultVolumeSmoothTimeInPCMFrames() uint32 {
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

// PreMixStackSizeInBytes returns the pre-mix stack size in bytes.
func (ec *EngineConfig) PreMixStackSizeInBytes() uint32 {
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

// NoAutoStart returns whether auto start is disabled.
func (ec *EngineConfig) NoAutoStart() bool {
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

// NoDevice returns whether device creation is disabled.
func (ec *EngineConfig) NoDevice() bool {
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

// MonoExpansionMode returns the mono expansion mode.
func (ec *EngineConfig) MonoExpansionMode() uint32 {
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

// PLog returns the log pointer.
func (ec *EngineConfig) PLog() *Log {
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

// PContext returns the context pointer.
func (ec *EngineConfig) PContext() *Context {
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

// PDevice returns the device pointer.
func (ec *EngineConfig) PDevice() *Device {
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

// PaymentDeviceID returns the playback device ID pointer.
func (ec *EngineConfig) PaymentDeviceID() *DeviceID {
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

// DataCallback returns the data callback.
func (ec *EngineConfig) DataCallback() C.ma_device_data_proc {
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

// NotificationCallback returns the notification callback.
func (ec *EngineConfig) NotificationCallback() C.ma_device_notification_proc {
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

// OnProcess returns the on process callback.
func (ec *EngineConfig) OnProcess() C.ma_engine_process_proc {
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

// PProcessUserData returns the user data for onProcess.
func (ec *EngineConfig) PProcessUserData() unsafe.Pointer {
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

// SampleRate returns the sample rate of the engine.
func (e *Engine) SampleRate() uint32 {
	if e.engine == nil {
		return 0
	}
	return uint32(C.ma_engine_get_sample_rate(e.engine))
}

// Channels returns the number of channels of the engine.
func (e *Engine) Channels() uint32 {
	if e.engine == nil {
		return 0
	}
	return uint32(C.ma_engine_get_channels(e.engine))
}

// TimeInPCMFrames returns the current global time in PCM frames.
func (e *Engine) TimeInPCMFrames() uint64 {
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

// Volume returns the engine volume.
func (e *Engine) Volume() float32 {
	if e.engine == nil {
		return 0
	}
	return float32(C.ma_engine_get_volume(e.engine))
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

// GainDB returns the engine gain in dB.
func (e *Engine) GainDB() float32 {
	if e.engine == nil {
		return 0
	}
	return float32(C.ma_engine_get_gain_db(e.engine))
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

// ListenerCount returns the number of listeners.
func (e *Engine) ListenerCount() uint32 {
	if e.engine == nil {
		return 0
	}
	return uint32(C.ma_engine_get_listener_count(e.engine))
}

// FindClosestListener finds the closest listener to the given position.
func (e *Engine) FindClosestListener(x, y, z float32) uint32 {
	if e.engine == nil {
		return 0
	}
	return uint32(C.ma_engine_find_closest_listener(e.engine, C.float(x), C.float(y), C.float(z)))
}

// ListenerPosition returns the position of the listener.
func (e *Engine) ListenerPosition(listenerIndex uint32) (x, y, z float32) {
	if e.engine == nil {
		return 0, 0, 0
	}
	pos := C.ma_engine_listener_get_position(e.engine, C.ma_uint32(listenerIndex))
	return float32(pos.x), float32(pos.y), float32(pos.z)
}

// ListenerDirection returns the direction of the listener.
func (e *Engine) ListenerDirection(listenerIndex uint32) (x, y, z float32) {
	if e.engine == nil {
		return 0, 0, 0
	}
	dir := C.ma_engine_listener_get_direction(e.engine, C.ma_uint32(listenerIndex))
	return float32(dir.x), float32(dir.y), float32(dir.z)
}

// ListenerVelocity returns the velocity of the listener.
func (e *Engine) ListenerVelocity(listenerIndex uint32) (x, y, z float32) {
	if e.engine == nil {
		return 0, 0, 0
	}
	vel := C.ma_engine_listener_get_velocity(e.engine, C.ma_uint32(listenerIndex))
	return float32(vel.x), float32(vel.y), float32(vel.z)
}

// ListenerCone returns the cone of the listener.
func (e *Engine) ListenerCone(listenerIndex uint32) (innerAngle, outerAngle, outerGain float32) {
	if e.engine == nil {
		return 0, 0, 0
	}
	var _innerAngle, _outerAngle, _outerGain C.float
	C.ma_engine_listener_get_cone(e.engine, C.ma_uint32(listenerIndex), &_innerAngle, &_outerAngle, &_outerGain)
	return float32(_innerAngle), float32(_outerAngle), float32(_outerGain)
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

// ListenerSetCone sets the cone of the listener.
func (e *Engine) ListenerSetCone(listenerIndex uint32, innerAngle, outerAngle, outerGain float32) error {
	if e.engine == nil {
		return ErrNilEngine
	}
	C.ma_engine_listener_set_cone(e.engine, C.ma_uint32(listenerIndex), C.float(innerAngle), C.float(outerAngle), C.float(outerGain))
	return nil
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
	C.ma_engine_listener_set_enabled(e.engine, C.ma_uint32(listenerIndex), C.ma_bool32(boolToCInt(enabled)))
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

// ReadPCMFrames reads PCM frames from the engine.
func (e *Engine) ReadPCMFrames(frameCount uint64, pFramesOut []float32) (framesRead uint64, err error) {
	if e.engine == nil {
		return 0, ErrNilEngine
	}
	var framesReadC C.ma_uint64
	var outPtr unsafe.Pointer
	//var lenC C.size_t
	if len(pFramesOut) > 0 {
		outPtr = unsafe.Pointer(&pFramesOut[0])
		//lenC = C.size_t(len(pFramesOut))
	}
	result := C.ma_engine_read_pcm_frames(e.engine, outPtr, C.ma_uint64(frameCount), &framesReadC)
	if result != 0 {
		return 0, ErrResult(result)
	}
	framesRead = uint64(framesReadC)
	return framesRead, nil
}

// NodeGraph returns the underlying node graph.
func (e *Engine) NodeGraph() *C.ma_node_graph {
	if e.engine == nil {
		return nil
	}
	return C.ma_engine_get_node_graph(e.engine)
}

// ResourceManager returns the underlying resource manager.
func (e *Engine) ResourceManager() *C.ma_resource_manager {
	if e.engine == nil {
		return nil
	}
	return C.ma_engine_get_resource_manager(e.engine)
}

// Log returns the underlying log.
func (e *Engine) Log() *C.ma_log {
	if e.engine == nil {
		return nil
	}
	return C.ma_engine_get_log(e.engine)
}

// Endpoint returns the endpoint node.
func (e *Engine) Endpoint() *Node {
	if e.engine == nil {
		return nil
	}
	return &Node{node: (*C.ma_node)(C.ma_engine_get_endpoint(e.engine))}
}

// TimeInMilliseconds returns the current global time in milliseconds.
func (e *Engine) TimeInMilliseconds() uint64 {
	if e.engine == nil {
		return 0
	}
	return uint64(C.ma_engine_get_time_in_milliseconds(e.engine))
}

// SetTimeInMilliseconds sets the current global time in milliseconds.
func (e *Engine) SetTimeInMilliseconds(time uint64) error {
	if e.engine == nil {
		return ErrNilEngine
	}
	result := C.ma_engine_set_time_in_milliseconds(e.engine, C.ma_uint64(time))
	if result != 0 {
		return ErrResult(result)
	}
	return nil
}

// Time returns the current global time in PCM frames (deprecated).
func (e *Engine) Time() uint64 {
	if e.engine == nil {
		return 0
	}
	return uint64(C.ma_engine_get_time(e.engine))
}

// SetTime sets the global time in PCM frames (deprecated).
func (e *Engine) SetTime(time uint64) error {
	if e.engine == nil {
		return ErrNilEngine
	}
	result := C.ma_engine_set_time(e.engine, C.ma_uint64(time))
	if result != 0 {
		return ErrResult(result)
	}
	return nil
}

// PlaySoundEx plays a sound with the given file path, parent node, and input bus index.
func (e *Engine) PlaySoundEx(filePath string, pNode *Node, nodeInputBusIndex uint32) error {
	if e.engine == nil {
		return ErrNilEngine
	}
	cPath := C.CString(filePath)
	defer C.free(unsafe.Pointer(cPath))
	var node unsafe.Pointer
	if pNode != nil {
		node = unsafe.Pointer(pNode.node)
	}
	result := C.ma_engine_play_sound_ex(e.engine, cPath, node, C.ma_uint32(nodeInputBusIndex))
	if result != 0 {
		return ErrResult(result)
	}
	return nil
}

// PlaySoundFromFile creates and starts a sound from a file path.
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
