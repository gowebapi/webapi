// Code generated. DO NOT EDIT.
//
// (C) Copyright Martin Juhlin 2019.
//
// See LICENSE.md for license information

package jsconv

//go:generate go run generator.go

import (
	"reflect"
	"runtime"
	"unsafe"

	"github.com/gowebapi/webapi/core/js"
)

// UInt8ToJs is converting []uint8 to a new javascript Uint8Array.
func UInt8ToJs(src []uint8) js.Value {
	array := js.Global().Get("Uint8Array").New(len(src) * 1)
	js.CopyBytesToJS(array, src)
	return array
}

// JsToUInt8 convert javascript Uint8Array to Go uint8.
func JsToUInt8(array js.Value) []uint8 {
	size := array.Get("length").Int()

	mem := make([]byte, size)
	js.CopyBytesToGo(mem, array)

	head := (*reflect.SliceHeader)(unsafe.Pointer(&mem))
	dst := *(*[]uint8)(unsafe.Pointer(head))
	runtime.KeepAlive(mem)
	return dst
}

// Int8ToJs is converting []int8 to a new javascript Int8Array.
func Int8ToJs(src []int8) js.Value {
	array := js.Global().Get("Uint8Array").New(len(src) * 1)

	// pointer converter
	head := (*reflect.SliceHeader)(unsafe.Pointer(&src))
	data := *(*[]byte)(unsafe.Pointer(head))

	// copy
	js.CopyBytesToJS(array, data)
	runtime.KeepAlive(src)

	// convert back
	buf := array.Get("buffer")
	return js.Global().Get("Int8Array").New(
		buf,
		array.Get("byteOffset"),
		array.Get("byteLength").Int()/1,
	)
}

// JsToInt8 convert javascript Int8Array to Go int8.
func JsToInt8(array js.Value) []int8 {
	size := array.Get("length").Int()
	buf := array.Get("buffer")
	tmp := js.Global().Get("Uint8Array").New(buf)

	mem := make([]byte, size*1)
	js.CopyBytesToGo(mem, tmp)

	head := (*reflect.SliceHeader)(unsafe.Pointer(&mem))
	head.Cap /= 1
	head.Len /= 1
	dst := *(*[]int8)(unsafe.Pointer(head))
	runtime.KeepAlive(mem)
	return dst
}

// UInt16ToJs is converting []uint16 to a new javascript Uint16Array.
func UInt16ToJs(src []uint16) js.Value {
	array := js.Global().Get("Uint8Array").New(len(src) * 2)

	// pointer converter
	head := (*reflect.SliceHeader)(unsafe.Pointer(&src))
	head.Len *= 2
	head.Cap *= 2
	data := *(*[]byte)(unsafe.Pointer(head))

	// copy
	js.CopyBytesToJS(array, data)
	runtime.KeepAlive(src)

	// convert back
	buf := array.Get("buffer")
	return js.Global().Get("Uint16Array").New(
		buf,
		array.Get("byteOffset"),
		array.Get("byteLength").Int()/2,
	)
}

// JsToUInt16 convert javascript Uint16Array to Go uint16.
func JsToUInt16(array js.Value) []uint16 {
	size := array.Get("length").Int()
	buf := array.Get("buffer")
	tmp := js.Global().Get("Uint8Array").New(buf)

	mem := make([]byte, size*2)
	js.CopyBytesToGo(mem, tmp)

	head := (*reflect.SliceHeader)(unsafe.Pointer(&mem))
	head.Cap /= 2
	head.Len /= 2
	dst := *(*[]uint16)(unsafe.Pointer(head))
	runtime.KeepAlive(mem)
	return dst
}

// Int16ToJs is converting []int16 to a new javascript Int16Array.
func Int16ToJs(src []int16) js.Value {
	array := js.Global().Get("Uint8Array").New(len(src) * 2)

	// pointer converter
	head := (*reflect.SliceHeader)(unsafe.Pointer(&src))
	head.Len *= 2
	head.Cap *= 2
	data := *(*[]byte)(unsafe.Pointer(head))

	// copy
	js.CopyBytesToJS(array, data)
	runtime.KeepAlive(src)

	// convert back
	buf := array.Get("buffer")
	return js.Global().Get("Int16Array").New(
		buf,
		array.Get("byteOffset"),
		array.Get("byteLength").Int()/2,
	)
}

// JsToInt16 convert javascript Int16Array to Go int16.
func JsToInt16(array js.Value) []int16 {
	size := array.Get("length").Int()
	buf := array.Get("buffer")
	tmp := js.Global().Get("Uint8Array").New(buf)

	mem := make([]byte, size*2)
	js.CopyBytesToGo(mem, tmp)

	head := (*reflect.SliceHeader)(unsafe.Pointer(&mem))
	head.Cap /= 2
	head.Len /= 2
	dst := *(*[]int16)(unsafe.Pointer(head))
	runtime.KeepAlive(mem)
	return dst
}

