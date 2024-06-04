package msgpack

import (
	"reflect"
	"unsafe"
)

func (e *Encoder) reflectStringSlice(v reflect.Value) []string {
	if v.CanInterface() {
		return v.Interface().([]string)
	}
	return unsafe.Slice((*string)(v.UnsafePointer()), v.Len())
}
