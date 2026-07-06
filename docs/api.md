# miniaudio API リファレンス

miniaudio v0.11.25 の主要な型をカテゴリ別に整理した簡易ドキュメントです。

---

## 1. ベース型 (Basic Types)

### プリミティブ型

| 型 | 定義 | 説明 |
|----|------|------|
| `ma_bool8` | `ma_uint8` | 8ビット boolean |
| `ma_bool32` | `ma_uint32` | 32ビット boolean |
| `ma_float` | `float` | 32ビット浮動小数点 |
| `ma_double` | `double` | 64ビット浮動小数点 |
| `ma_proc` | `void*` | 汎用関数ポインタ |
| `ma_handle` | `void*` | ハンドル |
| `ma_ptr` | `void*` | 汎用ポインタ |

### ウィンドウズワイド文字

| 型 | 定義 | 説明 |
|----|------|------|
| `ma_wchar_win32` | `ma_uint16` | 16ビットワイド文字 (UCS-2/UTF-16LE) |

---

## 2. 列挙型 (Enums)

### 結果コード

| 型 | 説明 |
|----|------|
| `ma_result` | 操作の結果を示す列挙 (`MA_SUCCESS`, `MA_ERROR`, `MA_INVALID_ARGS`, `MA_END_OF_FILE`, `MA_OUT_OF_MEMORY`, `MA_UNSUPPORTED` など) |

### オーディオフォーマット

| 型 | 説明 |
|----|------|
| `ma_format` | サンプルフォーマット (`ma_format_unknown`, `ma_format_f32`, `ma_format_s16`, `ma_format_s24`, `ma_format_u8`) |
| `ma_channel` | `ma_uint8` - チャネルインデックス |
| `ma_channel_mix_mode` | チャネルミキシングモード (`ma_channel_mix_mode_default`, `ma_channel_mix_mode_simple`) |
| `ma_dither_mode` | ディザリングモード (`ma_dither_mode_none`, `ma_dither_mode_rectangular`, ...) |
| `ma_standard_channel_map` | 標準チャネルマップ |

### デバイス・バックエンド

| 型 | 説明 |
|----|------|
| `ma_device_type` | デバイスタイプ (`ma_device_type_none`, `ma_device_type_playback`, `ma_device_type_capture`, `ma_device_type_duplex`, `ma_device_type_loopback`) |
| `ma_backend` | バックエンド (`ma_backend_default`, `ma_backend_wasapi`, `ma_backend_wasapi`, `ma_backend_directsound`, `ma_backend_winmm`, `ma_backend_alsa`, `ma_backend_pulseaudio`, `ma_backend_coreaudio`, `ma_backend_jack`, `ma_backend_aaudio`, `ma_backend_opensl`, `ma_backend_sndio`, `ma_backend_audio4`, `ma_backend_oss`, `ma_backend_nullapi`) |
| `ma_device_state` | デバイスの状態 (`ma_device_state_uninitialized`, `ma_device_state_init_failed`, `ma_device_state_started`, `ma_device_state_stopped`) |
| `ma_share_mode` | シャアモード (`ma_share_mode_auto`, `ma_share_mode_shared`, `ma_share_mode_exclusive`) |
| `ma_performance_profile` | パフォーマンスプロファイル (`ma_performance_profile_undefined`, `ma_performance_profile_latency`, `ma_performance_profile_throughput`) |

### 設定・モード

| 型 | 説明 |
|----|------|
| `ma_device_notification_type` | デバイス通知タイプ (`ma_device_notification_type_pre_init`, `ma_device_notification_type_post_init`, `ma_device_notification_type_unloaded`) |
| `ma_log_level` | ログレベル (`ma_log_level_trace`, `ma_log_level_debug`, `ma_log_level_info`, `ma_log_level_warn`, `ma_log_level_error`, `ma_log_level_fatal`) |
| `ma_resample_algorithm` | リサンプリングアルゴリズム (`ma_resample_algorithm_linear`, `ma_resample_algorithm_cubic`, `ma_resample_algorithm_sinc`, `ma_resample_algorithm_custom`) |
| `ma_attenuation_model` | 減衰モデル (`ma_attenuation_model_none`, `ma_attenuation_model_inverse`, `ma_attenuation_model_linear`, `ma_attenuation_model_exponential`) |
| `ma_positioning` | 位置決めモード (`ma_positioning_relative`, `ma_positioning_absolute`) |
| `ma_handedness` |  handedness (`ma_handedness_left_handed`, `ma_handedness_right_handed`) |
| `ma_pan_mode` | パンモード |
| `ma_mono_expansion_mode` | モノ展開モード |
| `ma_encoding_format` | エンコーディングフォーマット (`ma_encoding_format_unknown`, `ma_encoding_format_wav`, `ma_encoding_format_ogg`, `ma_encoding_format_mp3`, ...) |

### iOS/AudioToolbox 固有

| 型 | 説明 |
|----|------|
| `ma_ios_session_category` | iOSセッションカテゴリ |
| `ma_ios_session_category_option` | iOSセッションカテゴリオプション |
| `ma_opensl_stream_type` | OpenSL Streamタイプ |
| `ma_opensl_recording_preset` | OpenSL録音プリセット |

### WASAPI/AAudio 固有

