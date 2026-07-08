package miniaudio

/*
#include "miniaudio.h"
#include <stdlib.h>
*/
import "C"
import "unsafe"

type AllocationCallbacks struct {
	callbacks *C.ma_allocation_callbacks
}

func NewAllocationCallbacks() *AllocationCallbacks {
	cbs := (*C.ma_allocation_callbacks)(C.malloc(C.size_t(unsafe.Sizeof(*(*C.ma_allocation_callbacks)(nil)))))
	return &AllocationCallbacks{
		callbacks: cbs,
	}
}

func (c *AllocationCallbacks) Close() {
	if c.callbacks == nil {
		return
	}
	C.free(unsafe.Pointer(c.callbacks))
	c.callbacks = nil
}
