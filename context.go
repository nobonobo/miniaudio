package miniaudio

/*
#include "miniaudio.h"
#include <stdlib.h>
#include <string.h>
*/
import "C"
import "unsafe"

// ContextConfig wraps ma_context_config.
type ContextConfig struct {
	config *C.ma_context_config
}

// NewContextConfig creates a new ContextConfig and calls ma_context_config_init.
func NewContextConfig() *ContextConfig {
	config := C.ma_context_config_init()
	return &ContextConfig{config: &config}
}

func (cc *ContextConfig) Close() {
	if cc.config == nil {
		return
	}
	cc.config = nil
}

// PLog returns the log pointer.
func (cc *ContextConfig) PLog() *Log {
	if cc == nil || cc.config == nil {
		return nil
	}
	return &Log{log: cc.config.pLog}
}

// SetPLog sets the log pointer.
func (cc *ContextConfig) SetPLog(log *Log) {
	if cc != nil && cc.config != nil {
		cc.config.pLog = log.log
	}
}

// ThreadPriority returns the thread priority.
func (cc *ContextConfig) ThreadPriority() C.ma_thread_priority {
	if cc == nil || cc.config == nil {
		return 0
	}
	return cc.config.threadPriority
}

// SetThreadPriority sets the thread priority.
func (cc *ContextConfig) SetThreadPriority(priority C.ma_thread_priority) {
	if cc != nil && cc.config != nil {
		cc.config.threadPriority = priority
	}
}

// ThreadStackSize returns the thread stack size.
func (cc *ContextConfig) ThreadStackSize() uint64 {
	if cc == nil || cc.config == nil {
		return 0
	}
	return uint64(cc.config.threadStackSize)
}

// SetThreadStackSize sets the thread stack size.
func (cc *ContextConfig) SetThreadStackSize(size uint64) {
	if cc != nil && cc.config != nil {
		cc.config.threadStackSize = C.size_t(size)
	}
}

// PUserData returns the user data pointer.
func (cc *ContextConfig) PUserData() unsafe.Pointer {
	if cc == nil || cc.config == nil {
		return nil
	}
	return unsafe.Pointer(cc.config.pUserData)
}

// SetPUserData sets the user data pointer.
func (cc *ContextConfig) SetPUserData(userData unsafe.Pointer) {
	if cc != nil && cc.config != nil {
		cc.config.pUserData = userData
	}
}

// DSoundHWND returns the Windows window handle for DirectSound.
func (cc *ContextConfig) DSoundHWND() uintptr {
	if cc == nil || cc.config == nil {
		return 0
	}
	return uintptr(cc.config.dsound.hWnd)
}

// SetDSoundHWND sets the Windows window handle for DirectSound.
func (cc *ContextConfig) SetDSoundHWND(hWnd uintptr) {
	if cc != nil && cc.config != nil {
		cc.config.dsound.hWnd = C.ma_handle(hWnd)
	}
}

// AlsaUseVerboseDeviceEnumeration returns whether verbose device enumeration is enabled for ALSA.
func (cc *ContextConfig) AlsaUseVerboseDeviceEnumeration() bool {
	if cc == nil || cc.config == nil {
		return false
	}
	return cc.config.alsa.useVerboseDeviceEnumeration != 0
}

// SetAlsaUseVerboseDeviceEnumeration sets whether verbose device enumeration is enabled for ALSA.
func (cc *ContextConfig) SetAlsaUseVerboseDeviceEnumeration(enabled bool) {
	if cc != nil && cc.config != nil {
		cc.config.alsa.useVerboseDeviceEnumeration = C.ma_bool32(boolToCInt(enabled))
	}
}

// PulseApplicationName returns the PulseAudio application name.
func (cc *ContextConfig) PulseApplicationName() string {
	if cc == nil || cc.config == nil {
		return ""
	}
	if cc.config.pulse.pApplicationName == nil {
		return ""
	}
	return C.GoString(cc.config.pulse.pApplicationName)
}