| 型 | 説明 |
|----|------|
| `ma_wasapi_usage` | WASAPI使用モード (`ma_wasapi_usage_multimedia`, `ma_wasapi_usage_games`, `ma_wasapi_usage_native`) |
| `ma_aaudio_usage` | AAudio使用用途 |
| `ma_aaudio_content_type` | AAudioコンテンツタイプ |
| `ma_aaudio_input_preset` | AAudio入力プリセット |
| `ma_aaudio_allowed_capture_policy` | AAudio録音ポリシー |

### データソース

| 型 | 説明 |
|----|------|
| `ma_stream_format` | ストリームフォーマット |
| `ma_stream_layout` | ストリーmlayout |
| `ma_standard_sample_rate` | 標準サンプルレート (8000, 16000, 44100, 48000, 96000, ...) |
| `ma_waveform_type` | 波形タイプ (`ma_waveform_type_sine`, `ma_waveform_type_square`, `ma_waveform_type_triangle`, `ma_waveform_type_sawtooth`) |
| `ma_noise_type` | ノイズタイプ (`ma_noise_type_white`, `ma_noise_type_pink`, `ma_noise_type_brownian`) |

### リソースマネージャー

| 型 | 説明 |
|----|------|
| `ma_resource_manager_data_source_flags` | データソースフラグ |
| `ma_resource_manager_pipeline_notifications` | パイプライン通知タイプ |
| `ma_resource_manager_flags` | リソースマネージャーフラグ |
| `ma_resource_manager_data_supply_type` | データ供給タイプ |

### ノード

| 型 | 説明 |
|----|------|
| `ma_node_flags` | ノードフラグ |
| `ma_node_state` | ノード状態 (`ma_node_state_uninitialised`, `ma_node_state_initialising`, `ma_node_state_stopped`, `ma_node_state_started`) |

### ファイル/I/O

| 型 | 説明 |
|----|------|
| `ma_open_mode_flags` | ファイルオープンモードフラグ |
| `ma_seek_origin` | シーク元 (`ma_seek_origin_start`, `ma_seek_origin_current`, `ma_seek_origin_end`) |

### ジョブキュー

| 型 | 説明 |
|----|------|
| `ma_job_type` | ジョブタイプ |
| `ma_job_queue_flags` | ジョブキューフラグ |

---

## 3. 設定構造体 (Config Structures)

### Engine Config

```c
typedef struct {
    ma_resource_manager* pResourceManager;          // カスタムリソースマネージャー (NULL=自動生成)
    ma_context* pContext;                           // コンテキスト
    ma_device* pDevice;                             // デバイス
    ma_device_id* pPlaybackDeviceID;                // プレイバックデバイスID
    ma_device_data_proc dataCallback;               // カスタムデータコールバック
    ma_device_notification_proc notificationCallback;// 通知コールバック
    ma_log* pLog;                                   // ログ (NULL=コンテキストのログを使用)
    ma_uint32 listenerCount;                        // リスナー数 (1〜MA_ENGINE_MAX_LISTENERS)
    ma_uint32 channels;                             // チャネル数 (0=デバイスのネイティブ)
    ma_uint32 sampleRate;                           // サンプルレート (0=デバイスのネイティブ)
    ma_uint32 periodSizeInFrames;                   // フレーム単位の期間サイズ
    ma_uint32 periodSizeInMilliseconds;             // ミリ秒単位の期間サイズ
    ma_uint32 gainSmoothTimeInFrames;               // ゲイン平滑化時間 (フレーム)
    ma_uint32 gainSmoothTimeInMilliseconds;         // ゲイン平滑化時間 (ミリ秒)
    ma_uint32 defaultVolumeSmoothTimeInPCMFrames;   // デフォルト音量平滑化時間
    ma_uint32 preMixStackSizeInBytes;               // プリミックスタックサイズ
    ma_allocation_callbacks allocationCallbacks;    // カスタムアロケーション
    ma_bool32 noAutoStart;                          // 自動スタート無効化
    ma_bool32 noDevice;                             // デバイス作成無効化
    ma_mono_expansion_mode monoExpansionMode;       // モノ展開モード
    ma_vfs* pResourceManagerVFS;                    // リソースマネージャー用VFS
    ma_engine_process_proc onProcess;               // プリミックスターフni発火
    void* pProcessUserData;                         // onProcessのユーザーデータ
    ma_resampler_config resourceManagerResampling;  // リソースマネージャーのリサンプリング設定
    ma_resampler_config pitchResampling;            // ピッチ/ドップラー用リサンプリング設定
} ma_engine_config;
```

### Context Config

```c
struct ma_context_config {
    ma_log* pLog;                                   // ログオブジェクト
    void* pUserData;                                // ユーザーデータ
    ma_allocation_callbacks allocationCallbacks;    // カスタムアロケーション
    // バックエンド固有設定
    struct { ma_bool32 enumerateSoftwar; } wasapi;
    struct { /* ... */ } directsound;
    struct { /* ... */ } audioio;
    struct { /* ... */ } portaudio;
    struct { /* ... */ } alsa;
    struct { /* ... */ } pulseaudio;
    struct { /* ... */ } coreaudio;
    struct { /* ... */ } jack;
    struct { /* ... */ } oss;
};
```

### Device Config

