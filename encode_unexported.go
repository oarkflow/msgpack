package msgpack

import (
	"reflect"
	"sort"
	"unsafe"
)

func (e *Encoder) reflectStringSlice(v reflect.Value) []string {
	if v.CanInterface() {
		return v.Interface().([]string)
	}
	return unsafe.Slice((*string)(v.UnsafePointer()), v.Len())
}

func encodeMapStringBoolValueUnexported(e *Encoder, v reflect.Value) error {
	iter := v.MapRange()
	for iter.Next() {
		mk := iter.Key().String()
		mv := iter.Value().Bool()
		if err := e.EncodeString(mk); err != nil {
			return err
		}
		if err := e.EncodeBool(mv); err != nil {
			return err
		}
	}
	return nil
}

func encodeMapStringStringValueUnexported(e *Encoder, v reflect.Value) error {
	iter := v.MapRange()
	for iter.Next() {
		mk := iter.Key().String()
		mv := iter.Value().String()
		if err := e.EncodeString(mk); err != nil {
			return err
		}
		if err := e.EncodeString(mv); err != nil {
			return err
		}
	}
	return nil
}

func encodeSortedMapStringStringUnexported(e *Encoder, v reflect.Value) error {
	rkeys := v.MapKeys()
	keys := make([]string, len(rkeys))
	for i, rkey := range rkeys {
		keys[i] = rkey.String()
	}
	sort.Strings(keys)

	for _, k := range keys {
		err := e.EncodeString(k)
		if err != nil {
			return err
		}
		val := v.MapIndex(reflect.ValueOf(k)).String()
		if err = e.EncodeString(val); err != nil {
			return err
		}
	}

	return nil
}

func encodeSortedMapStringBoolUnexported(e *Encoder, v reflect.Value) error {
	rkeys := v.MapKeys()
	keys := make([]string, len(rkeys))
	for i, rkey := range rkeys {
		keys[i] = rkey.String()
	}
	sort.Strings(keys)

	for _, k := range keys {
		err := e.EncodeString(k)
		if err != nil {
			return err
		}
		val := v.MapIndex(reflect.ValueOf(k)).Bool()
		if err = e.EncodeBool(val); err != nil {
			return err
		}
	}

	return nil
}
