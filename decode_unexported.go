package msgpack

import (
	"reflect"
	"unsafe"
)

func reflectExportValue(v reflect.Value) reflect.Value {
	return reflect.NewAt(v.Type(), unsafe.Pointer(v.UnsafeAddr())).Elem()
}

func (d *Decoder) reflectSetBool(v reflect.Value, x bool) {
	if d.flags&decodeIncludeUnexportedFlag == 0 || v.CanSet() {
		v.SetBool(x)
		return
	}
	reflectExportValue(v).SetBool(x)
}

func (d *Decoder) reflectSetInt(v reflect.Value, x int64) {
	if d.flags&decodeIncludeUnexportedFlag == 0 || v.CanSet() {
		v.SetInt(x)
		return
	}
	reflectExportValue(v).SetInt(x)
}

func (d *Decoder) reflectSetUint(v reflect.Value, x uint64) {
	if d.flags&decodeIncludeUnexportedFlag == 0 || v.CanSet() {
		v.SetUint(x)
		return
	}
	reflectExportValue(v).SetUint(x)
}

func (d *Decoder) reflectSetFloat(v reflect.Value, x float64) {
	if d.flags&decodeIncludeUnexportedFlag == 0 || v.CanSet() {
		v.SetFloat(x)
		return
	}
	reflectExportValue(v).SetFloat(x)
}

func (d *Decoder) reflectSetString(v reflect.Value, x string) {
	if d.flags&decodeIncludeUnexportedFlag == 0 || v.CanSet() {
		v.SetString(x)
		return
	}
	reflectExportValue(v).SetString(x)
}

func (d *Decoder) reflectSetBytes(v reflect.Value, x []byte) {
	if d.flags&decodeIncludeUnexportedFlag == 0 || v.CanSet() {
		v.SetBytes(x)
		return
	}
	reflectExportValue(v).SetBytes(x)
}

func (d *Decoder) reflectSet(v reflect.Value, x reflect.Value) {
	if d.flags&decodeIncludeUnexportedFlag == 0 || v.CanSet() {
		v.Set(x)
		return
	}
	reflectExportValue(v).Set(x)
}

func (d *Decoder) reflectSetMapIndex(v reflect.Value, key reflect.Value, elem reflect.Value) {
	if d.flags&decodeIncludeUnexportedFlag == 0 || v.CanSet() {
		v.SetMapIndex(key, elem)
		return
	}
	reflectExportValue(v).SetMapIndex(key, elem)
}

func (d *Decoder) reflectStringSlicePtr(vPtr reflect.Value) *[]string {
	if d.flags&decodeIncludeUnexportedFlag == 0 || vPtr.CanInterface() {
		return vPtr.Interface().(*[]string)
	}
	return (*[]string)(vPtr.UnsafePointer())
}

func (d *Decoder) reflectMapStringStringPtr(vPtr reflect.Value) *map[string]string {
	if d.flags&decodeIncludeUnexportedFlag == 0 || vPtr.CanInterface() {
		return vPtr.Interface().(*map[string]string)
	}
	return (*map[string]string)(vPtr.UnsafePointer())
}
