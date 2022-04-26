//go:build !baremetal
// +build !baremetal

package bluetooth

import (
	"github.com/godbus/dbus/v5"
)

func convertVariantMapToAnyMap(in map[uint16]dbus.Variant) map[uint16]interface{} {
	out := make(map[uint16]interface{}, len(in))

	for k := range in {
		out[k] = in[k].Value()
	}

	return out
}
