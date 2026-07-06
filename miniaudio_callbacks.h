#ifndef MINIAUDIO_CALLBACKS_H
#define MINIAUDIO_CALLBACKS_H

#include "miniaudio.h"

#ifdef __cplusplus
extern "C" {
#endif

/* This header is intentionally minimal.
 * It is included for callback bridging only - when miniaudio needs to call back into Go code.
 * Do NOT add simple wrapper function declarations here that just call ma_* functions directly.
 * Those should be called from Go via CGO instead. */

#ifdef __cplusplus
}
#endif

#endif // MINIAUDIO_CALLBACKS_H