```c
struct ma_device_config {
    ma_device_type deviceType;                      // デバイスタイプ
    // プレイバック設定
    struct {
        ma_format format;                           // フォーマット
        ma_uint32 channels;                         // チャネル数
        ma_uint32 internalChannelCount;             // 内部チャネル数
        ma_format internalFormat;                   // 内部フォーマット
        const void* pDeviceID;                      // デバイスID
        ma_share_mode shareMode;                    // シャアモード
        ma_wasapi_usage wasapiUsage;                // WASAPI使用モード (WASAPIのみ)
        ma_bool32 noAutoConvertDevices;             // 自動デバイス変換無効化
    } playback;
    // キャプチャ設定
    struct {
        ma_format format;                           // フォーマット
        ma_uint32 channels;                         // チャネル数
        ma_uint32 internalChannelCount;             // 内部チャネル数
        ma_format internalFormat;                   // 内部フォーマット
        const void* pDeviceID;                      // デバイスID
        ma_share_mode shareMode;                    // シャアモード
        ma_opensl_stream_type openslStreamType;     // OpenSL Streamタイプ (Android)
        ma_opensl_recording_preset openslRecordingPreset; // OpenSL録音プリセット (Android)
        ma_bool32 noAutoConvertDevices;             // 自動デバイス変換無効化
    } capture;
    ma_uint32 sampleRate;                           // サンプルレート
    ma_device_data_proc dataCallback;               // データコールバック
    ma_device_notification_proc notificationCallback;// 通知コールバック
    void* pUserData;                                // ユーザーデータ
    ma_uint32 periodSizeInFrames;                   // フレーム単位の期間サイズ
    ma_uint32 periodSizeInMilliseconds;             // ミリ秒単位の期間サイズ
    ma_processing_mode processingMode;              // 処理モード
    ma_bool32 noFlags;                              // フラグ (0=デフォルト)
    ma_bool32 noDeviceAutomaticStartStop;           // 自動スタート/ストップ無効化
    ma_sample_rate_conversion_quality src;          // SRC品質
    ma_uint32 defaultHandlerIndex;                  // デフォルトハンドラインデックス
};
```

### Decoder Config

```c
typedef struct {
    ma_format format;                           // 出力フォーマット (0=ストリームの内部フォーマット)
    ma_uint32 channels;                         // チャネル数 (0=内部チャネル数を使用)
    ma_uint32 sampleRate;                       // サンプルレート (0=内部サンプルレートを使用)
    ma_channel* pChannelMap;                    // チャネルマップ
    ma_channel_mix_mode channelMixMode;         // チャネルミキシングモード
    ma_dither_mode ditherMode;                  // ディザリングモード
    ma_resampler_config resampling;             // リサンプリング設定
    ma_allocation_callbacks allocationCallbacks;// カスタムアロケーション
    ma_encoding_format encodingFormat;          // エンコーディングフォーマット (0=自動検出)
    ma_uint32 seekPointCount;                   // シークポイント数
    ma_decoding_backend_vtable** ppCustomBackendVTables;  // カスタムバックエンドvtable
    ma_uint32 customBackendCount;               // カスタムバックエンド数
    void* pCustomBackendUserData;               // カスタムバックエンドユーザーデータ
} ma_decoder_config;
```

### Resource Manager Config

```c
typedef struct {
    ma_vfs* pVFS;                                   // ファイルシステム
    ma_uint32 maxLoadedSounds;                      // 最大ロード済みサウンド数
    ma_uint32 maxStreamingSounds;                   // 最大ストリーミングサウンド数
    ma_uint32 supplyThreadCount;                    // サプライスレッド数
    ma_bool32 noSupplyThread;                       // サプライスレッド無効化
    ma_allocation_callbacks allocationCallbacks;    // カスタムアロケーション
    ma_log* pLog;                                   // ログ
    ma_resource_manager_pipeline_notifications pipelineNotifications; // パイプライン通知
    void* pUserData;                                // ユーザーデータ
} ma_resource_manager_config;
```

### Node Graph Config

```c
typedef struct {
    ma_uint32 preMixStackSizeInBytes;               // プリミックスタックサイズ (デフォルト: 1MB)
    ma_allocation_callbacks allocationCallbacks;    // カスタムアロケーション
} ma_node_graph_config;
```

### Sound Config

```c
typedef struct {
    ma_uint32 defaultVolumeSmoothTimeInPCMFrames;   // デフォルト音量平滑化時間
    ma_resampler_config pitchResampling;            // ピッチ用リサンプリング設定
} ma_sound_config;
```

---

## 4. メイン構造体 (Main Structures)

### Engine

```c
struct ma_engine {
    ma_node_graph nodeGraph;                        // ノードグラフ (最初のメンバー)
    ma_resource_manager* pResourceManager;          // リソースマネージャー
    ma_device* pDevice;                             // デバイス
    ma_log* pLog;                                   // ログ
    ma_uint32 sampleRate;                           // サンプルレート
    ma_uint32 listenerCount;                        // リスナー数
    ma_spatializer_listener listeners[MA_ENGINE_MAX_LISTENERS]; // リスナー
    ma_allocation_callbacks allocationCallbacks;    // カスタムアロケーション
    ma_bool8 ownsResourceManager;                   // リソースマネージャー所有権
    ma_bool8 ownsDevice;                            // デバイス所有権
    ma_spinlock inlinedSoundLock;                   // インラインサウンドロック
    ma_sound_inlined* pInlinedSoundHead;            // 最初のインラインサウンド
    MA_ATOMIC(4, ma_uint32) inlinedSoundCount;      // インラインサウンド数
    ma_uint32 gainSmoothTimeInFrames;               // ゲイン平滑化時間 (フレーム)
    ma_uint32 defaultVolumeSmoothTimeInPCMFrames;   // デフォルト音量平滑化時間
    ma_mono_expansion_mode monoExpansionMode;       // モノ展開モード
    ma_engine_process_proc onProcess;               // プリミックスターフni発火
    void* pProcessUserData;                         // onProcessのユーザーデータ
    ma_resampler_config pitchResamplingConfig;      // ピッチ用リサンプリング設定
};
```

