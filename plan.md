# Goal

WebAudioと同様の機能をネイティブで実現したい
その手始めとしてGoからminiaudioの高レベルAPIを利用可能にする

# Background

WebAudioに近いネイティブオーディオライブラリが欲しいが、いまだ存在していないのでこのminiaudioを足掛かりにする

# Requirements

必須要件
- Goから利用する機能は全てminiaudioのC APIを呼び出して実現する
- Go実装やCラッパー実装は最小限の機能に留める
- 音源の3D空間における進行方向や進行速度の設定が可能
- リスナーの3D空間における位置と向きが設定可能
- 距離減衰モデルを設定できるようにする
- Doppler効果に対応
- ノードグラフによるエフェクトの構築が可能

# Non-goals

今回はやらないこと
- miniaudioに無い機能や構造体定義を作る
- ソフトウェアDSPや3D処理をGo実装
- CGOのコメントで*.cを参照すること
- Cのポインタを持たない構造体をGoで定義しない

# Design

設計方針
- miniaudioの命名規則は「ma_object_method」で、
- miniaudio.hの構造体定義はそのポインタをGoの構造体でラップする
- NewObject関数でC側mallocしてC.ma_object_init関数を呼び出す
- C.ma_object_initに必要な引数はNewObject関数にも持たせる
- Object.Close()でC側メモリをfreeする。
- Object.Init()は不要
- Goでは Object.Method のようにメソッドとして表現する。
- メソッドの引数がオブジェクトポインタの場合、Goでラップした構造体ポインタを受け取り、CのAPI関数呼び出し時にCのオブジェクトポインタを渡す
- CのAPIとGoのメソッドは１対１に対応させ、CのAPIを複数呼び出すような操作は実装しない
- GoのAPI上必要な定数はGoで再定義するが最小限にすること
- Context,ContextConfig -> context.go
- Device,DeviceConfig -> device.go
- Engine,EngineConfig -> engine.go
- Sound,SoundConfig -> sound.go
- Decoder,DecoderConfig -> decoder.go
- Encoder,EncoderConfig -> encoder.go
- Node,NodeGraph,NodeGraphConfig -> nodegraph.go
- デバイスのコールバックは miniaudio_callbacks.h と miniaudio_callbacks.c に Go の関数を呼び出すためのブリッジ関数を定義
- miniaudio_callbacks.h と miniaudio_callbacks.c はC-APIを呼び出すための定義を書くのは禁止
- ma_data_source*の引数はGo側では type DataSource interface {dataSourcePtr() unsafe.Pointer}定義の DataSourceインターフェス型引数とする。
- CGOでは同パッケージ内の*.cはgo buildでCコンパイルされリンクする挙動なので別途コンパイルやリンカー指定は無用


# Tasks

- [x] Step1: miniaudio.hの定義の確認
- [x] Step2: コールバック処理のブリッジ実装
- [x] Step3: Engine、Soundの各関数のラッパー実装
- [x] Step4: 音が鳴るサンプルTestコードの実装
- [x] Step5: テストの実行と成功を確認
- [x] Step6: Engineのテストコードを追加実装

# Constraints

守るべきルール
- Step5までの実装は成功しているのでStep5までの実装を参考にして
- miniaudio.h, miniaudio.c はパスも中身も改変禁止
- miniaudio_callbacks.hとminiaudio_callbacks.cはGoからCを呼び出すためのコードではない。
- miniaudio.hの定義をそのまま再定義することは極力避けて、再定義数を極力少なく保って。
- Go1.23対応

# Acceptance Criteria

完了条件
- Engineのテストコード完走
