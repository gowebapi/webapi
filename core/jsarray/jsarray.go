// Code generated. DO NOT EDIT.
//
// (C) Copyright Martin Juhlin 2019.
//
// See LICENSE.md for license information

package jsarray

//go:generate go run generator.go

import (
	"github.com/gowebapi/webapi/core/js"
)

func UInt8ToJS(values []uint8) js.Value {
	tmp := make([]interface{}, len(values))
	for i := 0; i < len(values); i++ {
		tmp[i] = values[i]
	}
	array := js.Global().Get("Array").New(tmp...)
	return array
}

func UInt8ToGo(array js.Value) []uint8 {
	size := array.Length()
	out := make([]uint8, size)
	for i := 0; i < size; i++ {
		out[i] = uint8(array.Index(i).Int())
	}
	return out
}

func Int8ToJS(values []int8) js.Value {
	tmp := make([]interface{}, len(values))
	for i := 0; i < len(values); i++ {
		tmp[i] = values[i]
	}
	array := js.Global().Get("Array").New(tmp...)
	return array
}

func Int8ToGo(array js.Value) []int8 {
	size := array.Length()
	out := make([]int8, size)
	for i := 0; i < size; i++ {
		out[i] = int8(array.Index(i).Int())
	}
	return out
}

func UInt16ToJS(values []uint16) js.Value {
	tmp := make([]interface{}, len(values))
	for i := 0; i < len(values); i++ {
		tmp[i] = values[i]
	}
	array := js.Global().Get("Array").New(tmp...)
	return array
}

func UInt16ToGo(array js.Value) []uint16 {
	size := array.Length()
	out := make([]uint16, size)
	for i := 0; i < size; i++ {
		out[i] = uint16(array.Index(i).Int())
	}
	return out
}

func Int16ToJS(values []int16) js.Value {
	tmp := make([]interface{}, len(values))
	for i := 0; i < len(values); i++ {
		tmp[i] = values[i]
	}
	array := js.Global().Get("Array").New(tmp...)
	return array
}

func Int16ToGo(array js.Value) []int16 {
	size := array.Length()
	out := make([]int16, size)
	for i := 0; i < size; i++ {
		out[i] = int16(array.Index(i).Int())
	}
	return out
}

func UInt32ToJS(values []uint32) js.Value {
	tmp := make([]interface{}, len(values))
	for i := 0; i < len(values); i++ {
		tmp[i] = values[i]
	}
	array := js.Global().Get("Array").New(tmp...)
	return array
}

func UInt32ToGo(array js.Value) []uint32 {
	size := array.Length()
	out := make([]uint32, size)
	for i := 0; i < size; i++ {
		out[i] = uint32(array.Index(i).Int())
	}
	return out
}

func Int32ToJS(values []int32) js.Value {
	tmp := make([]interface{}, len(values))
	for i := 0; i < len(values); i++ {
		tmp[i] = values[i]
	}
	array := js.Global().Get("Array").New(tmp...)
	return array
}

func Int32ToGo(array js.Value) []int32 {
	size := array.Length()
	out := make([]int32, size)
	for i := 0; i < size; i++ {
		out[i] = int32(array.Index(i).Int())
	}
	return out
}

func Float32ToJS(values []float32) js.Value {
	tmp := make([]interface{}, len(values))
	for i := 0; i < len(values); i++ {
		tmp[i] = values[i]
	}
	array := js.Global().Get("Array").New(tmp...)
	return array
}

func Float32ToGo(array js.Value) []float32 {
	size := array.Length()
	out := make([]float32, size)
	for i := 0; i < size; i++ {
		out[i] = float32(array.Index(i).Float())
	}
	return out
}

func Float64ToJS(values []float64) js.Value {
	tmp := make([]interface{}, len(values))
	for i := 0; i < len(values); i++ {
		tmp[i] = values[i]
	}
	array := js.Global().Get("Array").New(tmp...)
	return array
}

func Float64ToGo(array js.Value) []float64 {
	size := array.Length()
	out := make([]float64, size)
	for i := 0; i < size; i++ {
		out[i] = float64(array.Index(i).Float())
	}
	return out
}
