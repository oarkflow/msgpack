package msgpack

import "reflect"

type DecodeHook interface {
	AfterMsgpackUnmarshal() error
}

func decodeHook(v reflect.Value) error {
	if !v.CanAddr() {
		return nil
	}
	if !v.CanSet() {
		return nil
	}
	if hook, ok := v.Addr().Interface().(DecodeHook); ok {
		if err := hook.AfterMsgpackUnmarshal(); err != nil {
			return err
		}
	}
	return nil
}