// SetPulseApplicationName sets the PulseAudio application name.
func (cc *ContextConfig) SetPulseApplicationName(name string) {
	if cc != nil && cc.config != nil {
		cStr := C.CString(name)
		defer C.free(unsafe.Pointer(cStr))
		cc.config.pulse.pApplicationName = cStr
	}
}

// PulseServerName returns the PulseAudio server name.
func (cc *ContextConfig) PulseServerName() string {
	if cc == nil || cc.config == nil {
		return ""
	}
	if cc.config.pulse.pServerName == nil {
		return ""
	}
	return C.GoString(cc.config.pulse.pServerName)
}

// SetPulseServerName sets the PulseAudio server name.
func (cc *ContextConfig) SetPulseServerName(name string) {
	if cc != nil && cc.config != nil {
		cStr := C.CString(name)
		defer C.free(unsafe.Pointer(cStr))
		cc.config.pulse.pServerName = cStr
	}
}

// PulseTryAutoSpawn returns whether auto-spawning is enabled for PulseAudio.
func (cc *ContextConfig) PulseTryAutoSpawn() bool {
	if cc == nil || cc.config == nil {
		return false
	}
	return cc.config.pulse.tryAutoSpawn != 0
}

// SetPulseTryAutoSpawn sets whether auto-spawning is enabled for PulseAudio.
func (cc *ContextConfig) SetPulseTryAutoSpawn(enabled bool) {
	if cc != nil && cc.config != nil {
		cc.config.pulse.tryAutoSpawn = C.ma_bool32(boolToCInt(enabled))
	}
}

// CoreAudioSessionCategory returns the CoreAudio session category.
func (cc *ContextConfig) CoreAudioSessionCategory() C.ma_ios_session_category {
	if cc == nil || cc.config == nil {
		return 0
	}
	return cc.config.coreaudio.sessionCategory
}

// SetCoreAudioSessionCategory sets the CoreAudio session category.
func (cc *ContextConfig) SetCoreAudioSessionCategory(category C.ma_ios_session_category) {
	if cc != nil && cc.config != nil {
		cc.config.coreaudio.sessionCategory = category
	}
}

// CoreAudioSessionCategoryOptions returns the CoreAudio session category options.
func (cc *ContextConfig) CoreAudioSessionCategoryOptions() uint32 {
	if cc == nil || cc.config == nil {
		return 0
	}
	return uint32(cc.config.coreaudio.sessionCategoryOptions)
}

// SetCoreAudioSessionCategoryOptions sets the CoreAudio session category options.
func (cc *ContextConfig) SetCoreAudioSessionCategoryOptions(options uint32) {
	if cc != nil && cc.config != nil {
		cc.config.coreaudio.sessionCategoryOptions = C.ma_uint32(options)
	}
}

// CoreAudioNoAudioSessionActivate returns whether audio session activation is disabled for iOS.
func (cc *ContextConfig) CoreAudioNoAudioSessionActivate() bool {
	if cc == nil || cc.config == nil {
		return false
	}
	return cc.config.coreaudio.noAudioSessionActivate != 0
}

// SetCoreAudioNoAudioSessionActivate sets whether audio session activation is disabled for iOS.
func (cc *ContextConfig) SetCoreAudioNoAudioSessionActivate(disabled bool) {
	if cc != nil && cc.config != nil {
		cc.config.coreaudio.noAudioSessionActivate = C.ma_bool32(boolToCInt(disabled))
	}
}

// CoreAudioNoAudioSessionDeactivate returns whether audio session deactivation is disabled for iOS.
func (cc *ContextConfig) CoreAudioNoAudioSessionDeactivate() bool {
	if cc == nil || cc.config == nil {
		return false
	}
	return cc.config.coreaudio.noAudioSessionDeactivate != 0
}

// SetCoreAudioNoAudioSessionDeactivate sets whether audio session deactivation is disabled for iOS.
func (cc *ContextConfig) SetCoreAudioNoAudioSessionDeactivate(disabled bool) {
	if cc != nil && cc.config != nil {
		cc.config.coreaudio.noAudioSessionDeactivate = C.ma_bool32(boolToCInt(disabled))
	}
}

