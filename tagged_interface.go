package msgpack

import (
	"fmt"
	"reflect"

	"github.com/KyberNetwork/msgpack/v5/msgpcode"
)

const taggedInterfaceExtID = 127

var taggedInterfaceTypeMap = map[string]reflect.Type{}

func getConcreteTypeTag(v reflect.Value) (string, reflect.Type) {
	if v.Kind() == reflect.Pointer {
		if v.IsNil() || v.Elem().Kind() != reflect.Struct {
			return "", nil
		}
	} else if v.Kind() != reflect.Struct {
		return "", nil
	}
	var (
		isPointer bool
		typ       reflect.Type
	)
	if v.Kind() == reflect.Pointer {
		isPointer = true
		typ = v.Elem().Type()
	} else {
		typ = v.Type()
	}
	if isPointer {
		return fmt.Sprintf("*%s %s", typ.PkgPath(), typ.Name()), typ
	}
	return fmt.Sprintf("%s %s", typ.PkgPath(), typ.Name()), typ
}

func RegisterConcreteType(v interface{}) error {
	tag, typ := getConcreteTypeTag(reflect.ValueOf(v))
	if typ == nil {
		return fmt.Errorf("expect struct or pointer to struct")
	}
	taggedInterfaceTypeMap[tag] = typ
	return nil
}

func (e *Encoder) encodeTaggedInterface(tag string, v reflect.Value) error {
	err := e.writeCode(msgpcode.FixExt1)
	if err != nil {
		return err
	}
	err = e.writeCode(taggedInterfaceExtID)
	if err != nil {
		return err
	}
	err = e.EncodeString(tag)
	if err != nil {
		return err
	}
	err = e.EncodeValue(v.Elem())
	if err != nil {
		return err
	}
	return nil
}

func (d *Decoder) decodeTaggedInterface() (interface{}, error) {
	tag, err := d.DecodeString()
	if err != nil {
		return nil, err
	}

	typ, ok := taggedInterfaceTypeMap[tag]
	if !ok {
		return nil, fmt.Errorf("unregistered concrete type")
	}

	v := reflect.New(typ)
	err = d.DecodeValue(v)
	if err != nil {
		return nil, err
	}

	return v.Interface(), nil
}