### Sound

```c
struct ma_sound {
    ma_node_base node;                              // ノードベース (最初のメンバー)
    ma_resource_manager_data_source* pDataSource;   // データソース
    ma_uint32 channels;                             // チャネル数
    ma_uint32 sampleRate;                           // サンプルレート
    ma_format format;                               // フォーマット
    ma_bool8 loop;                                  // ループ有効
    ma_bool8 paused;                                // ポーズ中
    ma_bool8 manualLoopSetting;                     // 手動ループ設定
    float volume;                                   // 音量 (0.0〜1.0)
    float pitch;                                    // ピッチ (1.0=標準)
    float pan;                                      // パン (-1.0〜1.0)
    ma_uint64 startPCMFrame;                        // 開始位置 (PCMフレーム)
    ma_uint64 endPCMFrame;                          // 終了位置 (PCMフレーム)
    ma_uint32 gainSmoothTimeInFrames;               // ゲイン平滑化時間
    ma_uint32 defaultVolumeSmoothTimeInPCMFrames;   // デフォルト音量平滑化時間
    ma_resampler* pPitchResampler;                  // ピッチ用リサンプリャ
    void* pHeap;                                    // ヒープメモリ
};
```

### Context

```c
struct ma_context {
    ma_backend backends[8];                         // バックエンド配列
    ma_uint32 backendCount;                         // バックエンド数
    ma_context_command** ppCommands;                // コマンドキュー
    ma_uint32 commandCount;                         // コマンド数
    ma_uint32 commandCapacity;                      // コマンド容量
    ma_log* pLog;                                   // ログ
    void* pUserData;                                // ユーザーデータ
    ma_allocation_callbacks allocationCallbacks;    // カスタムアロケーション
    ma_backend_callbacks backendCallbacks[8];       // バックエンドコールバック
    ma_bool32 isRunning;                            // 実行中フラグ
    // バックエンド固有データ
    struct { /* wasapi */ } pBackendData;
};
```

### Device

```c
struct ma_device {
    ma_context* pContext;                           // コンテキスト
    ma_device_config config;                        // デバイス設定
    ma_device_state state;                          // 状態
    ma_uint32 sampleRate;                           // サンプルレート
    void* pUserData;                                // ユーザーデータ
    ma_device_data_proc dataCallback;               // データコールバック
    ma_device_notification_proc notificationCallback;// 通知コールバック
    ma_allocation_callbacks allocationCallbacks;    // カスタムアロケーション
    // バックエンド固有データ
    void* pBackendData;
    // ジョブスレッド
    ma_device_job_thread jobThread;
};
```

### Decoder

```c
struct ma_decoder {
    ma_data_source_base ds;                         // データソースベース (最初のメンバー)
    ma_decoder_config config;                       // デコーダー設定
    ma_uint64 cursor;                               // カーソル位置
    ma_uint64 outputFrameCount;                     // 出力フレーム数
    ma_bool32 atEnd;                                // 末尾フラグ
    void* pCustomBackendData;                       // カスタムバックエンドデータ
    ma_decoding_backend_vtable* pVTable;            // vtable
    ma_log* pLog;                                   // ログ
    ma_allocation_callbacks allocationCallbacks;    // カスタムアロケーション
};
```

### Resource Manager

```c
struct ma_resource_manager {
    ma_vfs* pVFS;                                   // ファイルシステム
    ma_uint32 maxLoadedSounds;                      // 最大ロード済みサウンド数
    ma_uint32 maxStreamingSounds;                   // 最大ストリーミングサウンド数
    ma_allocation_callbacks allocationCallbacks;    // カスタムアロケーション
    ma_log* pLog;                                   // ログ
    // データソース管理
    ma_resource_manager_data_source* pDataSourceHead;
    ma_resource_manager_data_source* pDataSourceTail;
    ma_uint32 dataSourceCount;
    // バッファ管理
    ma_resource_manager_data_buffer* pBufferHead;
    ma_resource_manager_data_buffer* pBufferTail;
    ma_uint32 bufferCount;
};
```

---

## 5. フィルター・エフェクト構造体 (Filters & Effects)

### Biquad Filter

```c
typedef struct {
    ma_float q;                                     // Qファクター
    ma_float gainLinear;                            // ゲイン (線形)
    ma_float frequency;                             // 周波数
} ma_biquad_config;

struct ma_biquad {
    ma_biquad_coefficient a[3];                     // a係数
    ma_biquad_coefficient b[2];                     // b係数
    ma_float x1a;                                   // ステートx1
    ma_float x1b;                                   // ステートx1 (B)
    ma_float y1a;                                   // ステートy1
    ma_float y1b;                                   // ステートy1 (B)
};
```

### Low-Pass Filter (LPF)