// JackClientName returns the JACK client name.
func (cc *ContextConfig) JackClientName() string {
	if cc == nil || cc.config == nil {
		return ""
	}
	if cc.config.jack.pClientName == nil {
		return ""
	}
	return C.GoString(cc.config.jack.pClientName)
}

// SetJackClientName sets the JACK client name.
func (cc *ContextConfig) SetJackClientName(name string) {
	if cc != nil && cc.config != nil {
		cStr := C.CString(name)
		defer C.free(unsafe.Pointer(cStr))
		cc.config.jack.pClientName = cStr
	}
}

// JackTryStartServer returns whether auto-starting the server is enabled for JACK.
func (cc *ContextConfig) JackTryStartServer() bool {
	if cc == nil || cc.config == nil {
		return false
	}
	return cc.config.jack.tryStartServer != 0
}

// SetJackTryStartServer sets whether auto-starting the server is enabled for JACK.
func (cc *ContextConfig) SetJackTryStartServer(enabled bool) {
	if cc != nil && cc.config != nil {
		cc.config.jack.tryStartServer = C.ma_bool32(boolToCInt(enabled))
	}
}

// Context wraps ma_context.
type Context struct {
	context *C.ma_context
}

// NewContext creates a new Context instance, initializes it with ma_context_init, and returns it.
// The backends parameter specifies which backends to use. If nil, miniaudio's default priorities are used.
// The config parameter can be nil for defaults.
func NewContext(backends []C.ma_backend, config *ContextConfig) (*Context, error) {
	var cBackends *C.ma_backend
	var backendCount C.ma_uint32
	if len(backends) > 0 {
		cBackends = &backends[0]
		backendCount = C.ma_uint32(len(backends))
	}

	var cConfig *C.ma_context_config
	if config != nil {
		cConfig = config.config
	}

	c := (*C.ma_context)(C.malloc(C.size_t(unsafe.Sizeof(*(*C.ma_context)(nil)))))
	if c == nil {
		return nil, ErrResult(MA_NO_MEMORY)
	}
	C.memset(unsafe.Pointer(c), 0, C.size_t(unsafe.Sizeof(*c)))

	result := C.ma_context_init(cBackends, backendCount, cConfig, c)
	if result != 0 {
		C.free(unsafe.Pointer(c))
		return nil, ErrResult(result)
	}

	return &Context{context: c}, nil
}

// Close uninitializes and releases the context memory.
func (c *Context) Close() {
	if c.context == nil {
		return
	}
	C.ma_context_uninit(c.context)
	C.free(unsafe.Pointer(c.context))
	c.context = nil
}

// Devices retrieves the devices from the context.
// It returns playback infos pointer, playback count, capture infos pointer, and capture count.
// Do not free the returned buffers as their memory is managed internally by miniaudio.
func (c *Context) Devices() (playbackInfo unsafe.Pointer, playbackCount uint32, captureInfo unsafe.Pointer, captureCount uint32, err error) {
	if c.context == nil {
		return nil, 0, nil, 0, ErrNilEngine
	}

	var pPlaybackInfos *C.ma_device_info
	var pCaptureInfos *C.ma_device_info
	var playbackCountC C.ma_uint32
	var captureCountC C.ma_uint32

	result := C.ma_context_get_devices(c.context, &pPlaybackInfos, &playbackCountC, &pCaptureInfos, &captureCountC)
	if result != 0 {
		return nil, 0, nil, 0, ErrResult(result)
	}

	playbackCount = uint32(playbackCountC)
	captureCount = uint32(captureCountC)

	if playbackCount > 0 {
		playbackInfo = unsafe.Pointer(pPlaybackInfos)
	}
	if captureCount > 0 {
		captureInfo = unsafe.Pointer(pCaptureInfos)
	}

	return playbackInfo, playbackCount, captureInfo, captureCount, nil
}

// Backend returns the backend of the context.
func (c *Context) Backend() C.ma_backend {
	if c.context == nil {
		return 0
	}
	return c.context.backend
}
