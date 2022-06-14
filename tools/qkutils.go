package tools

import (
	"encoding/json"
	"reflect"
)

func GetTypeByName(v string) reflect.Type {
	var av interface{}
	switch v {
	case "string":
		av = ""
	case "int":
		av = 0
	case "int64":
		av = int64(0)
	case "int32":
		av = int32(0)
	case "int16":
		av = int16(0)
	case "int8":
		av = int8(0)
	case "uint":
		av = uint(0)
	case "uint64":
		av = uint64(0)
	case "uint32":
		av = uint32(0)
	case "uint16":
		av = uint16(0)
	case "uint8":
		av = uint8(0)
	case "float64":
		av = float64(0)
	case "float32":
		av = float32(0)
	case "bool":
		av = false
	case "byte":
		av = byte(0)
	case "rune":
		av = rune(0)
	case "struct":
		av = struct{}{}
	case "slice":
		av = []interface{}{}
	case "map":
		av = map[interface{}]interface{}{}
	case "chan":
		av = make(chan interface{})
	case "func":
		av = func() {}
	case "pointer":
		av = &struct{}{}
	case "json.RawMessage":
		av = json.RawMessage("")
	default:
		av = nil
	}
	return reflect.TypeOf(av)
}
