package miniaudio

/*
#include "miniaudio.h"
#include <stdlib.h>
#include <string.h>
*/
import "C"
import "unsafe"

// NodeGraphConfig wraps ma_node_graph_config.
type NodeGraphConfig struct {
	config *C.ma_node_graph_config
}

// NewNodeGraphConfig creates a new NodeGraphConfig and calls ma_node_graph_config_init.
func NewNodeGraphConfig(channels uint32) *NodeGraphConfig {
	config := C.ma_node_graph_config_init(C.ma_uint32(channels))
	return &NodeGraphConfig{config: &config}
}

func (ngc *NodeGraphConfig) Close() {
	if ngc.config == nil {
		return
	}
	ngc.config = nil
}

type Node struct {
	node *C.ma_node
}

// NodeGraph wraps ma_node_graph.
type NodeGraph struct {
	nodeGraph *C.ma_node_graph
}

// NewNodeGraph creates a new NodeGraph instance, initializes it with ma_node_graph_init, and returns it.
func NewNodeGraph(config *NodeGraphConfig) (*NodeGraph, error) {
	ng := (*C.ma_node_graph)(C.malloc(C.size_t(unsafe.Sizeof(*(*C.ma_node_graph)(nil)))))
	if ng == nil {
		return nil, ErrNilEngine
	}
	C.memset(unsafe.Pointer(ng), 0, C.size_t(unsafe.Sizeof(*ng)))

	var cConfig *C.ma_node_graph_config
	if config != nil {
		cConfig = config.config
	}

	var allocationCallbacks *C.ma_allocation_callbacks
	result := C.ma_node_graph_init(cConfig, allocationCallbacks, ng)
	if result != 0 {
		C.free(unsafe.Pointer(ng))
		return nil, ErrResult(result)
	}

	return &NodeGraph{nodeGraph: ng}, nil
}

// Close uninitializes and releases the node graph memory.
func (ng *NodeGraph) Close() {
	if ng.nodeGraph == nil {
		return
	}
	C.ma_node_graph_uninit(ng.nodeGraph, nil)
	C.free(unsafe.Pointer(ng.nodeGraph))
	ng.nodeGraph = nil
}

// Endpoint returns the endpoint node of the node graph.
func (ng *NodeGraph) Endpoint() *Node {
	if ng.nodeGraph == nil {
		return nil
	}
	node := (*C.ma_node)(unsafe.Pointer(C.ma_node_graph_get_endpoint(ng.nodeGraph)))
	return &Node{node: node}
}

// ReadPCMFrames reads PCM frames from the node graph.
func (ng *NodeGraph) ReadPCMFrames(framesOut []float32, frameCount uint64) (uint64, error) {
	if ng.nodeGraph == nil {
		return 0, ErrNilEngine
	}
	var framesRead C.ma_uint64
	result := C.ma_node_graph_read_pcm_frames(ng.nodeGraph, unsafe.Pointer(&framesOut[0]), C.ma_uint64(frameCount), &framesRead)
	if result != 0 {
		return uint64(framesRead), ErrResult(result)
	}
	return uint64(framesRead), nil
}

// Channels returns the number of channels of the node graph.
func (ng *NodeGraph) Channels() uint32 {
	if ng.nodeGraph == nil {
		return 0
	}
	return uint32(C.ma_node_graph_get_channels(ng.nodeGraph))
}

// Time returns the current time of the node graph.
func (ng *NodeGraph) Time() uint64 {
	if ng.nodeGraph == nil {
		return 0
	}
	return uint64(C.ma_node_graph_get_time(ng.nodeGraph))
}

// SetTime sets the time of the node graph.
func (ng *NodeGraph) SetTime(globalTime uint64) error {
	if ng.nodeGraph == nil {
		return ErrNilEngine
	}
	result := C.ma_node_graph_set_time(ng.nodeGraph, C.ma_uint64(globalTime))
	if result != 0 {
		return ErrResult(result)
	}
	return nil
}

// ProcessingSizeInFrames returns the processing size in frames of the node graph.
func (ng *NodeGraph) ProcessingSizeInFrames() uint32 {
	if ng.nodeGraph == nil {
		return 0
	}
	return uint32(C.ma_node_graph_get_processing_size_in_frames(ng.nodeGraph))
}
