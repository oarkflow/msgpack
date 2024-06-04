package msgpack

import (
	"reflect"
)

type EncodeHook interface {
	BeforeMsgpackMarshal() error
}

func encodeHook(strct reflect.Value) error {
	if !strct.CanAddr() {
		return nil
	}
	if !strct.CanSet() {
		return nil
	}
	if hook, ok := strct.Addr().Interface().(EncodeHook); ok {
		if err := hook.BeforeMsgpackMarshal(); err != nil {
			return err
		}
	}
	return nil
}