```c
typedef struct {
    ma_float frequency;                             // 遮断周波数
    ma_float resonance;                             // 共鳴
} ma_lpf_config;

struct ma_lpf {
    ma_float frequency;                             // 遮断周波数
    ma_float resonance;                             // 共鳴
    ma_float gain;                                  // ゲイン
    ma_float z1;                                    // ステートz1
};
```

### High-Pass Filter (HPF)

```c
typedef struct {
    ma_float frequency;                             // 遮断周波数
    ma_float resonance;                             // 共鳴
} ma_hpf_config;

struct ma_hpf {
    ma_float frequency;                             // 遮断周波数
    ma_float resonance;                             // 共鳴
    ma_float gain;                                  // ゲイン
    ma_float z1;                                    // ステートz1
};
```

### Band-Pass Filter (BPF)

```c
typedef struct {
    ma_float frequency;                             // 中心周波数
    ma_float resonance;                             // 共鳴
} ma_bpf_config;

struct ma_bpf {
    ma_float frequency;                             // 中心周波数
    ma_float resonance;                             // 共鳴
    ma_float gain;                                  // ゲイン
    ma_float z1;                                    // ステートz1
};
```

### Delay

```c
typedef struct {
    ma_uint32 maxDelayInFrames;                     // 最大遅延時間 (フレーム)
    ma_bool32 delayInFrames;                        // フレーム指定有効
} ma_delay_config;

struct ma_delay {
    ma_float* pBuffer;                              // バッファ
    ma_uint32 bufferLength;                         // バッファ長
    ma_uint32 writeIndex;                           // 書き込みインデックス
    ma_float gain;                                  // ゲイン
};
```

### Gainer

```c
typedef struct {
    ma_float gain;                                  // ゲイン
} ma_gainer_config;

struct ma_gainer {
    ma_float gain;                                  // ゲイン
};
```

### Panner

```c
typedef struct {
    ma_uint32 channelCount;                         // チャネル数
    ma_uint32 sampleRate;                           // サンプルレート
    ma_pan_mode mode;                               // パンモード
} ma_panner_config;

struct ma_panner {
    ma_float pan;                                   // パン値 (-1.0〜1.0)
};
```

### Fader

```c
typedef struct {
    ma_uint32 sampleRate;                           // サンプルレート
} ma_fader_config;

struct ma_fader {
    float currentGain;                              // 現在のゲイン
    float targetGain;                               // ターゲットゲイン
    ma_uint32 samplesUntilComplete;                 // 完了までのサンプル数
};
```

### Spatializer

```c
typedef struct {
    ma_uint32 listenerCount;                        // リスナー数
    ma_attenuation_model attenuationModel;          // 減衰モデル
    ma_float minGain;                               // 最小ゲイン
    ma_float maxGain;                               // 最大ゲイン
    ma_float rolloffFactor;                         // ロールオフファクター
} ma_spatializer_config;

struct ma_spatializer {
    ma_vector_3f position;                          // 位置
    ma_vector_3f forward;                           // 前方ベクトル
    ma_vector_3f upward;                            // 上ベクトル
};
```

### Spatializer Listener

```c
typedef struct {
    ma_uint32 listenerCount;                        // リスナー数
} ma_spatializer_listener_config;

struct ma_spatializer_listener {
    ma_vector_3f position;                          // 位置
    ma_vector_3f forward;                           // 前方ベクトル
    ma_vector_3f upward;                            // 上ベクトル
    ma_vector_3f velocity;                          // 速度
};
```

---

## 6. データコンバーター構造体 (Data Converters)

### Linear Resampler

```c
typedef struct {
    ma_uint32 lpFilterOrder;                        // LPF次数
} ma_linear_resampler_config;

struct ma_linear_resampler {
    ma_float* pInBuffer;                            // 入力バッファ
    ma_float* pOutBuffer;                           // 出力バッファ
    ma_uint32 inChannelCount;                       // 入力チャネル数
    ma_uint32 outChannelCount;                      // 出力チャネル数
    ma_float sampleRateIn;                          // 入力サンプルレート
    ma_float sampleRateOut;                         // 出力サンプルレート
};
```

### Resampler (High Quality)

```c
typedef struct ma_resampler_config {
    ma_format format;                               // フォーマット (f32またはs16)
    ma_uint32 channels;                             // チャネル数
    ma_uint32 sampleRateIn;                         // 入力サンプルレート
    ma_uint32 sampleRateOut;                        // 出力サンプルレート
    ma_resample_algorithm algorithm;                // アルゴリズム
    ma_resampling_backend_vtable* pBackendVTable;   // バックエンドvtable (custom用)
    void* pBackendUserData;                         // バックエンドユーザーデータ
    struct {
        ma_uint32 lpfOrder;                         // LPF次数 (linearアルゴリズム用)
    } linear;
} ma_resampler_config;

struct ma_resampler {
    ma_format format;                               // フォーマット
    ma_uint32 channels;                             // チャネル数
    ma_uint32 sampleRateIn;                         // 入力サンプルレート
    ma_uint32 sampleRateOut;                        // 出力サンプルレート
    void* pBackend;                                 // バックエンド
    ma_resample_algorithm algorithm;                // アルゴリズム
};
```

### Channel Converter

