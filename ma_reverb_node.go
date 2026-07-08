package miniaudio

/*
#include "ma_reverb_node.h"
#include "verblib.h"
#include "stdlib.h"
*/
import "C"
import "unsafe"

type ReverveNodeConfig struct {
	config *C.ma_reverb_node_config
}

type ReverveNode struct {
	node *C.ma_reverb_node
}

func NewReverveNodeConfig(channels, sampleRate uint32) *ReverveNodeConfig {
	config := (*C.ma_reverb_node_config)(C.malloc(C.size_t(unsafe.Sizeof(*(*C.ma_reverb_node_config)(nil)))))
	*config = C.ma_reverb_node_config_init(C.ma_uint32(channels), C.ma_uint32(sampleRate))
	return &ReverveNodeConfig{
		config: config,
	}
}

func (c *ReverveNodeConfig) Free() {
	if c.config != nil {
		C.free(unsafe.Pointer(c.config))
		c.config = nil
	}
}

// RoomSize は部屋のサイズを返します。範囲: 0.0 〜 1.0
func (c *ReverveNodeConfig) RoomSize() float32 {
	return float32(c.config.roomSize)
}

// SetRoomSize は部屋のサイズを設定します。範囲: 0.0 〜 1.0
func (c *ReverveNodeConfig) SetRoomSize(roomSize float32) {
	c.config.roomSize = C.float(roomSize)
}

// Damping はダンピング量を返します。範囲: 0.0 〜 1.0
func (c *ReverveNodeConfig) Damping() float32 {
	return float32(c.config.damping)
}

// SetDamping はダンピング量を設定します。範囲: 0.0 〜 1.0
func (c *ReverveNodeConfig) SetDamping(damping float32) {
	c.config.damping = C.float(damping)
}

// Width はステレオ幅を返します。範囲: 0.0 〜 1.0
func (c *ReverveNodeConfig) Width() float32 {
	return float32(c.config.width)
}

// SetWidth はステレオ幅を設定します。範囲: 0.0 〜 1.0
func (c *ReverveNodeConfig) SetWidth(width float32) {
	c.config.width = C.float(width)
}

// WetVolume はウェット信号の音量を返します。範囲: 0.0 〜 1.0
func (c *ReverveNodeConfig) WetVolume() float32 {
	return float32(c.config.wetVolume)
}

// SetWetVolume はウェット信号の音量を設定します。範囲: 0.0 〜 1.0
func (c *ReverveNodeConfig) SetWetVolume(wetVolume float32) {
	c.config.wetVolume = C.float(wetVolume)
}

// DryVolume はドライ信号の音量を返します。範囲: 0.0 〜 1.0
func (c *ReverveNodeConfig) DryVolume() float32 {
	return float32(c.config.dryVolume)
}

// SetDryVolume はドライ信号の音量を設定します。範囲: 0.0 〜 1.0
func (c *ReverveNodeConfig) SetDryVolume(dryVolume float32) {
	c.config.dryVolume = C.float(dryVolume)
}

// Mode はリバーブのモードを返します。0.5 未満は通常、それ以上はフリーズ（凍結）モード
func (c *ReverveNodeConfig) Mode() float32 {
	return float32(c.config.mode)
}

// SetMode はリバーブのモードを設定します。0.5 未満は通常、それ以上はフリーズ（凍結）モード
func (c *ReverveNodeConfig) SetMode(mode float32) {
	c.config.mode = C.float(mode)
}

func NewReverveNode(nodeGraph *NodeGraph, config *ReverveNodeConfig, callbacks *AllocationCallbacks) *ReverveNode {
	n := (*C.ma_reverb_node)(C.malloc(C.size_t(unsafe.Sizeof(*(*C.ma_reverb_node)(nil)))))
	var cbs *C.ma_allocation_callbacks
	if callbacks != nil {
		cbs = callbacks.callbacks
	}
	result := C.ma_reverb_node_init(nodeGraph.nodeGraph, config.config, cbs, n)
	if result != 0 {
		return nil
	}
	node := &ReverveNode{
		node: n,
	}
	node.SetRoomSize(config.RoomSize())
	node.SetDamping(config.Damping())
	node.SetWidth(config.Width())
	node.SetWetVolume(config.WetVolume())
	node.SetDryVolume(config.DryVolume())
	node.SetMode(config.Mode())
	return node
}

func (n *ReverveNode) Close() {
	if n.node == nil {
		return
	}
	C.ma_reverb_node_uninit(n.node, nil)
	C.free(unsafe.Pointer(n.node))
	n.node = nil
}

// RoomSize は部屋のサイズを返します。範囲: 0.0 〜 1.0
func (n *ReverveNode) RoomSize() float32 {
	return float32(C.verblib_get_room_size(&n.node.reverb))
}

// SetRoomSize は部屋のサイズを設定します。範囲: 0.0 〜 1.0
func (n *ReverveNode) SetRoomSize(roomSize float32) {
	C.verblib_set_room_size(&n.node.reverb, C.float(roomSize))
}

// Damping はダンピング量を返します。範囲: 0.0 〜 1.0
func (n *ReverveNode) Damping() float32 {
	return float32(C.verblib_get_damping(&n.node.reverb))
}

// SetDamping はダンピング量を設定します。範囲: 0.0 〜 1.0
func (n *ReverveNode) SetDamping(damping float32) {
	C.verblib_set_damping(&n.node.reverb, C.float(damping))
}

// Width はステレオ幅を返します。範囲: 0.0 〜 1.0
func (n *ReverveNode) Width() float32 {
	return float32(C.verblib_get_width(&n.node.reverb))
}

// SetWidth はステレオ幅を設定します。範囲: 0.0 〜 1.0
func (n *ReverveNode) SetWidth(width float32) {
	C.verblib_set_width(&n.node.reverb, C.float(width))
}

// WetVolume はウェット信号の音量を返します。範囲: 0.0 〜 1.0
func (n *ReverveNode) WetVolume() float32 {
	return float32(C.verblib_get_wet(&n.node.reverb))
}

// SetWetVolume はウェット信号の音量を設定します。範囲: 0.0 〜 1.0
func (n *ReverveNode) SetWetVolume(wetVolume float32) {
	C.verblib_set_wet(&n.node.reverb, C.float(wetVolume))
}

// DryVolume はドライ信号の音量を返します。範囲: 0.0 〜 1.0
func (n *ReverveNode) DryVolume() float32 {
	return float32(C.verblib_get_dry(&n.node.reverb))
}

// SetDryVolume はドライ信号の音量を設定します。範囲: 0.0 〜 1.0
func (n *ReverveNode) SetDryVolume(dryVolume float32) {
	C.verblib_set_dry(&n.node.reverb, C.float(dryVolume))
}

// Mode はリバーブのモードを返します。0.5 未満は通常、それ以上はフリーズ（凍結）モード
func (n *ReverveNode) Mode() float32 {
	return float32(C.verblib_get_mode(&n.node.reverb))
}

// SetMode はリバーブのモードを設定します。0.5 未満は通常、それ以上はフリーズ（凍結）モード
func (n *ReverveNode) SetMode(mode float32) {
	C.verblib_set_mode(&n.node.reverb, C.float(mode))
}
