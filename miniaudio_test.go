package miniaudio

import (
	"testing"
	"time"
)

// ファイルから再生するテスト
func TestEngineInitAndPlaySoundFromFile(t *testing.T) {
	config := NewEngineConfig()
	// Create a new engine
	engine, err := NewEngine(config)
	if err != nil {
		t.Fatalf("Failed to create engine: %v", err)
	}
	defer engine.Close()
	snd, err := engine.PlaySoundFromFile("examples/sample.wav", 0, nil, nil)
	if err != nil {
		t.Fatalf("Failed to create sound: %v", err)
	}
	defer snd.Close()
	time.Sleep(1 * time.Second)
	// Verify sample rate and channels (should be 0 in no device mode)
	sampleRate := engine.SampleRate()
	channels := engine.Channels()
	t.Logf("Engine sample rate: %d, channels: %d", sampleRate, channels)
}

// Engineの基本的な初期化テスト
func TestEngineNewConfig(t *testing.T) {
	config := NewEngineConfig()
	if config == nil {
		t.Fatal("Expected non-nil config")
	}
	// Configが正しく初期化されているか確認
	t.Log("EngineConfig created successfully")
}

// EngineのStart/Stopテスト
func TestEngineStartStop(t *testing.T) {
	config := NewEngineConfig()
	engine, err := NewEngine(config)
	if err != nil {
		t.Fatalf("Failed to create engine: %v", err)
	}
	defer engine.Close()

	err = engine.Start()
	if err != nil {
		t.Fatalf("Failed to start engine: %v", err)
	}

	err = engine.Stop()
	if err != nil {
		t.Fatalf("Failed to stop engine: %v", err)
	}
}

// EngineのPlaySoundテスト
func TestEnginePlaySound(t *testing.T) {
	config := NewEngineConfig()
	engine, err := NewEngine(config)
	if err != nil {
		t.Fatalf("Failed to create engine: %v", err)
	}
	defer engine.Close()

	err = engine.PlaySound("examples/sample.wav")
	if err != nil {
		t.Fatalf("Failed to play sound: %v", err)
	}

	time.Sleep(1 * time.Second)
}

// EngineのGetSampleRateとGetChannelsテスト
func TestEngineGetSampleRateAndChannels(t *testing.T) {
	config := NewEngineConfig()
	engine, err := NewEngine(config)
	if err != nil {
		t.Fatalf("Failed to create engine: %v", err)
	}
	defer engine.Close()

	sampleRate := engine.SampleRate()
	channels := engine.Channels()
	t.Logf("Sample rate: %d, Channels: %d", sampleRate, channels)

	if engine.engine == nil {
		if sampleRate != 0 || channels != 0 {
			t.Errorf("Expected 0 for nil engine, got rate: %d, channels: %d", sampleRate, channels)
		}
	}
}

// EngineのSetVolumeテスト
func TestEngineSetVolume(t *testing.T) {
	config := NewEngineConfig()
	engine, err := NewEngine(config)
	if err != nil {
		t.Fatalf("Failed to create engine: %v", err)
	}
	defer engine.Close()

	err = engine.SetVolume(0.5)
	if err != nil {
		t.Fatalf("Failed to set volume: %v", err)
	}

	err = engine.SetVolume(1.0)
	if err != nil {
		t.Fatalf("Failed to set volume to 1.0: %v", err)
	}
}

// EngineのSetGainDBテスト
func TestEngineSetGainDB(t *testing.T) {
	config := NewEngineConfig()
	engine, err := NewEngine(config)
	if err != nil {
		t.Fatalf("Failed to create engine: %v", err)
	}
	defer engine.Close()

	err = engine.SetGainDB(0.0)
	if err != nil {
		t.Fatalf("Failed to set gain: %v", err)
	}

	err = engine.SetGainDB(-10.0)
	if err != nil {
		t.Fatalf("Failed to set negative gain: %v", err)
	}
}

// EngineのSetTimeInPCMFramesとGetTimeInPCMFramesテスト
func TestEngineTimeFunctions(t *testing.T) {
	config := NewEngineConfig()
	engine, err := NewEngine(config)
	if err != nil {
		t.Fatalf("Failed to create engine: %v", err)
	}
	defer engine.Close()

	time1 := engine.TimeInPCMFrames()
	t.Logf("Initial time: %d", time1)

	err = engine.SetTimeInPCMFrames(100)
	if err != nil {
		t.Fatalf("Failed to set time: %v", err)
	}

	time2 := engine.TimeInPCMFrames()
	t.Logf("After set time: %d", time2)

	if time2 != 100 {
		t.Errorf("Expected time 100, got %d", time2)
	}
}

// EngineのListenerSetPositionテスト
func TestEngineListenerPosition(t *testing.T) {
	config := NewEngineConfig()
	engine, err := NewEngine(config)
	if err != nil {
		t.Fatalf("Failed to create engine: %v", err)
	}
	defer engine.Close()

	err = engine.ListenerSetPosition(0, 0, 0, 0)
	if err != nil {
		t.Fatalf("Failed to set listener position: %v", err)
	}

	err = engine.ListenerSetPosition(0, 1.0, 2.0, 3.0)
	if err != nil {
		t.Fatalf("Failed to set listener position: %v", err)
	}

	x, y, z := engine.ListenerPosition(0)
	t.Logf("Listener position: (%f, %f, %f)", x, y, z)

	if x != 1.0 || y != 2.0 || z != 3.0 {
		t.Errorf("Expected (1, 2, 3), got (%f, %f, %f)", x, y, z)
	}
}