```c
typedef struct {
    ma_uint32 inChannelCount;                       // 入力チャネル数
    ma_uint32 outChannelCount;                      // 出力チャネル数
    const ma_channel* pInChannelMap;                // 入力チャネルマップ
    const ma_channel* pOutChannelMap;               // 出力チャネルマップ
} ma_channel_converter_config;

struct ma_channel_converter {
    ma_channel_conversion_path path;                // 変換パス
};
```

### Data Converter

```c
typedef struct {
    ma_format format;                               // フォーマット
    ma_uint32 inChannelCount;                       // 入力チャネル数
    ma_uint32 outChannelCount;                      // 出力チャネル数
    const ma_channel* pInChannelMap;                // 入力チャネルマップ
    const ma_channel* pOutChannelMap;               // 出力チャネルマップ
    ma_channel_mix_mode channelMixMode;             // チャネルミキシングモード
    ma_dither_mode ditherMode;                      // ディザリングモード
    ma_uint32 sampleRateIn;                         // 入力サンプルレート
    ma_uint32 sampleRateOut;                        // 出力サンプルレート
} ma_data_converter_config;

struct ma_data_converter {
    ma_channel_converter channelConverter;          // チャネルコンバーター
    ma_linear_resampler resampler;                  // リサンプリャ
};
```

---

## 7. データソース構造体 (Data Sources)

### Data Source Base

```c
typedef struct ma_data_source ma_data_source;       // 前方宣言

struct ma_data_source_vtable {
    const char* (* pName)(void);                    // 名前
    ma_result (* onGetFormat)(ma_data_source* pSource, ma_format* pFormat, ma_channel_count* pChannelCount, ma_uint32* pSampleRate, ma_channel_map* pChannelMap, ma_uint32 channelMapCap, size_t* pChannelMapSize);
    ma_result (* onRead)(ma_data_source* pSource, void* pBufferOut, ma_uint64 framesIn, ma_uint64* pFramesRead);
    ma_bool8 (* atEnd)(ma_data_source* pSource);
};

struct ma_data_source_base {
    ma_data_source_vtable vtable;                   // vtable (最初のメンバー)
    ma_data_source* pNext;                          // 次のデータソース
    void* pUserData;                                // ユーザーデータ
};
```

### Audio Buffer

```c
typedef struct {
    const void* pData;                             // データポインタ
    ma_uint32 dataSize;                            // データサイズ
} ma_audio_buffer_ref;

typedef struct {
    ma_format format;                              // フォーマット
    ma_uint32 channelCount;                        // チャネル数
    ma_uint32 sampleRate;                          // サンプルレート
    ma_bool8 loop;                                 // ループ有効
    ma_uint64 loopStartInclusivePCMFrame;          // ループ開始位置
    ma_uint64 loopEndExclusivePCMFram;             // ループ終了位置
} ma_audio_buffer_config;

struct ma_audio_buffer {
    ma_audio_buffer_ref buffer;                    // バッファ参照
};
```

### Paged Audio Buffer

```c
struct ma_paged_audio_buffer_page {
    ma_paged_audio_buffer_page* pNext;             // 次のページ
    ma_uint8* pData;                               // データ
    ma_uint32 dataSize;                            // データサイズ
};

typedef struct {
    ma_uint32 pageSizeInBytes;                     // ページサイズ
} ma_paged_audio_buffer_config;

struct ma_paged_audio_buffer {
    ma_paged_audio_buffer_data data;               // データ
};
```

### Ring Buffer

```c
struct ma_rb {                                    // 汎用リングバッファ
    void* pMemory;                                 // メモリ
    ma_uint32 capacityInBytes;                     // 容量 (バイト)
    ma_uint32 head;                                // ヘッドポインタ
    ma_uint32 tail;                                // テールポインタ
};

struct ma_pcm_rb {                                // PCM専用リングバッファ
    ma_rb rb;                                      // 汎用rb
    ma_format format;                              // フォーマット
    ma_uint32 channelCount;                        // チャネル数
};

struct duplex_rb {                                // デュプレックスリングバッファ
    ma_pcm_rb playback;                            // プレイバック
    ma_pcm_rb capture;                             // キャプチャ
};
```

---

## 8. コールバック型 (Callback Types)

### ログコールバック

```c
typedef void (*ma_log_callback_proc)(void* pUserData, ma_uint32 level, const char* pMessage);

struct ma_log_callback {
    ma_log_callback_proc onLog;                    // コールバック関数
    void* pUserData;                               // ユーザーデータ
};

struct ma_log {
    ma_log_callback callbacks[8];                  // コールバック配列
    ma_uint32 callbackCount;                       // コールバック数
};
```

### データソースプロシージャー

```c
typedef ma_data_source* (*ma_data_source_get_next_proc)(ma_data_source* pDataSource);
```

### デバイスコールバック

```c
// データコールバック
typedef void (*ma_device_data_proc)(
    ma_device* pDevice, 
    void* pOutput, 
    const void* pInput, 
    ma_uint32 frameCount
);

// 通知コールバック
typedef void (*ma_device_notification_proc)(
    ma_device* pDevice, 
    ma_device_notification_type notificationType
);

// デバイス初期化後の通知データ
struct ma_device_notification {
    ma_device_notification_type type;              // タイプ
    union {
        struct { /* プレイバック特有のデータ */ } playback;
        struct { /* キャプチャ特有のデータ */ } capture;
    };
};
```

### エンジンプロセスコールバック