// UInt32ToJs is converting []uint32 to a new javascript Uint32Array.
func UInt32ToJs(src []uint32) js.Value {
	array := js.Global().Get("Uint8Array").New(len(src) * 4)

	// pointer converter
	head := (*reflect.SliceHeader)(unsafe.Pointer(&src))
	head.Len *= 4
	head.Cap *= 4
	data := *(*[]byte)(unsafe.Pointer(head))

	// copy
	js.CopyBytesToJS(array, data)
	runtime.KeepAlive(src)

	// convert back
	buf := array.Get("buffer")
	return js.Global().Get("Uint32Array").New(
		buf,
		array.Get("byteOffset"),
		array.Get("byteLength").Int()/4,
	)
}

// JsToUInt32 convert javascript Uint32Array to Go uint32.
func JsToUInt32(array js.Value) []uint32 {
	size := array.Get("length").Int()
	buf := array.Get("buffer")
	tmp := js.Global().Get("Uint8Array").New(buf)

	mem := make([]byte, size*4)
	js.CopyBytesToGo(mem, tmp)

	head := (*reflect.SliceHeader)(unsafe.Pointer(&mem))
	head.Cap /= 4
	head.Len /= 4
	dst := *(*[]uint32)(unsafe.Pointer(head))
	runtime.KeepAlive(mem)
	return dst
}

// Int32ToJs is converting []int32 to a new javascript Int32Array.
func Int32ToJs(src []int32) js.Value {
	array := js.Global().Get("Uint8Array").New(len(src) * 4)

	// pointer converter
	head := (*reflect.SliceHeader)(unsafe.Pointer(&src))
	head.Len *= 4
	head.Cap *= 4
	data := *(*[]byte)(unsafe.Pointer(head))

	// copy
	js.CopyBytesToJS(array, data)
	runtime.KeepAlive(src)

	// convert back
	buf := array.Get("buffer")
	return js.Global().Get("Int32Array").New(
		buf,
		array.Get("byteOffset"),
		array.Get("byteLength").Int()/4,
	)
}

// JsToInt32 convert javascript Int32Array to Go int32.
func JsToInt32(array js.Value) []int32 {
	size := array.Get("length").Int()
	buf := array.Get("buffer")
	tmp := js.Global().Get("Uint8Array").New(buf)

	mem := make([]byte, size*4)
	js.CopyBytesToGo(mem, tmp)

	head := (*reflect.SliceHeader)(unsafe.Pointer(&mem))
	head.Cap /= 4
	head.Len /= 4
	dst := *(*[]int32)(unsafe.Pointer(head))
	runtime.KeepAlive(mem)
	return dst
}

// Float32ToJs is converting []float32 to a new javascript Float32Array.
func Float32ToJs(src []float32) js.Value {
	array := js.Global().Get("Uint8Array").New(len(src) * 4)

	// pointer converter
	head := (*reflect.SliceHeader)(unsafe.Pointer(&src))
	head.Len *= 4
	head.Cap *= 4
	data := *(*[]byte)(unsafe.Pointer(head))

	// copy
	js.CopyBytesToJS(array, data)
	runtime.KeepAlive(src)

	// convert back
	buf := array.Get("buffer")
	return js.Global().Get("Float32Array").New(
		buf,
		array.Get("byteOffset"),
		array.Get("byteLength").Int()/4,
	)
}

// JsToFloat32 convert javascript Float32Array to Go float32.
func JsToFloat32(array js.Value) []float32 {
	size := array.Get("length").Int()
	buf := array.Get("buffer")
	tmp := js.Global().Get("Uint8Array").New(buf)

	mem := make([]byte, size*4)
	js.CopyBytesToGo(mem, tmp)

	head := (*reflect.SliceHeader)(unsafe.Pointer(&mem))
	head.Cap /= 4
	head.Len /= 4
	dst := *(*[]float32)(unsafe.Pointer(head))
	runtime.KeepAlive(mem)
	return dst
}

// Float64ToJs is converting []float64 to a new javascript Float64Array.
func Float64ToJs(src []float64) js.Value {
	array := js.Global().Get("Uint8Array").New(len(src) * 8)

	// pointer converter
	head := (*reflect.SliceHeader)(unsafe.Pointer(&src))
	head.Len *= 8
	head.Cap *= 8
	data := *(*[]byte)(unsafe.Pointer(head))

	// copy
	js.CopyBytesToJS(array, data)
	runtime.KeepAlive(src)

	// convert back
	buf := array.Get("buffer")
	return js.Global().Get("Float64Array").New(
		buf,
		array.Get("byteOffset"),
		array.Get("byteLength").Int()/8,
	)
}

// JsToFloat64 convert javascript Float64Array to Go float64.
func JsToFloat64(array js.Value) []float64 {
	size := array.Get("length").Int()
	buf := array.Get("buffer")
	tmp := js.Global().Get("Uint8Array").New(buf)

	mem := make([]byte, size*8)
	js.CopyBytesToGo(mem, tmp)

	head := (*reflect.SliceHeader)(unsafe.Pointer(&mem))
	head.Cap /= 8
	head.Len /= 8
	dst := *(*[]float64)(unsafe.Pointer(head))
	runtime.KeepAlive(mem)
	return dst
}