// EngineのListenerSetDirectionテスト
func TestEngineListenerDirection(t *testing.T) {
	config := NewEngineConfig()
	engine, err := NewEngine(config)
	if err != nil {
		t.Fatalf("Failed to create engine: %v", err)
	}
	defer engine.Close()

	err = engine.ListenerSetDirection(0, 0, 0, 1)
	if err != nil {
		t.Fatalf("Failed to set listener direction: %v", err)
	}

	x, y, z := engine.ListenerDirection(0)
	t.Logf("Listener direction: (%f, %f, %f)", x, y, z)

	if x != 0 || y != 0 || z != 1 {
		t.Errorf("Expected (0, 0, 1), got (%f, %f, %f)", x, y, z)
	}
}

// EngineのListenerSetVelocityテスト（ドップラー効果用）
func TestEngineListenerVelocity(t *testing.T) {
	config := NewEngineConfig()
	engine, err := NewEngine(config)
	if err != nil {
		t.Fatalf("Failed to create engine: %v", err)
	}
	defer engine.Close()

	err = engine.ListenerSetVelocity(0, 1.0, 0, 0)
	if err != nil {
		t.Fatalf("Failed to set listener velocity: %v", err)
	}

	x, y, z := engine.ListenerVelocity(0)
	t.Logf("Listener velocity: (%f, %f, %f)", x, y, z)

	if x != 1.0 || y != 0 || z != 0 {
		t.Errorf("Expected (1, 0, 0), got (%f, %f, %f)", x, y, z)
	}
}

// EngineのListenerSetWorldUpテスト
func TestEngineListenerWorldUp(t *testing.T) {
	config := NewEngineConfig()
	engine, err := NewEngine(config)
	if err != nil {
		t.Fatalf("Failed to create engine: %v", err)
	}
	defer engine.Close()

	err = engine.ListenerSetWorldUp(0, 0, 1, 0)
	if err != nil {
		t.Fatalf("Failed to set world up: %v", err)
	}
}

// EngineのListenerSetEnabled/IsEnabledテスト
func TestEngineListenerEnabled(t *testing.T) {
	config := NewEngineConfig()
	engine, err := NewEngine(config)
	if err != nil {
		t.Fatalf("Failed to create engine: %v", err)
	}
	defer engine.Close()

	enabled := engine.ListenerIsEnabled(0)
	t.Logf("Listener enabled (default): %v", enabled)

	err = engine.ListenerSetEnabled(0, true)
	if err != nil {
		t.Fatalf("Failed to enable listener: %v", err)
	}

	if !engine.ListenerIsEnabled(0) {
		t.Error("Expected listener to be enabled")
	}

	err = engine.ListenerSetEnabled(0, false)
	if err != nil {
		t.Fatalf("Failed to disable listener: %v", err)
	}

	if engine.ListenerIsEnabled(0) {
		t.Error("Expected listener to be disabled")
	}
}

// Engineのnilポインタテスト
func TestEngineNilPointer(t *testing.T) {
	// Goのメソッドはnilレシーバーでも呼び出されるため、
	// パニックが発生する場合はrecoverでキャッチする
	defer func() {
		if r := recover(); r != nil {
			t.Log("Nil pointer dereference caught (expected behavior in Go)")
		}
	}()

	var engine *Engine = nil
	_ = engine.SampleRate()
	// This will panic because Go methods receive the receiver first,
	// and the nil check happens inside the method.
	// The panic is expected, but we catch it above.
	t.Error("Should have panicked before here")
}

// Engineのnilポインタテスト（安全なバージョン）
func TestEngineNilPointerSafe(t *testing.T) {
	engine := &Engine{engine: nil}

	if engine.SampleRate() != 0 {
		t.Errorf("Expected 0 for nil engine GetSampleRate")
	}
	if engine.Channels() != 0 {
		t.Errorf("Expected 0 for nil engine GetChannels")
	}
	if engine.TimeInPCMFrames() != 0 {
		t.Errorf("Expected 0 for nil engine GetTimeInPCMFrames")
	}

	x, y, z := engine.ListenerPosition(0)
	if x != 0 || y != 0 || z != 0 {
		t.Errorf("Expected (0, 0, 0) for nil engine ListenerGetPosition")
	}

	err := engine.Start()
	if err == nil {
		t.Error("Expected error for nil engine Start")
	} else {
		t.Logf("Got expected error: %v", err)
	}

	x, y, z = engine.ListenerDirection(0)
	if x != 0 || y != 0 || z != 0 {
		t.Errorf("Expected (0, 0, 0) for nil engine ListenerGetDirection")
	}

	x, y, z = engine.ListenerVelocity(0)
	if x != 0 || y != 0 || z != 0 {
		t.Errorf("Expected (0, 0, 0) for nil engine ListenerGetVelocity")
	}
}

// EngineのClose複数呼び出しテスト（パニックしないことを確認）
func TestEngineMultipleClose(t *testing.T) {
	config := NewEngineConfig()
	engine, err := NewEngine(config)
	if err != nil {
		t.Fatalf("Failed to create engine: %v", err)
	}

	engine.Close()
	// Should not panic
	engine.Close()
}