```c
typedef void (*ma_engine_process_proc)(
    ma_engine* pEngine, 
    void* pOutput, 
    const void* pInput, 
    ma_uint64 frameCount
);
```

### ジョブプロシージャー

```c
typedef ma_result (*ma_job_proc)(ma_job* pJob);

struct ma_job {
    ma_job_type type;                              // タイプ
    union {
        struct { /* デバイス初期化 */ } deviceInit;
        struct { /* デバイス未初期化 */ } deviceUninit;
        // ... 他のジョブ固有データ
    };
};
```

---

## 9. ユーティリティ構造体 (Utility Structures)

### Allocation Callbacks

```c
typedef void* (* ma_allocate_proc)(size_t size, void* pUserData, ma_allocation_mode allocationMode);
typedef void (* ma_free_proc)(void* pObject, void* pUserData);
typedef ma_allocation_callbacks* (* ma_get_allocation_pointers_proc)(const ma_allocation_callbacks* pCallbacks, ma_allocation_callbacks* pPointers);

struct ma_allocation_callbacks {                 // カスタムアロケーションコールバック
    void* pUserData;                               // ユーザーデータ
    ma_allocate_proc onAllocate;                   // 割り当て関数
    ma_free_proc onFree;                           // 解放関数
    ma_get_allocation_pointers_proc onGetPointers; // ポインタ取得関数
};
```

### LCG (Linear Congruential Generator)

```c
struct ma_lcg {                                   // 線形合同法乱数生成器
    ma_uint64 state;                               // 状態
};
```

### Spinlock

```c
typedef ma_uint32 ma_spinlock;                    // スピンロック型
```

### Timer

```c
struct ma_timer {                                 // タイマー
    ma_int64 time;                                 // タイムスタンプ
};
```

### Device ID

```c
struct ma_device_id {                             // デバイスID
    const void* pBackendData;                      // バックエンドデータ
};
```

### Stack

```c
typedef struct {
    size_t itemSizeInBytes;                        // アイテムサイズ
    size_t capacityInItems;                        // 容量 (アイテム数)
    void* pBuffer;                                 // バッファ
} ma_stack;                                        // スタック構造体
```

### Slot Allocator

```c
typedef struct {
    size_t slotSizeInBytes;                        // スロットサイズ
    ma_uint32 slotCount;                           // スロット数
} ma_slot_allocator_config;

struct ma_slot_allocator_group {                  // スロットアロケーターグループ
    ma_uint8* pMemory;                            // メモリ
};

struct ma_slot_allocator {                        // スロットアロケーター
    ma_vector_map map;                             // マップ
};
```

### Job Queue

```c
typedef enum {
    ma_job_queue_flags_auto_start_thread = (1 << 0),
} ma_job_queue_flags;

typedef struct {
    ma_uint32 maxJobs;                             // 最大ジョブ数
    ma_uint32 jobThreadPriority;                   // ジョスレッド優先度
    ma_job_queue_flags flags;                      // フラグ
} ma_job_queue_config;

struct ma_job_queue {                              // ジョブキュー
    ma_spinlock lock;                               // ロック
    ma_job* pJobs;                                 // ジョブ配列
    ma_uint32 jobCount;                            // ジョブ数
};
```

### Fence

```c
struct ma_fence {                                  // フェンス同期オブジェクト
    const char* pName;                             // 名前
    ma_bool8 triggered;                            // トリガー済み
};
```

### Async Notification

```c
typedef struct {
    void (* onPoll)(void* pUserData, ma_uint64 currentFrame);
} ma_async_notification_callbacks;

struct ma_async_notification_poll {                // ポリング通知
    ma_async_notification_callbacks callbacks;     // コールバック
};

struct ma_async_notification_event {               // イベント通知
    HANDLE handle;                                 // Windowsハンドル
};
```

### VFS (Virtual File System)

```c
typedef struct {
    ma_result (*onOpen)(ma_vfs* pVFS, const char* pFilePath, ma_uint32 openModeFlags, ma_file** ppFile);
    ma_result (*onOpenW)(ma_vfs* pVFS, const wchar_t* pFilePath, ma_uint32 openModeFlags, ma_file** ppFile);
    ma_result (*onRead)(ma_vfs* pVFS, ma_file* pFile, void* pBufferOut, size_t bytesToRead, size_t* pBytesRead);
    // ... 他のVFSコールバック
} ma_vfs_callbacks;

struct ma_default_vfs {                            // デフォルトVFS
    ma_vfs vfs;                                    // ベースVFS
};

typedef struct {
    const char* pName;                             // ファイル名
    ma_uint64 sizeInBytes;                         // サイズ (バイト)
} ma_file_info;
```

---

## 10. ノード構造体 (Node Structures)

### Node Base

```c
struct ma_node_vtable {
    const char* pName;                             // ノード名
    ma_result (* onInitialize)(ma_node_config* pConfig, ma_node* pNode);
    ma_result (* onUninitialize)(ma_node* pNode);
    void (* onStart)(ma_node* pNode);
    void (* onStop)(ma_node* pNode);
    ma_result (* onMix)(ma_node* pNode, ma_node_mix_args* pMixArgs);
};

struct ma_node_config {
    ma_node_graph* pGraph;                         // ノードグラフ
    ma_node_flags flags;                           // フラグ
};

struct ma_node_base {                              // ノードベース (最初のメンバー)
    ma_node_vtable vtable;                         // vtable
    ma_node_graph* pGraph;                         // ノードグラフ
    ma_node_state state;                           // 状態
    ma_uint32 flags;                               // フラグ
};
```

