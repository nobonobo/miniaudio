/* Bridge functions for Go callbacks from miniaudio C APIs. */
#include "miniaudio.h"

/* This file is intentionally left mostly empty.
 * It is included for callback bridging only - when miniaudio needs to call back into Go code.
 * Do NOT add simple wrapper functions here that just call ma_* functions directly.
 * Those should be called from Go via CGO instead. */