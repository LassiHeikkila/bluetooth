// +build softdevice,s113v7

package bluetooth

/*
// Add the correct SoftDevice include path to CFLAGS, so #include will work as
// expected.
#cgo CFLAGS: -Is113_nrf52_7.0.1/s113_nrf52_7.0.1_API/include

#include "nrf_nvic.h"
nrf_nvic_state_t nrf_nvic_state = {0};
*/
import "C"