### Node Graph

```c
struct ma_node_output_bus {                        // ノード出力バス
    ma_node* pNode;                                // ノード
    ma_uint32 busIndex;                            // バスインデックス
};

struct ma_node_input_bus {                         // ノード入力バス
    ma_node* pNode;                                // ノード
    ma_uint32 busIndex;                            // バスインデックス
};

struct ma_node_graph {                             // ノードグラフ
    ma_node_base nodeBase;                         // ノードベース (最初のメンバー)
    ma_stack preMixStack;                          // プリミックスタック
    ma_spinlock lock;                               // ロック
};
```

### Data Source Node

```c
typedef struct {
    ma_uint32 outputChannelCount;                  // 出力チャネル数
} ma_data_source_node_config;

struct ma_data_source_node {                       // データソースノード
    ma_node_base nodeBase;                         // ノードベース
    ma_data_source* pDataSource;                   // データソース
};
```

### Splitter Node

```c
typedef struct {
    ma_uint32 outputChannelCount;                  // 出力チャネル数
} ma_splitter_node_config;

struct ma_splitter_node {                          // スプリッターノード
    ma_node_base nodeBase;                         // ノードベース
};
```

### Filter Nodes (Biquad, LPF, HPF, BPF, Notch, Peak, LoShelf, HiShelf)

各フィルターノードは以下の構造を持つ:

```c
typedef struct {
    ma_uint32 outputChannelCount;                  // 出力チャネル数
} ma_<filter>_node_config;

struct ma_<filter>_node {                          // フィルターノード
    ma_node_base nodeBase;                         // ノードベース
    // フィルター固有のデータ
};
```

### Delay Node

```c
typedef struct {
    ma_uint32 outputChannelCount;                  // 出力チャネル数
} ma_delay_node_config;

struct ma_delay_node {                             // デレイノード
    ma_node_base nodeBase;                         // ノードベース
};
```

### Engine Node

```c
typedef enum {
    ma_engine_node_type_sound,                     // サウンドノード
    ma_engine_node_type_sound_group,               // サウンドグループノード
} ma_engine_node_type;

typedef struct {
    ma_uint32 outputChannelCount;                  // 出力チャネル数
} ma_engine_node_config;

struct ma_engine_node {                            // エンジンノード
    ma_node_base nodeBase;                         // ノードベース
    ma_engine_node_type type;                      // タイプ
};
```

### Sound Inline

```c
struct ma_sound_inlined {                          // インラインサウンド
    ma_sound_inlined* pNext;                       // 次のインラインサウンド
    ma_sound sound;                                // サウンド
};
```

---

## 11. コールバック関数ポインター (Callback Function Pointers)

### デコーディングバックエンド

```c
typedef struct {
    ma_result (* onGetBackendName)(void* pBackendData, const char** ppBackendName);
    ma_result (* onInitializeDecoder)(const ma_decoder_config* pConfig, void** ppBackendData, ma_decoding_backend_vtable** ppVTable);
} ma_decoding_backend_config;

typedef struct {
    ma_result (* onRead)(void* pBackendData, const void** ppFramesOut, ma_uint64 frameCount, ma_uint64* pFramesRead);
    ma_bool8 (* atEnd)(void* pBackendData);
    ma_result (* getLengthInPCMFrames)(void* pBackendData, ma_uint64* pFrameCountOut);
    ma_result (* seekToPCMFrame)(void* pBackendData, ma_uint64 frameNumber);
    ma_result (* getPositionInPCMFrames)(void* pBackendData, ma_uint64* pFrameCountOut);
} ma_decoding_backend_vtable;
```

### リサンプリングバックエンド

```c
typedef struct {
    ma_result (* onGetHeapSize)(void* pUserData, const ma_resampler_config* pConfig, size_t* pHeapSizeInBytes);
    ma_result (* onInit)(void* pUserData, const ma_resampler_config* pConfig, void* pHeap, ma_resampling_backend** ppBackend);
} ma_resampling_backend_vtable;
```

---

## 12. Go bindings の型マッピング

| C型 | Go型 | パッケージ |
|-----|------|------------|
| `ma_engine_config` | `EngineConfig` | miniaudio (miniaudio_types.go) |
| `ma_context_config` | `ContextConfig` | miniaudio (context.go) |
| `ma_device_config` | `DeviceConfig` | miniaudio (device.go) |
| `ma_decoder_config` | `DecoderConfig` | miniaudio (decoder.go) |
| `ma_engine` | `Engine` | miniaudio (engine.go) |
| `ma_context` | `Context` | miniaudio (context.go) |
| `ma_device` | `Device` | miniaudio (device.go) |
| `ma_decoder` | `Decoder` | miniaudio (decoder.go) |
| `ma_sound` | `Sound` | miniaudio (sound.go) |
| `ma_node_graph` | `NodeGraph` | miniaudio (nodegraph.go) |

---

## 付録: 主要定数

| 定数 | 値 | 説明 |
|-----|-----|------|
| `MA_MAX_CHANNELS` | 32 | 最大チャネル数 |
| `MA_ENGINE_MAX_LISTENERS` | 8 | 最大リスナー数 |