package msgpack

import (
	"reflect"
	"unsafe"
)

var encodeHookType = reflect.TypeOf((*EncodeHook)(nil)).Elem()

type EncodeHook interface {
	BeforeMsgpackMarshal() error
}

func encodeHook(v reflect.Value) error {
	if !v.CanAddr() {
		return nil
	}
	vPtr := v.Addr()

	// early exit
	if !vPtr.Type().AssignableTo(encodeHookType) {
		return nil
	}

	if !vPtr.CanInterface() {
		vPtr = reflect.NewAt(v.Type(), unsafe.Pointer(v.UnsafeAddr()))
	}

	if hook, ok := vPtr.Interface().(EncodeHook); ok {
		if err := hook.BeforeMsgpackMarshal(); err != nil {
			return err
		}
	}
	return nil
}
