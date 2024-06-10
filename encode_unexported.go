package msgpack

import (
	"reflect"
	"sort"
	"unsafe"
)

func (e *Encoder) reflectStringSlice(v reflect.Value) []string {
	if e.flags&encodeIncludeUnexportedFlag == 0 || v.CanInterface() {
		return v.Interface().([]string)
	}
	return unsafe.Slice((*string)(v.UnsafePointer()), v.Len())
}

func (e *Encoder) writeByteArrayUnexportedToBuf(v reflect.Value) {
	n := v.Len()
	for i := 0; i < n; i++ {
		e.buf[i] = byte(v.Index(i).Uint())
	}
}

// Iterate over reflect.Value which is a map[string]string and encode its entries.
// Since reflect.Value v is an unexported field, we can not `.Interface().(map[string]string)`.
// Instead, we use `.MapRange()` to iterate over its entries.
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

// Iterate over reflect.Value which is a map[string]bool and encode its entries.
// Since reflect.Value v is an unexported field, we can not `.Interface().(map[string]bool)`.
// Instead, we use `.MapRange()` to iterate over its entries.
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

// Iterate over reflect.Value which is a map[string]string and encode its entries ordered by keys.
// Since reflect.Value v is an unexported field, we can not `.Interface().(map[string]string)`.
// Instead, we use `.MapRange()` to iterate over its entries.
func encodeSortedMapStringStringUnexported(e *Encoder, v reflect.Value) error {
	if v.Kind() != reflect.Map {
		panic("expect reflect.Map")
	}
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

// Iterate over reflect.Value which is a map[string]bool and encode its entries ordered by keys.
// Since reflect.Value v is an unexported field, we can not `.Interface().(map[string]bool)`.
// Instead, we use `.MapRange()` to iterate over its entries.
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
