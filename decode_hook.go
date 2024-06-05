package msgpack

import (
	"reflect"
	"unsafe"
)

var decodeHookType = reflect.TypeOf((*DecodeHook)(nil)).Elem()

type DecodeHook interface {
	AfterMsgpackUnmarshal() error
}

func decodeHook(v reflect.Value) error {
	if !v.CanAddr() {
		return nil
	}
	vPtr := v.Addr()

	// early exit
	if !vPtr.Type().AssignableTo(decodeHookType) {
		return nil
	}

	if !vPtr.CanInterface() {
		vPtr = reflect.NewAt(v.Type(), unsafe.Pointer(v.UnsafeAddr()))
	}

	if hook, ok := vPtr.Interface().(DecodeHook); ok {
		if err := hook.AfterMsgpackUnmarshal(); err != nil {
			return err
		}
	}

	return nil
}
